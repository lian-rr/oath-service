FROM golang:1.8
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build model/*.go
RUN go build -o main *.go
CMD ["/go/src/app/main"]
