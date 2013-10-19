PREFIX=/usr/local
BINDIR=${PREFIX}/bin

all: build/splitd

build:
	mkdir build

build/splitd: build
	go build -o build/splitd

clean:
	rm -rf build

install: build/splitd
	install -m 755 -d ${BINDIR}
	install -m 755 build/splitd ${BINDIR}/splitd

uninstall:
	rm ${BINDIR}/splitd

test:
	cd splitd; go test

.PHONY: install uninstall clean all test