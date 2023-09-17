FROM golang:1.21.1-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/ssm ./cmd/ssm

FROM alpine:3.14.2

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=0 /app/bin/ssm /app/bin/ssm

ENTRYPOINT ["/app/bin/ssm"]

CMD ["--help"]