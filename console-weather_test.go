package weather

import "testing"

func Test_toDigital_9(t *testing.T) {
	expectedDigital := " _ \n" +
		"|_|\n" +
		" _|"
	digital := toDigital(9)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_25(t *testing.T) {
	expectedDigital := " _  _ \n" +
		" _||_ \n" +
		"|_  _|"
	digital := toDigital(25)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_34(t *testing.T) {
	expectedDigital := " _    \n" +
		" _||_|\n" +
		" _|  |"
	digital := toDigital(34)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_17(t *testing.T) {
	expectedDigital := "    _ \n" +
		"  |  |\n" +
		"  |  |"
	digital := toDigital(17)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}
func Test_toDigital_89(t *testing.T) {
	expectedDigital := " _  _ \n" +
		"|_||_|\n" +
		"|_| _|"
	digital := toDigital(89)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_889(t *testing.T) {
	expectedDigital := " _  _  _ \n" +
		"|_||_||_|\n" +
		"|_||_| _|"
	digital := toDigital(889)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}
