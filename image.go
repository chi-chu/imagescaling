package imagescaling

import (
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