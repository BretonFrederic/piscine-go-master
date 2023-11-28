package piscine

// test := []string{"Hello", "how", "are", "you?"}

func ConcatParams(args []string) string {
	var aString string
	sizeParams := len(args)
	for i := 0; i < sizeParams; i++ {
		aString = aString + args[i]
		if i < sizeParams-1 {
			aString = aString + "\n"
		}
	}
	return aString
}
