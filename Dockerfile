FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download
WORKDIR /build/cmd
RUN go build -o main .

FROM alpine:latest  

WORKDIR /small
COPY --from=builder /build/cmd/main ./binary
COPY --from=builder /build/.env .

EXPOSE 8080

ENTRYPOINT [ "./binary" ]