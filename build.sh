#!/bin/sh

make release

docker build -t happy-endpoint .