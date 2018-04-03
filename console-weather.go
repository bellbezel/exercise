package weather

func toDigital(number int) string {
	top := [10]string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "}
	middle := [10]string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"}
	bottom := [10]string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"}

	divideNum := number / 10
	modNum := number % 10

	if divideNum > 0 {
		return top[divideNum] + top[modNum] + "\n" + middle[divideNum] + middle[modNum] + "\n" + bottom[divideNum] + bottom[modNum]
	} else {
		return top[number] + "\n" + middle[number] + "\n" + bottom[number]
	}
}
