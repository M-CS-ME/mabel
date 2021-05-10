tangle:
	mabel mabel.md > src/mabel.go

install:
	go build src/mabel.go
	sudo cp mabel /usr/bin/mabel

package:
	mabel mabel.md > src/mabel.go
	go build src/mabel.go
	sudo cp mabel /usr/bin/mabel
	make install
	pandoc mabel.md -o mabel.pdf
