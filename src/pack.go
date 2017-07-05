package pronto

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"strings"
)

func Pack(pronto string) (packed string, err error) {
	return pack(pronto)
}

func Unpack(packed string) (pronto string, err error) {
	return unpack(packed)
}

func pack(input string) (output string, e error) {
	output = strings.Replace(input, " ", "", -1)
	if compressedData, err := compress(output); err == nil {
		output = base64.StdEncoding.EncodeToString(compressedData)
	} else {
		e = err
	}
	return
}

func compress(input string) (output []byte, e error) {
	inflated := new(bytes.Buffer)
	if flateWriter, err := flate.NewWriter(inflated, flate.BestCompression); err == nil {
		flateWriter.Write([]byte(input))
		flateWriter.Close()
		output = inflated.Bytes()
	} else {
		e = err
	}
	return
}

func unpack(input string) (output string, e error) {
	if decodedData, err := base64.StdEncoding.DecodeString(input); err == nil {
		if decompressedData, err := decompress(decodedData); err == nil {
			if len(decompressedData)%4 != 0 {
				e = errors.New("Invalid input length")
			}

			outputData := make([]byte, len(decompressedData)*5/4)
			j := 0
			step := 4
			for i := 0; i < len(decompressedData); i += step {
				if copy(outputData[j:], decompressedData[i:i+step]) == step {
					outputData[j+step] = ' '
					j += step + 1
				} else {
					e = errors.New("Invalid input length")
				}
			}
			output = strings.TrimSpace(string(outputData))
		} else {
			e = err
		}
	} else {
		e = err
	}
	return
}

func decompress(input []byte) (output []byte, e error) {
	reader := bytes.NewReader(input)
	flateReader := flate.NewReader(reader)
	if inflated, err := ioutil.ReadAll(flateReader); err == nil {
		output = inflated
	} else {
		e = err
	}
	return
}
