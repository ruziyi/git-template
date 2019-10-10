package op

func Get(b bool, val1, val2 interface{}) interface{} {
	if b {
		return val1
	}
	return val2
}
