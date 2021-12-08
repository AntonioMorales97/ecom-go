package util

const (
	PRODUCT int64 = 1
	SERVICE int64 = 2
)

func IsSupportedProductType(productTypeID int64) bool {
	switch productTypeID {
	case PRODUCT, SERVICE:
		return true
	}
	return false
}
