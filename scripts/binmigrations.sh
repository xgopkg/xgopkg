#!/bin/sh
cd db/migrations
go-bindata -pkg migrate .
mv bindata.go ../../pkg/migrate/
gofmt -s -w ../../pkg/migrate/bindata.go