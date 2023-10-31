FROM ubuntu:latest
LABEL authors="tylermorton"

ENTRYPOINT ["top", "-b"]