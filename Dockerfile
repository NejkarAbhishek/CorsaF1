FROM golang:1.22.3

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o server ./cmd/main.go

EXPOSE 8080

CMD ["./server"]
