#! /bin/bash

path="/Users/wuyajun/docker/deploy"
proto_path=${path}/proto/
protoc --go_out=${path}/protogo/ *.proto
