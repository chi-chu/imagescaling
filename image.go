package imagescaling

import (
    "github.com/nfnt/resize"
    "golang.org/x/image/bmp"
    "image"
    "image/gif"
    "image/jpeg"
    "image/png"
    "io"
)

const (
    DefaultQuality       = 100
)

type Image struct {
    image           image.Image
    opImage         image.Image
    ext             string
    quality         int
    err             error
}

func New(in io.Reader)(*Image, error) {
    origin, fm, err := image.Decode(in)
    if err != nil {
        return nil, err
    }
    if fm != "jpeg" && fm != "png" && fm != "gif" && fm != "bmp" {
        return nil, ERRUNKNOWEXT
    }
    o := Image{}
    o.ext = fm
    o.image = origin
    return &o, nil
}

func (m *Image) calcClip(s *Mode)(int, int, int, int){
    if m.err != nil {
        return 0, 0, 0, 0
    }
    if s == nil {
        if setedclip == false {
            m.err = ERRNILCLIPMODE
            return 0,0,0,0
        }
        s = &globalClipMode
    }
    switch s.Mode {
    case CenterMode:
        if m.opImage != nil {
            halfWidth := int(m.opImage.Bounds().Max.X/2)
            halfHeight := int(m.opImage.Bounds().Max.Y/2)
            if m.opImage.Bounds().Max.Y > m.opImage.Bounds().Max.X {
                return 0,halfHeight-halfWidth, m.opImage.Bounds().Max.X, halfHeight+halfWidth
            }
            return halfWidth-halfHeight,0, halfHeight+halfWidth, m.opImage.Bounds().Max.Y
        }
        halfWidth := int(m.image.Bounds().Max.X/2)
        halfHeight := int(m.image.Bounds().Max.Y/2)
        if m.image.Bounds().Max.Y > m.image.Bounds().Max.X{
            return 0,halfHeight-halfWidth, m.image.Bounds().Max.X, halfHeight+halfWidth
        }
        return halfWidth-halfHeight,0, halfHeight+halfWidth, m.image.Bounds().Max.Y
    case CustomMode:
        if (s.Coordinate[0] >= s.Coordinate[2]) ||
            (s.Coordinate[0] >= s.Coordinate[3]) {
            m.err = ERRCUSTOMCLIP
            return 0, 0, 0, 0
        }
        if m.opImage != nil {
            if s.Coordinate[2] > uint(m.opImage.Bounds().Max.X) ||
                s.Coordinate[3] > uint(m.opImage.Bounds().Max.Y) {
                m.err = ERRCUSTOMCLIPPARA
                return 0, 0, 0, 0
            }
            return int(s.Coordinate[0]), int(s.Coordinate[1]), int(s.Coordinate[2]), int(s.Coordinate[3])
        }
        if s.Coordinate[2] > uint(m.image.Bounds().Max.X) ||
            s.Coordinate[3] > uint(m.image.Bounds().Max.Y) {
            m.err = ERRCUSTOMCLIPPARA
            return 0, 0, 0, 0
        }
        return int(s.Coordinate[0]), int(s.Coordinate[1]), int(s.Coordinate[2]), int(s.Coordinate[3])
    }

    m.err =  ERRCLIP
    return 0,0,0,0
}

