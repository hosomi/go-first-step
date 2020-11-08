# Go Thumbnail Picture

* go version go1.15.2 windows/amd64
* [draw - GoDoc](https://godoc.org/golang.org/x/image/draw)  

## initial setup:  

```powershell
PS go-thumbnail-picture> go mod init hosomi/go-thumbnail-picture
go: creating new go.mod: module hosomi/go-thumbnail-picture
```

## run setup:

```powershell
PS go-thumbnail-picture> go build
go: finding module for package golang.org/x/image/draw
go: downloading golang.org/x/image v0.0.0-20200927104501-e162460cd6b5
go: found golang.org/x/image/draw in golang.org/x/image v0.0.0-20200927104501-e162460cd6b5
```

## go run halfsize.go

Halve the original image and output it.  

args:  
* [0] ... Original Image File Name.
* [1] ... Output Image File Name.

```powershell
PS go-thumbnail-picture> go run halfsize.go material/cat0056-051.jpg out.jpg
```

## Thanks.

* [無料の写真素材 - 無料画像・フリー素材のpro.foto（プロ・フォト）](https://pro-foto.jp/)

