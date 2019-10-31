package imagescaling

import(
    "errors"
)
var (
    ERRUNKNOWEXT        = errors.New("unknown image type")

    ERRSCALE            = errors.New("[Scale Mode] only support ProportionMode or CustomMode")
    ERRPROPORTION       = errors.New("[Scale ProportionMode] proportion must greater than zero")
    ERRCUSTOMSCALE      = errors.New("[Scale CustomMode] custom height and width are zero at the same time")
    ERRNILSCALEMODE     = errors.New("[Scale Mode] is nil and you forget to Set Global Scale Mode")

    ERRCLIP             = errors.New("[Clip Mode] only support CenterMode or CustomMode")
    ERRCUSTOMCLIP       = errors.New("[Clip CustomMode] coordinate should be set")
    ERRCUSTOMCLIPPARA   = errors.New("[Clip CustomMode] wrong coordinate parameter:  X2<=X1 or Y2<=Y1 or Greater than image X,Y)")
    ERRNILCLIPMODE      = errors.New("[Clip Mode] is nil and you forget to Set Global Clip Mode")
)
