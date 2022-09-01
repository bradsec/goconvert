package main

import "testing"

type tests struct {
	dec int64
	bin string
	hex string
}

func loadTestData() []tests {
	testData := []tests{
		{255, "11111111", "ff"},
		{128, "10000000", "80"},
		{64, "1000000", "40"},
		{32, "100000", "20"},
		{17, "10001", "11"},
		{21, "10101", "15"},
		{42, "101010", "2a"},
	}
	return testData
}

func TestBinaryToDecimal(t *testing.T) {
	testData := loadTestData()
	for _, v := range testData {
		x := binaryToDecimal(v.bin)
		if x != v.dec {
			t.Errorf("Got: %v, Want: %v", x, v.dec)
		}
	}
}

func TestBinaryToHex(t *testing.T) {
	testData := loadTestData()
	for _, v := range testData {
		x := binaryToHex(v.bin)
		if x != v.hex {
			t.Errorf("Got: %v, Want: %v", x, v.hex)
		}
	}
}

func TestDecimalToBinary(t *testing.T) {
	testData := loadTestData()
	for _, v := range testData {
		x := decimalToBinary(v.dec)
		if x != v.bin {
			t.Errorf("Got: %v, Want: %v", x, v.bin)
		}
	}
}

func TestDecimalToBinaryAlt(t *testing.T) {
	testData := loadTestData()
	for _, v := range testData {
		x := decimalToBinaryAlt(v.dec)
		if x != v.bin {
			t.Errorf("Got: %v, Want: %v", x, v.bin)
		}
	}
}

func TestDecimalToHex(t *testing.T) {
	testData := loadTestData()
	for _, v := range testData {
		x := decimalToHex(v.dec)
		if x != v.hex {
			t.Errorf("Got: %v, Want: %v", x, v.hex)
		}
	}
}

func TestDecimalToHexAlt(t *testing.T) {
	testData := loadTestData()
	for _, v := range testData {
		x := decimalToHexAlt(v.dec)
		if x != v.hex {
			t.Errorf("Got: %v, Want: %v", x, v.hex)
		}
	}
}

func TestHextoBinary(t *testing.T) {
	testData := loadTestData()
	for _, v := range testData {
		x := hexToBinary(v.hex)
		if x != v.bin {
			t.Errorf("Got: %v, Want: %v", x, v.bin)
		}
	}
}

func TestHextoDecimal(t *testing.T) {
	testData := loadTestData()
	for _, v := range testData {
		x := hexToDecimal(v.hex)
		if x != v.dec {
			t.Errorf("Got: %v, Want: %v", x, v.dec)
		}
	}
}

func BenchmarkDecimalToBinary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decimalToBinary(255)
	}
}

func BenchmarkDecimalToBinaryAlt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decimalToBinary(255)
	}
}

func BenchmarkDecimalToHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decimalToBinary(255)
	}
}

func BenchmarkDecimalToHexAlt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decimalToBinary(255)
	}
}
