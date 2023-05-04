FROM golang:alpine AS builder
RUN apk update
RUN apk add upx

WORKDIR $GOPATH/src/app/

COPY . .

RUN go mod tidy

RUN go mod download

WORKDIR $GOPATH/src/app/cmd/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/mycatmarcat
RUN upx --best --lzma /go/bin/mycatmarcat

FROM scratch

COPY --from=builder /go/bin/mycatmarcat /go/bin/mycatmarcat

ENTRYPOINT [ "/go/bin/mycatmarcat" ]

#