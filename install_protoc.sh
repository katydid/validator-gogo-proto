#!/usr/bin/env bash

set -ex

cd ${HOME}

basename=protoc-29.3-linux-x86_64
wget https://github.com/protocolbuffers/protobuf/releases/download/v29.3/$basename.zip
unzip $basename.zip