func (m *Image) calcScale(s *Mode)(uint, uint) {
    if m.err != nil {
        return 0, 0
    }
    if s == nil {
        if setedscale == false {
            m.err = ERRNILSCALEMODE
            return 0,0
        }
        s = &globalScaleMode
    }
    switch s.Mode {
    case ProportionMode:
        if s.Proportion < 0 {
            m.err = ERRPROPORTION
            return 0,0
        }
        if m.opImage != nil {
            return uint(float64(m.opImage.Bounds().Max.Y)*s.Proportion), uint(float64(m.opImage.Bounds().Max.X)*s.Proportion)
        }
        return uint(float64(m.image.Bounds().Max.Y)*s.Proportion), uint(float64(m.image.Bounds().Max.X)*s.Proportion)
    case FixLengthMode:
        if s.FixHeight == 0 && s.FixWidth == 0 {
            m.err = ERRFIXLENGTHSCALE
            return 0,0
        }
        if m.opImage != nil {
            if s.FixHeight == 0 && s.FixWidth != 0 {
                ratio := s.FixWidth/uint(m.opImage.Bounds().Max.X)
                return uint(s.FixHeight*ratio), s.FixWidth
            }
            if s.FixHeight != 0 && s.FixWidth == 0 {
                ratio := s.FixHeight/uint(m.opImage.Bounds().Max.Y)
                return s.FixHeight, uint(s.FixHeight*ratio)
            }
        }
        if s.FixHeight == 0 && s.FixWidth != 0 {
            ratio := s.FixWidth/uint(m.image.Bounds().Max.X)
            return uint(s.FixHeight*ratio), s.FixWidth
        }
        if s.FixHeight != 0 && s.FixWidth == 0 {
            ratio := s.FixHeight/uint(m.image.Bounds().Max.Y)
            return s.FixHeight, uint(s.FixHeight*ratio)
        }
        return s.FixHeight, s.FixWidth
    }
    m.err = ERRSCALE
    return 0,0
}

func (m *Image) GetExt() string {
    return m.ext
}

func (m *Image) SetQuality(q int) *Image {
    m.quality = q
    return m
}

func (m *Image) GetHeightWidth() (int, int) {
    return m.image.Bounds().Max.Y, m.image.Bounds().Max.X
}

func (m *Image) Clip(s *Mode) *Image {
    x0, y0, x1, y1 := m.calcClip(s)
    if m.err != nil {
        return m
    }
    switch m.ext {
    case "jpeg":
        if m.opImage == nil {
            m.opImage = m.image.(*image.YCbCr).SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
        }else{
            m.opImage = m.opImage.(*image.YCbCr).SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)
        }
    case "png":
        switch m.image.(type) {
        case *image.NRGBA:
            if m.opImage == nil {
                m.opImage = m.image.(*image.NRGBA).SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
            }else{
                m.opImage = m.opImage.(*image.NRGBA).SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
            }
        case *image.RGBA:
            if m.opImage == nil {
                m.opImage = m.image.(*image.RGBA).SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
            }else{
                m.opImage = m.opImage.(*image.RGBA).SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
            }
        }
    case "gif":
        if m.opImage == nil {
            m.opImage = m.image.(*image.Paletted).SubImage(image.Rect(x0, y0, x1, y1)).(*image.Paletted)
        }else{
            m.opImage = m.opImage.(*image.Paletted).SubImage(image.Rect(x0, y0, x1, y1)).(*image.Paletted)
        }
    case "bmp":
        if m.opImage == nil {
            m.opImage = m.image.(*image.RGBA).SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
        }else{
            m.opImage = m.opImage.(*image.RGBA).SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
        }
    }
    return m
}

func (m *Image) Scale(s *Mode) *Image {
    h, w := m.calcScale(s)
    if m.err != nil {
        return m
    }
    if m.opImage == nil {
        m.opImage = resize.Resize(w, h, m.image, resize.Lanczos3)
    }else{
        m.opImage = resize.Resize(w, h, m.opImage, resize.Lanczos3)
    }
    return m
}

func (m *Image) Draw(out io.Writer) error {
    if m.err != nil {
        return m.err
    }
    quality := DefaultQuality
    if m.quality != 0 {
        quality = m.quality
    }
    if m.opImage == nil {
        m.opImage = m.image
    }
    switch m.ext {
    case "jpeg":
        return jpeg.Encode(out, m.opImage, &jpeg.Options{Quality:quality})
    case "png":
        return png.Encode(out, m.opImage)
    case "gif":
        return gif.Encode(out, m.opImage, &gif.Options{})
    case "bmp":
        return bmp.Encode(out, m.opImage)
    default:
        return ERRUNKNOWEXT
    }
}

func (m *Image) ReSet() *Image {
    m.opImage = nil
    m.err = nil
    return m
}