FROM golang:1.17.5 AS builder

WORKDIR /go/src/github.com/Armatorix/payment-gateway

COPY go.mod go.sum ./

RUN go mod download
COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./

FROM scratch

WORKDIR /app
COPY --from=builder /go/src/github.com/Armatorix/payment-gateway/main /app/

EXPOSE 8080

ENTRYPOINT [ "./main" ]