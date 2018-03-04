#!/bin/sh
cd db
go-bindata -pkg migrate ./...
mv bindata.go ../pkg/migrate/
gofmt -s -w ../pkg/migrate/bindata.go