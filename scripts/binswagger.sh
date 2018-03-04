#!/bin/sh
cd frontend/public
go-bindata -pkg assets swagger/...
mv bindata.go ../../pkg/assets