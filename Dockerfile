FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /build

COPY . .

RUN go build -o /bin/app ./cmd/server/main.go

FROM alpine:3.21.3

WORKDIR /app

COPY --from=builder /bin/app app

CMD [ "./app" ]

