package imagescaling

type MODE int

const (
    CenterMode      MODE = 1
    ProportionMode  MODE = 2
    FixLengthMode   MODE = 3
    CustomMode      MODE = 4
)

type Mode struct {
    Mode            MODE
    FixHeight       uint
    FixWidth        uint
    Proportion      float64
    Coordinate      [4]uint
}

var globalClipMode Mode
var setedclip bool
var globalScaleMode Mode
var setedscale bool
func SetGlobalClipMode(s Mode){
    globalClipMode = s
    setedclip = true
}

func SetGlobalScaleMode(s Mode){
    globalScaleMode = s
    setedscale = true
}