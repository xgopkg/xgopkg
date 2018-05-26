#!/bin/sh
for i in "$@"
do
case $i in
	-v=*|--version=*)
	VERSION="${i#*=}"
	shift
	;;
esac
done
if [ ! "$VERSION" -o "$VERSION" == " " ]; then
        echo "Version number is null or space. Please input version number, just like 0.0.1"
        exit 0
fi

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o xgopkg ./../cmd/xgopkg
docker build -t xgopkg/xgopkg:"$VERSION" .
docker push xgopkg/xgopkg:"$VERSION"