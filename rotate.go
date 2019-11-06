package imagescaling

import (
    "image"
    "runtime"
    "sync"
)

type RotateMode int
const(
    Rotate90Degrees     RotateMode = 90
    Rotate180Degrees    RotateMode = 180
    Rotate270Degrees    RotateMode = 270
)

func (m *Image) Rotate(r RotateMode) *Image {
    if m.err != nil {
        return m
    }
    if m.opImage == nil {
        m.opImage = m.image
    }
    var wImg, hImg int
    switch r {
    case Rotate90Degrees, Rotate270Degrees:
        wImg = m.opImage.Bounds().Dy()
        hImg = m.opImage.Bounds().Dx()
    case Rotate180Degrees:
        wImg = m.opImage.Bounds().Dx()
        hImg = m.opImage.Bounds().Dy()
    default:
        m.err = ERRUNKNOWROTATEMODE
        return m
    }
    rowSize := wImg * 4
    m.tempImage = image.NewNRGBA(image.Rect(0, 0, wImg, hImg))
    procs := runtime.GOMAXPROCS(0)
    if procs > hImg {
        procs = hImg
    }
    c := make(chan int, hImg)
    for i:=0; i< hImg; i++{
        c <- i
    }
    close(c)
    var wg sync.WaitGroup
    for i := 0; i < procs; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for nh := range c {
                i := nh * m.tempImage.Stride
                switch r {
                case Rotate90Degrees:
                    srcX := nh
                    m.adjust(srcX, 0, srcX+1, m.opImage.Bounds().Max.Y, m.tempImage.Pix[i:i+rowSize])
                    reverse(m.tempImage.Pix[i : i+rowSize])
                case Rotate180Degrees:
                    srcY := m.opImage.Bounds().Max.Y - nh - 1
                    m.adjust(0, srcY, m.opImage.Bounds().Max.X, srcY+1, m.tempImage.Pix[i:i+rowSize])
                    reverse(m.tempImage.Pix[i : i+rowSize])
                case Rotate270Degrees:
                    srcX := hImg - nh - 1
                    m.adjust(srcX, 0, srcX+1, m.opImage.Bounds().Max.Y, m.tempImage.Pix[i:i+rowSize])
                }
            }
        }()
    }
    wg.Wait()
    m.opImage = m.tempImage
    m.tempImage = nil
    return m
}


