FROM golang:1.24-alpine AS build

WORKDIR /usr/app

COPY app/payment/go.mod .
COPY app/payment/go.sum .
RUN go mod download

COPY app/payment/ .

RUN go build -o /usr/app/payment /usr/app/cmd/main.go

FROM alpine AS production

WORKDIR /usr/app

COPY --from=build /usr/app ./

CMD ["./payment"]
