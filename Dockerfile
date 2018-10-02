FROM golang:1.11.0-stretch as builder

WORKDIR /data-service

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.8

WORKDIR /root/

COPY --from=builder /data-service .

ENTRYPOINT ./data-service

EXPOSE 80