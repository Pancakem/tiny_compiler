build-debug:
	go build -gcflags '-N -l' .

build:
	go build .

clean:
	rm tiny_compiler
