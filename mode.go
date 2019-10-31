package imagescaling

type MODE int

const (
    CenterMode      MODE = 0
    ProportionMode  MODE = 1
    CustomMode      MODE = 2
)

type Mode struct {
    Mode            MODE
    CustomHeight    uint
    CustomWidth     uint
    Proportion      float64
    Coordinate      [4]uint
}

var globalClipMode Mode
var setedclip bool
var globalScaleMode Mode
var setedscale bool
func SetGlobalClipMode(s Mode) error{
    globalClipMode = s
    setedclip = true

}

func SetGlobalScaleMode(s Mode) error{
    globalScaleMode = s
    setedscale = true
}