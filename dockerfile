FROM golang:alpine AS builder

WORKDIR $GOPATH/src/app/

COPY . .

RUN go mod tidy

RUN go mod download

WORKDIR $GOPATH/src/app/cmd/

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/mycatmarcat

FROM scratch

COPY --from=builder /go/bin/mycatmarcat /go/bin/mycatmarcat

ENTRYPOINT [ "/go/bin/mycatmarcat" ]

#