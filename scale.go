package imagescaling

import "github.com/nfnt/resize"

func (m *Image) scale(h, w uint) {
    if m.opImage == nil {
        m.opImage = resize.Resize(w, h, m.image, resize.Lanczos3)
    }else{
        m.opImage = resize.Resize(w, h, m.opImage, resize.Lanczos3)
    }
}

func (m *Image) ProportionScale(p float64) *Image {
    if m.err != nil {
        return m
    }
    if p < 0 {
        m.err = ERRPROPORTION
        return m
    }
    if m.opImage != nil {
        m.scale(uint(float64(m.opImage.Bounds().Max.Y) * p), uint(float64(m.opImage.Bounds().Max.X)*p))
        return m
    }
    m.scale(uint(float64(m.image.Bounds().Max.Y)*p), uint(float64(m.image.Bounds().Max.X)*p))
    return m
}

func (m *Image) FixScale(h, w uint) *Image {
    if m.err != nil {
        return m
    }
    if h == 0 && w == 0 {
        m.err = ERRFIXLENGTHSCALE
        return m
    }
    if m.opImage != nil {
        if h == 0 && w != 0 {
            ratio := w/uint(m.opImage.Bounds().Max.X)
            m.scale(uint(m.opImage.Bounds().Max.Y)*ratio, w)
            return m
        }
        if h != 0 && w == 0 {
            ratio := h/uint(m.opImage.Bounds().Max.Y)
            m.scale(h, uint(m.opImage.Bounds().Max.X)*ratio)
            return m
        }
    }
    if h == 0 && w != 0 {
        ratio := w/uint(m.image.Bounds().Max.X)
        m.scale(uint(m.image.Bounds().Max.Y)*ratio, w)
        return m
    }
    if h != 0 && w == 0 {
        ratio := h/uint(m.image.Bounds().Max.Y)
        m.scale(h, uint(m.image.Bounds().Max.X)*ratio)
    }
    return m
}