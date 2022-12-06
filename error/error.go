package error

func Error(e error) {
	if e != nil {
		panic(e)
	}
}
