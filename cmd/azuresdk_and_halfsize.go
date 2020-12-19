package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
	"golang.org/x/image/draw"
)

type Options struct {
	container    string
	downloadFile string
	uploadFile   string
}

var (
	o Options
)

func init() {
	flag.CommandLine.Init("command", flag.ExitOnError)

	flag.StringVar(&o.container, "container", "go-test", "Azure Storage Container name.")
	flag.StringVar(&o.downloadFile, "downloadFile", "cat0056-051.jpg", "Azure Storage Container Blob download file name.")
	flag.StringVar(&o.uploadFile, "uploadFile", "cat0056-051-out.jpg", "Azure Storage Container Blob upload file name.")
}

func handleErrors(err error) {
	if err != nil {
		if serr, ok := err.(azblob.StorageError); ok {
			switch serr.ServiceCode() {
			case azblob.ServiceCodeContainerNotFound:
				fmt.Println("Received 404. Container not exists")
				return
			}
		}
		log.Fatal(err)
	}
}

func main() {

	flag.Parse()
	fmt.Println("container:", o.container)
	fmt.Println("download file:", o.downloadFile)
	fmt.Println("upload file:", o.uploadFile)

	accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")
	if len(accountName) == 0 || len(accountKey) == 0 {
		log.Fatal("Either the AZURE_STORAGE_ACCOUNT or AZURE_STORAGE_ACCESS_KEY environment variable is not set")
	}

	// emulator
	URL, _ := url.Parse(
		fmt.Sprintf("http://127.0.0.1:10000/%s/%s", accountName, o.container))

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	containerURL := azblob.NewContainerURL(*URL, p)
	ctx := context.Background() // never-expiring context

	// download
	blobURL := containerURL.NewBlockBlobURL(o.downloadFile)

	downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)
	handleErrors(err)

	imageMemory, err := jpeg.Decode(downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20}))
	handleErrors(err)

	rct := imageMemory.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, rct.Dx()/2, rct.Dy()/2))
	draw.CatmullRom.Scale(dst, dst.Bounds(), imageMemory, rct, draw.Over, nil)

	// upload
	blobUploadURL := containerURL.NewBlockBlobURL(o.uploadFile)
	handleErrors(err)

	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, dst, &jpeg.Options{Quality: 90})
	handleErrors(err)
	imageBytes := buffer.Bytes()

	_, err = azblob.UploadBufferToBlockBlob(ctx, imageBytes, blobUploadURL,
		azblob.UploadToBlockBlobOptions{
			BlockSize:   4 * 1024 * 1024,
			Parallelism: 16})
	handleErrors(err)
}
