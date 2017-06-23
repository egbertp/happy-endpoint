#!/bin/sh

make release

VERSION_TAG=`git describe 2>/dev/null | cut -f 1 -d '-' 2>/dev/null`

docker build -t happy-endpoint:${VERSION_TAG} .