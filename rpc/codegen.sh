#!/bin/bash

rm -f ./protos.pb/*

protoc --proto_path=./protos/ --go_out=plugins=grpc,paths=source_relative:./protos.pb/ ./protos/*.proto