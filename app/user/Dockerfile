FROM golang:1.24-alpine AS build

WORKDIR /usr/app

COPY app/user/go.mod .
COPY app/user/go.sum .
RUN go mod download

COPY app/user/ .

RUN go build -o /usr/app/user /usr/app/cmd/main.go

FROM alpine AS production

WORKDIR /usr/app

COPY --from=build /usr/app ./

CMD ["./user"]
