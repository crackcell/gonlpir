all: test

install_deps:
	sudo cp deps/include/*.h /usr/include
	if [ -e /usr/lib64 ]; then sudo cp deps/lib/linux64/*.so /usr/lib64/; else sudo cp deps/lib/linux64/*.so /usr/lib/; fi

test:
	go test -v
