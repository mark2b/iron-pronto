package main

import (
	"fmt"
	"mark-berner.com/iron-pronto/pronto"
)

func main() {
	prontoCode := "0000 006D 0022 0002 0155 00AA 0015 0016 0015 0014 0016 0015 0015 0016 0015 0015 0016 0015 0015 0015 0015 0015 0016 0040 0015 003F 0015 0040 0016 0040 0015 003F 0015 0040 0016 003F 0015 0040 0015 0016 0015 0040 0015 0015 0015 0040 0015 0040 0016 0014 0015 0016 0015 0015 0016 0040 0015 0014 0016 0040 0015 0015 0016 0015 0015 003F 0016 0040 0015 003F 0015 05E4 0154 0056 0015 0E2E"
	fmt.Println(prontoCode)

	if decoded, err := pronto.Decode(prontoCode); err == nil {
		fmt.Println(decoded.Format())
	}
	if packed, err := pronto.Pack(prontoCode); err == nil {
		fmt.Println(packed)
		if unpacked, err := pronto.Unpack(packed); err == nil {
			fmt.Println(unpacked)
		}
	}
}
