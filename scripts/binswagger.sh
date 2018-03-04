#!/bin/sh
cd frontend/public
go-bindata -pkg assets swagger/...
mv bindata.go ../../pkg/assets
gofmt -s -w ../../pkg/assets/bindata.go