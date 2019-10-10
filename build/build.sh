#!/bin/sh

project=gts
export GOOS=linux
export GOARCH=amd64
export BASEDIR=$(pwd)

echo "============================="
echo "==== building"
echo "============================="
cd "$BASEDIR"
go build -ldflags "-s -w" -o ${project}

if [ $? -ne 0 ]
then
    echo "build failed"
    exit -1
fi

echo "============================="
echo "==== packaging"
echo "============================="
cd "$BASEDIR"
tar -czf ${project}.tar.gz ${project} configs static

cd $BASEDIR
rm -rf dist
mkdir -p dist
mv $BASEDIR/${project}.tar.gz dist/
cd dist
tar -xvf ${project}.tar.gz
rm -f ${project}.tar.gz

echo "============================="
echo "==== clean"
echo "============================="
rm -rf "$BASEDIR/${project}"
