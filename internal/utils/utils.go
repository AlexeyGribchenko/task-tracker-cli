package utils

func ValueFromPointer(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func PointerFromValue(val string) *string {
	return &val
}
