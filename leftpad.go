package leftpad

func Leftpad(str string, pad int) string {
	return generatePadding(pad) + str
}

func Rightpad(str string, pad int) string {
	return str + generatePadding(pad)
}

func generatePadding(qty int) string {
	padding := ""
	for i := 0; i < qty; i++ {
		padding += " "
	}
	return padding
}
