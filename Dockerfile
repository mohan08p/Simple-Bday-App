FROM golang:alpine AS builder

WORKDIR /go/src/app
COPY /src/Simple-Bday-App/ .

RUN go build -o simple-bday-app .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/ /app/

CMD ["./simple-bday-app"]
