FROM golang:alpine 

COPY hello.go /tmp 

CMD mkdir pictures

COPY pictures /go/pictures

CMD go run /tmp/hello.go 