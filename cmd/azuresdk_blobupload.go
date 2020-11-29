package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/Azure/azure-storage-blob-go/azblob"
)

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
	args := flag.Args()

	f, err := os.Open(args[0])
	if err != nil {
		fmt.Println("open:", err)
		return
	}
	defer f.Close()

	accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT"), os.Getenv("AZURE_STORAGE_ACCESS_KEY")
	if len(accountName) == 0 || len(accountKey) == 0 {
		log.Fatal("Either the AZURE_STORAGE_ACCOUNT or AZURE_STORAGE_ACCESS_KEY environment variable is not set")
	}

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Invalid credentials with error: " + err.Error())
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	containerName := "go-test"

	// Azure
	// URL, _ := url.Parse(
	// 	fmt.Sprintf("https://%s.blob.core.windows.net/%s", accountName, containerName))

	// emulator
	URL, _ := url.Parse(
		fmt.Sprintf("http://127.0.0.1:10000/%s/%s", accountName, containerName))

	containerURL := azblob.NewContainerURL(*URL, p)
	ctx := context.Background() // never-expiring context

	fileName := args[0]
	blobURL := containerURL.NewBlockBlobURL(fileName)
	file, err := os.Open(fileName)
	defer f.Close()
	handleErrors(err)

	_, err = azblob.UploadFileToBlockBlob(ctx, file, blobURL,
		azblob.UploadToBlockBlobOptions{
			BlockSize:   4 * 1024 * 1024,
			Parallelism: 16})
	handleErrors(err)
}
