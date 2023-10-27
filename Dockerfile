FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
