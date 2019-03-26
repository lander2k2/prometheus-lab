FROM golang:1.12.1-alpine3.9 as builder
COPY main.go .
RUN mkdir -p /go/src
COPY vendor /go/src/
RUN go build -o /server main.go

FROM alpine:3.9
CMD ["./server"]
COPY --from=builder /server .
