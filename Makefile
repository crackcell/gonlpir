all: test

install_deps:
	sudo cp deps/lib/linux64/*.so /usr/lib64/
	sudo cp deps/include/*.h /usr/include

test:
	go test -v
