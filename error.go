package imagescaling

import(
    "errors"
)

var (
    ERRUNKNOWEXT        = errors.New("unknown image type")

    ERRPROPORTION       = errors.New("[Scale ProportionMode] proportion must greater than zero")
    ERRFIXLENGTHSCALE   = errors.New("[Scale FixLengthMode] custom height and width are zero at the same time")

    ERRCUSTOMCLIP       = errors.New("[Clip CustomMode] coordinate should be set")
    ERRCUSTOMCLIPPARA   = errors.New("[Clip CustomMode] wrong coordinate parameter:  X2<=X1 or Y2<=Y1 or Greater than image X,Y)")
    ERRUNKNOWROTATEMODE = errors.New("unknown rotate mode")
)
