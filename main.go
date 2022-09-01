package main

import (
	"fmt"
	"strconv"
)

func main() {

	b := "11111111"
	fmt.Printf("binaryToDecimal:\t %v \t----> %v\n", b, binaryToDecimal(b))
	fmt.Printf("binaryToHex:\t\t %v \t----> %v\n", b, binaryToHex(b))

	d := int64(255)
	fmt.Printf("decimalToBinary:\t %v \t\t----> %v\n", d, decimalToBinary(d))
	fmt.Printf("decimalToBinaryAlt:\t %v \t\t----> %v\n", d, decimalToBinaryAlt(d))
	fmt.Printf("decimalToHex:\t\t %v \t\t----> %v\n", d, decimalToHex(d))
	fmt.Printf("decimalToHexAlt:\t %v \t\t----> %v\n", d, decimalToHexAlt(d))

	h := "ff"
	fmt.Printf("hexToBinary:\t\t %v \t\t----> %v\n", h, hexToBinary(h))
	fmt.Printf("hexToDecimal:\t\t %v \t\t----> %v\n", h, hexToDecimal(h))
}

// binarytoDecimal converts a binary string to a number
func binaryToDecimal(b string) int64 {
	d, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return d
}

// binaryToHex converts a binary string to a hexadecimal base 16 string
func binaryToHex(b string) string {
	d := binaryToDecimal(b)
	h := strconv.FormatInt(d, 16)
	return h
}

// decimalToBinary converts a number to a binary string
func decimalToBinary(d int64) string {
	b := strconv.FormatInt(d, 2)
	return b
}

// decimalToBinary converts a number to a binary string using an alternate method
func decimalToBinaryAlt(d int64) string {
	b := fmt.Sprintf("%b", d)
	return b
}

// decimaltoHex converts a number to hexadecimal string
func decimalToHex(d int64) string {
	h := strconv.FormatInt(d, 16)
	return h
}

// decimaltoHex converts a number to hexadecimal string using an alternate method
func decimalToHexAlt(d int64) string {
	b := fmt.Sprintf("%x", d)
	return b
}

// hexToDecimal converts a hexadecimal string to a number
func hexToDecimal(h string) int64 {
	d, err := strconv.ParseInt(h, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return d
}

// hexToDecimal converts a hexadecimal string to a binary string
func hexToBinary(h string) string {
	d := hexToDecimal(h)
	b := decimalToBinary(d)
	return b
}
