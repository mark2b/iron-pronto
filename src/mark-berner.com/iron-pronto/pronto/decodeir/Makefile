TARGET=decodeir

$(TARGET): libdecodeir.a
	go build .

libdecodeir.a: DecodeIR.o
	ar r $@ $^

%.o: %.cpp
	g++ -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)
