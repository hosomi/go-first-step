package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

type Options struct {
	container string
	file      string
}

var (
	o Options
)

func init() {
	flag.CommandLine.Init("command", flag.ExitOnError)

	flag.StringVar(&o.container, "container", "go-test", "Azure Storage Container name.")
	flag.StringVar(&o.file, "file", "cat0056-051.jpg", "Azure Storage Container Blob file name.")
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
	fmt.Println("file:", o.file)

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
	blobURL := containerURL.NewBlockBlobURL(o.file)

	downloadResponse, err := blobURL.Download(ctx, 0, azblob.CountToEnd, azblob.BlobAccessConditions{}, false)
	handleErrors(err)

	body, err := ioutil.ReadAll(downloadResponse.Body(azblob.RetryReaderOptions{MaxRetryRequests: 20}))
	handleErrors(err)

	file, err := os.OpenFile("out_"+o.file, os.O_CREATE|os.O_WRONLY, 0666)
	handleErrors(err)

	defer func() {
		file.Close()
	}()

	file.Write(body)

	fmt.Println("download filename:", "out_"+o.file)
}
