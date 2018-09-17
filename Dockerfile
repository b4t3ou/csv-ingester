FROM golang:1.11


RUN mkdir -p /go/src/github.com/b4t3ou/csv-ingester
WORKDIR /go/src/github.com/b4t3ou/csv-ingester
COPY . /go/src/github.com/b4t3ou/csv-ingester

RUN go install github.com/b4t3ou/csv-ingester

CMD ["/go/bin/csv-ingester"]
EXPOSE 3000