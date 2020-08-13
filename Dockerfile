FROM golang:1.14 as builder
COPY . /app
WORKDIR /app
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/*
#second stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app .
CMD ["./app"]