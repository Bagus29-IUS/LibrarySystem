FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o main .

EXPOSE 8080

CMD ["/app/main"]

