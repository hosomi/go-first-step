# Go First Step

* go version go1.15.2 windows/amd64
* [draw - GoDoc](https://godoc.org/golang.org/x/image/draw)  

## initial setup:  

```powershell
PS go-thumbnail-picture> go mod init hosomi/go-thumbnail-picture
go: creating new go.mod: module hosomi/go-thumbnail-picture
```

## go run cmd/halfsize.go

Halve the original image and output it.  

args:  
* [0] ... Original Image File Name.
* [1] ... Output Image File Name.

```powershell
PS go-thumbnail-picture> go run cmd/halfsize.go material/cat0056-051.jpg out.jpg
```

---

## go run cmd/superposition.go

Superposition two images.  

```powershell
PS go-thumbnail-picture> go run cmd/superposition.go
```

lower image(100x100):  
![lower image](material/100x100.jpg)  

upper image(50x50):  
![upper image](material/50x50.jpg)  

output image(Starting position 25,25):  
![output image](material/superposition.jpg)  

---

## go run cmd/azuresdk_blobupload.go

Upload the files to Azure Storage Blob.  
(Uploading files to Azure Blob with Azure SDK for Go)    

args:  
* [0] ... Upload File Name.

``The container name must be created in go-test before it is run.``

```powershell
PS go-thumbnail-picture> go run .\cmd\azuresdk_blobupload.go .\material\cat0056-051.jpg
```

setup:  
* Windows (Azure Storage Emulator settings)
:link: [接続文字列を構成する - Azure Storage | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/storage/common/storage-configure-connection-string)

```powershell
setx AZURE_STORAGE_ACCOUNT "devstoreaccount1"
setx AZURE_STORAGE_ACCESS_KEY "Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw=="
```

:link: [Azure/azure-storage-blob-go: Microsoft Azure Blob Storage Library for Go](https://github.com/Azure/azure-storage-blob-go)  
:link: [Go 開発者向けの Azure | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/developer/go/)  
:link: [Azure クイック スタート - Go を使用してオブジェクト ストレージに BLOB を作成する | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/storage/blobs/storage-quickstart-blobs-go)  
:link: [setx | Microsoft Docs](https://docs.microsoft.com/ja-jp/windows-server/administration/windows-commands/setx)  

---


## go run cmd/azuresdk_blobdownload.go

Download the files from Azure Storage Blob.

args:  
* --container ... Azure Storage Container name.
* --file ... Azure Storage Container Blob file name. 

usage: go run cmd/azuresdk_blobdownload.go --help  


## Thanks.

* [無料の写真素材 - 無料画像・フリー素材のpro.foto（プロ・フォト）](https://pro-foto.jp/)

