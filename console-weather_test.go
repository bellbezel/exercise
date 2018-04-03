package weather

import "testing"

func Test_toDigital_0(t *testing.T) {
	expectedDigital := " _ \n" +
		"| |\n" +
		"|_|"
	digital := toDigital(0)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_1(t *testing.T) {
	expectedDigital := "   \n" +
		"  |\n" +
		"  |"
	digital := toDigital(1)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_2(t *testing.T) {
	expectedDigital := " _ \n" +
		" _|\n" +
		"|_ "
	digital := toDigital(2)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_3(t *testing.T) {
	expectedDigital := " _ \n" +
		" _|\n" +
		" _|"
	digital := toDigital(3)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_4(t *testing.T) {
	expectedDigital := "   \n" +
		"|_|\n" +
		"  |"
	digital := toDigital(4)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_5(t *testing.T) {
	expectedDigital := " _ \n" +
		"|_ \n" +
		" _|"
	digital := toDigital(5)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_6(t *testing.T) {
	expectedDigital := " _ \n" +
		"|_ \n" +
		"|_|"
	digital := toDigital(6)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

func Test_toDigital_7(t *testing.T) {
	expectedDigital := " _ \n" +
		"  |\n" +
		"  |"
	digital := toDigital(7)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}
func Test_toDigital_8(t *testing.T) {
	expectedDigital := " _ \n" +
		"|_|\n" +
		"|_|"
	digital := toDigital(8)
	if digital != expectedDigital {
		t.Error("Expected \n" + expectedDigital + " But actual is " + digital)
	}
}

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
