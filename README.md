# imagescaling
	imagescaling is a go lib for clip or scaling image. support for jpg/png/bmp/gif 【chaining operations】
	一个go图片 裁剪&缩放 库。 支持 jpg/png/bmp/gif 格式， 【链式操作】
	
#### 1.get
```bash
go get github/chi-chu/imagescaling
```
#### 2.Mode Explain
	

#### 3.Usage
```golang
import github/chi-chu/imagescaling

func main(){
	imagedata, err := os.Open("your/image/path/file.jpg")
	if err != nil {
		panic(err)
	}
	img, err := imagescaling.New(imagedata)
	if err != nil {
		panic(err)
	}
	outPutPath := "your/output/image/path/"
	outPutName := "filename" + img.GetExt()
	err = img.Clip(nil).Scale(nil).Draw()
	if err != nil {
		panic(err)
	}
}
```