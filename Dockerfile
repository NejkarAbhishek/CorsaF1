FROM golang:1.22.3

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o server .

EXPOSE 8080

CMD ["./server"]
