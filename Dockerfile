FROM golang:1.8

WORKDIR /go/src/praypi
COPY . .

WORKDIR /go/src/praypi/cmd/praypi
RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["praypi"]
