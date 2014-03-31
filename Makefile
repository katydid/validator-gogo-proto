.PHONY: nuke regenerate gofmt build test

all: nuke regenerate gofmt build test

test:
	go test -v ./...

build:
	go build ./...

bench:
	go test -v -test.run=XXX -test.bench=. ./...

regenerate:
	(cd asm && gocc asm.bnf)
	(cd asm/ast && protoc --gogo_out=. -I=.:../../../../../ asm.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ person.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ srctree.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ taxonomy.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ treeregister.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ typewriterprison.proto)
	(cd asm/test && protoc --gogo_out=. -I=.:../../../../../ puddingmilkshake.proto)
	make gofmt

clean:
	(cd asm && rm *.txt || true)
	rm -rf asm/test/example/*
	go clean ./...

example:
	(cd asm/test && go test -c && ./test.test)

nuke: clean
	rm -rf asm/errors
	rm -rf asm/lexer
	rm -rf asm/parser
	rm -rf asm/token
	rm -rf asm/util
	go clean -i ./...

gofmt:
	gofmt -l -s -w .