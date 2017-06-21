# AlpineLinux with a glibc-2.23 and Oracle Java 8
FROM alpine:3.4

MAINTAINER Jeroen

add target/linux_amd64/happy-endpoint /usr/local/bin/happy-endpoint

CMD happy-endpoint