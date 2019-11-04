package imagescaling

import "image"

func (m *Image) clip(x0, y0, x1, y1 int) {
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
}

func (m *Image) CenterClip() *Image {
    if m.err != nil {
        return m
    }
    if m.opImage != nil {
        halfWidth := int(m.opImage.Bounds().Max.X/2)
        halfHeight := int(m.opImage.Bounds().Max.Y/2)
        if m.opImage.Bounds().Max.Y > m.opImage.Bounds().Max.X {
            m.clip(0,halfHeight-halfWidth, m.opImage.Bounds().Max.X, halfHeight+halfWidth)
            return m
        }
        m.clip(halfWidth-halfHeight,0, halfHeight+halfWidth, m.opImage.Bounds().Max.Y)
        return m
    }
    halfWidth := int(m.image.Bounds().Max.X/2)
    halfHeight := int(m.image.Bounds().Max.Y/2)
    if m.image.Bounds().Max.Y > m.image.Bounds().Max.X{
        m.clip(0,halfHeight-halfWidth, m.image.Bounds().Max.X, halfHeight+halfWidth)
        return m
    }
    m.clip(halfWidth-halfHeight,0, halfHeight+halfWidth, m.image.Bounds().Max.Y)
    return m
}

func (m *Image) CustomClip(x0,y0,x1,y1 uint) *Image {
    if m.err != nil {
        return m
    }
    if (x0 >= x1) ||
        (y0 >= y1) {
        m.err = ERRCUSTOMCLIP
        return m
    }
    if m.opImage != nil {
        if x1 > uint(m.opImage.Bounds().Max.X) ||
            y1 > uint(m.opImage.Bounds().Max.Y) {
            m.err = ERRCUSTOMCLIPPARA
            return m
        }
    }
    if x1 > uint(m.image.Bounds().Max.X) ||
        y1 > uint(m.image.Bounds().Max.Y) {
        m.err = ERRCUSTOMCLIPPARA
        return m
    }
    m.clip(int(x0), int(y0), int(x1), int(y1))
    return m
}