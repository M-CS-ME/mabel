tangle:
	mabel mabel.md > src/mabel.go

install:
	go build src/mabel.go
	sudo cp mabel /usr/bin/mabel

package:
	make tangle
	make install
	pandoc mabel.md -o mabel.pdf
