FROM golang:1.24-alpine AS build

WORKDIR /usr/app

COPY app/gateway/go.mod .
COPY app/gateway/go.sum .
RUN go mod download

COPY app/gateway/ .

RUN go build -o /usr/app/gateway /usr/app/cmd/main.go

FROM alpine AS production

WORKDIR /usr/app

COPY --from=build /usr/app ./

CMD ["./gateway"]
