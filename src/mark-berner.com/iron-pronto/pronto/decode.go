package pronto

// #cgo LDFLAGS: -lstdc++
// #include "DecodeIR.h"
import "C"
import (
	"fmt"
	"strings"
	"unsafe"
)

type Decoded struct {
	Protocol  string
	Device    int
	Subdevice int
	Obc       int
	Hex       [4]int
}

func (self Decoded) Format() string {
	formatted := make([]string, 0)
	formatted = append(formatted, self.Protocol)
	formatted = append(formatted, fmt.Sprintf("device:%d", self.Device))
	if self.Subdevice != -1 {
		formatted = append(formatted, fmt.Sprintf("subdevice:%d", self.Subdevice))
	}
	if self.Obc != -1 {
		formatted = append(formatted, fmt.Sprintf("obc:%d", self.Obc))
	}
	for i, h := range self.Hex {
		if h != -1 {
			formatted = append(formatted, fmt.Sprintf("hex%d:%x", i, h))
		}
	}
	return strings.Join(formatted, ",")
}

func Decode(pronto string) (decoded Decoded, err error) {
	var frequency int32
	var intro_length int32
	var rep_length int32
	data := make([]int32, 512)
	C.get_hex_code(C.CString(pronto),
		(*C.int)(unsafe.Pointer(&frequency)),
		(*C.int)(unsafe.Pointer(&intro_length)),
		(*C.int)(unsafe.Pointer(&rep_length)),
		(*C.int)(unsafe.Pointer(&data[0])))

	decodeir_context := []uint32{0, 0}
	protocol := make([]byte, 255)
	var device int32 = -1
	var subdevice int32 = -1
	var obc int32 = -1
	hex := []int32{-1, -1, -1, -1}
	misc_message := make([]byte, 255)
	error_message := make([]byte, 255)

	decodeIR(decodeir_context, data, frequency, intro_length, rep_length,
		protocol, &device, &subdevice, &obc, hex, misc_message, error_message)

	decoded.Protocol = strings.TrimRight(string(protocol), "\x00")
	decoded.Device = int(device)
	decoded.Subdevice = int(subdevice)
	decoded.Obc = int(obc)
	for i, h := range hex {
		decoded.Hex[i] = int(h)
	}
	return
}

func decodeIR(decodeir_context []uint32,
	data []int32,
	frequency int32,
	intro_length int32,
	rep_length int32,
	protocol []byte,
	device *int32,
	subdevice *int32,
	obc *int32,
	hex []int32,
	misc_message []byte,
	error_message []byte) {
	C.DecodeIR((*C.uint)(unsafe.Pointer(&decodeir_context[0])),
		(*C.int)(unsafe.Pointer(&data[0])),
		C.int(frequency),
		C.int(intro_length),
		C.int(rep_length),
		(*C.char)(unsafe.Pointer(&protocol[0])),
		(*C.int)(unsafe.Pointer(device)),
		(*C.int)(unsafe.Pointer(subdevice)),
		(*C.int)(unsafe.Pointer(obc)),
		(*C.int)(unsafe.Pointer(&hex[0])),
		(*C.char)(unsafe.Pointer(&misc_message)),
		(*C.char)(unsafe.Pointer(&error_message)))
}
