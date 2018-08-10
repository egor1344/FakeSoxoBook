FROM golang:1.9

WORKDIR /go/src/fakesoxobook
COPY . .

RUN go get -d -v ./...
RUN go build

CMD ./fakesoxobook
