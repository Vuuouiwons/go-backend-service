FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /build

COPY /src .

RUN go build -o /bin/app ./main.go

FROM alpine:3.21.3

WORKDIR /app

COPY --from=builder /bin/app app

CMD [ "./app" ]

