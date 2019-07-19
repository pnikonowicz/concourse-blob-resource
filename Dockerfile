FROM pivnet/golang

ADD . /go/src/concourse-blob-resource

WORKDIR /go/src

RUN go build -o concourse-blob-resource/src/main/check/check concourse-blob-resource/src/main/check/main.go
RUN go build -o concourse-blob-resource/src/main/in/in concourse-blob-resource/src/main/in/main.go
RUN go build -o concourse-blob-resource/src/main/out/out concourse-blob-resource/src/main/out/main.go

RUN mkdir -p /opt/resource/
RUN mv concourse-blob-resource/src/main/check/check /opt/resource/
RUN mv concourse-blob-resource/src/main/in/in /opt/resource/
RUN mv concourse-blob-resource/src/main/out/out /opt/resource/