package weather

import (
	"strconv"
)

func toDigital(number int) string {
	top := [10]string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "}
	middle := [10]string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"}
	bottom := [10]string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"}

	str := strconv.Itoa(number)

	var retIntTop string
	var retIntMid string
	var retIntBtm string
	
	for _, value := range str {
	toInt,_ := strconv.Atoi(string(value))
		retIntTop += top[toInt] 
		retIntMid += middle[toInt]
		retIntBtm += bottom[toInt]
	}
	
	return retIntTop + "\n" + retIntMid + "\n" + retIntBtm
}