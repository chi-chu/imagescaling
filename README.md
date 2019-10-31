# imagescaling
	imagescaling is a go lib for clip or scaling image. support for jpg/png/bmp/gif 【chaining operations】
	
	imagescaling 是一个go图片 裁剪&缩放 库。 支持 jpg/png/bmp/gif 格式， 【链式操作】
#### 1.get
```bash
go get github/chi-chu/imagescaling
```
#### 2.Mode Explain
* Clip()  &emsp;Api example:  
&emsp;&emsp;CenterMode:			Mode{Mode:CenterMode}  
&emsp;&emsp;![Image text]()  
&emsp;&emsp;CustomMode:    		Mode{Mode:CustomMode, Coordinate: [4]uint{&ensp;X0,&ensp;Y0,&ensp;X1,&ensp;Y1&ensp;}}  
&emsp;&emsp;![Image text]()  
* Scale() &emsp; Api example:  
	&emsp;&emsp;ProportionMode:		Mode{Mode:ProportionMode, Proportion: 0.5}	&emsp;&emsp;&emsp;&emsp;**(half size)**  
	&emsp;&emsp;![Image text]()  
	&emsp;&emsp;FixLengthMode:		Mode{Mode:FixLengthMode, FixHeight:80} 	&emsp;&emsp;&emsp;&emsp;&emsp;**(auto fix width)**  
	&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;                 or Mode{Mode:FixLengthMode, FixWidth:100}		&ensp;&emsp;&emsp;&emsp;&emsp;&emsp;**(auto fix height)**  
	&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;&emsp;  				 or Mode{Mode:FixLengthMode, FixHeight:123, FixWidth:456}  &emsp;&emsp;**(stretch 拉伸)**  
&emsp;&emsp;![Image text]()  


#### 3.Usage
```golang
import github/chi-chu/imagescaling

func main(){
	imageData, err := os.Open("/your/image/path/filename.jpg")
    if err != nil {
        panic(err)
    }
    defer imageData.Close()
    // here set the global opreation mode
    imagescaling.SetGlobalClipMode(imagescaling.Mode{Mode:imagescaling.CustomMode, Coordinate: [4]uint{0,0,123,300}})
    imagescaling.SetGlobalScaleMode(imagescaling.Mode{Mode:imagescaling.FixLengthMode, FixHeight:500})
    img, err := imagescaling.New(imageData)
    if err != nil {
        panic(err)
    }
    outPutPath := "/your/output/image/path/filename."+ img.GetExt()
    fd, err := os.OpenFile(outPutPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
    if err != nil {
        panic(err)
    }
    defer fd.Close()
    err = img.Clip(nil).Scale(nil).Draw(fd)
    if err != nil {
        panic(err)
    }
}
```