# Go Thumbnail Picture

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



## Thanks.

* [無料の写真素材 - 無料画像・フリー素材のpro.foto（プロ・フォト）](https://pro-foto.jp/)

