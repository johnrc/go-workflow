FROM docker.io/library/golang:1.15.15 as builder
WORKDIR /go/src/app
COPY . .
RUN go build -o main main.go

FROM docker.io/library/busybox:1.34.1-glibc
COPY --from=builder /go/src/app/main /main
CMD ["/main"]
