package utils

// PrettyLoad prettifies the loading output
func PrettyLoad(data string, leng int) string {
	for {
		lengtString := len(data)
		if lengtString < leng {
			data = data + ":"
			continue
		}
		break
	}
	return data
}
