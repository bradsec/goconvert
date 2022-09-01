# Go Convert
### Learning go basics testing and benchmarks

`main.go` includes various small functions to convert between decimal, hexadecimal and binary.

Sample output `go run main.go`:  

```terminal
binaryToDecimal:	 11111111 	----> 255
binaryToHex:		 11111111 	----> ff
decimalToBinary:	 255 		----> 11111111
decimalToBinaryAlt:	 255 		----> 11111111
decimalToHex:		 255 		----> ff
decimalToHexAlt:	 255 		----> ff
hexToBinary:		 ff 		----> 11111111
hexToDecimal:		 ff 		----> 255
```

Sample output `go test -v`:  

```terminal
go test -v      
=== RUN   TestBinaryToDecimal
--- PASS: TestBinaryToDecimal (0.00s)
=== RUN   TestBinaryToHex
--- PASS: TestBinaryToHex (0.00s)
=== RUN   TestDecimalToBinary
--- PASS: TestDecimalToBinary (0.00s)
=== RUN   TestDecimalToBinaryAlt
--- PASS: TestDecimalToBinaryAlt (0.00s)
=== RUN   TestDecimalToHex
--- PASS: TestDecimalToHex (0.00s)
=== RUN   TestDecimalToHexAlt
--- PASS: TestDecimalToHexAlt (0.00s)
=== RUN   TestHextoBinary
--- PASS: TestHextoBinary (0.00s)
=== RUN   TestHextoDecimal
--- PASS: TestHextoDecimal (0.00s)
PASS
ok  	github.com/bradsec/goconvert	0.001s
```

Sample output `go test -bench=.`

```terminal
BenchmarkDecimalToBinary-12       	38662970	        28.96 ns/op
BenchmarkDecimalToBinaryAlt-12    	41990098	        26.90 ns/op
BenchmarkDecimalToHex-12          	44431413	        23.88 ns/op
BenchmarkDecimalToHexAlt-12       	49089007	        23.91 ns/op
PASS
ok  	github.com/bradsec/goconvert	4.604s
```

