# Copyright 2013 Walter Schulze
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: nuke dep regenerate gofmt build test

all: nuke dep regenerate build test vet

dep:
	go install github.com/gogo/protobuf/protoc-gen-gogo
	go install -v github.com/goccmack/gocc
	go install -v github.com/awalterschulze/goderive

checklicense:
	go get github.com/katydid/checklicense
	checklicense . \
	person.proto \
	srctree.proto \
	puddingmilkshake.proto \
	taxonomy.proto \
	treeregister.proto \
	typewriterprison.proto \
	proto/tokens/test.proto \
	parser/debug/debug.proto \
	doc.go \
	tools/tools.go \
	.svg \
	.txt \
	install_protoc.sh

test:
	go test ./...

build:
	go build ./...

install:
	go install ./...

bench:
	go test -test.v -test.run=XXX -test.bench=. ./...

vet:
	go vet ./encode/...

regenerate:
	goderive ./...
	(cd parser && make regenerate)
	(cd encode && make regenerate)

clean:
	go clean ./...

nuke: clean
	(cd parser && make nuke)
	go clean -i ./...

gofmt:
	gofmt -l -s -w .

travis:
	make all
	make errcheck
	make checklicense
	make diff

errcheck:
	go get github.com/kisielk/errcheck
	errcheck -ignore 'fmt:[FS]?[Pp]rint*' ./...

diff:
	git diff --exit-code .
