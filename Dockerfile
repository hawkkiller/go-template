FROM golang:1.18-alpine3.14 AS builder

WORKDIR /usr/local/go/src/
COPY app/ .

RUN go mod download
RUN go build -o app cmd/main/app.go

FROM alpine:3.14
# copy binary app
COPY --from=builder /usr/local/go/src/ /

CMD ["/app"]
