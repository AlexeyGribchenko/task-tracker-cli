package utils

func ValueFromPointer(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}
