FROM golang:1.20 AS builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux  go build -o main main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /go/src/app/main ./
EXPOSE 8080

CMD ["./main"]
