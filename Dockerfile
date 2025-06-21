FROM golang:1.23rc1

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o server ./cmd/main.go

EXPOSE 8080

CMD ["./server"]

ENV DB_USER=user
ENV DB_PASS=pass
ENV DB_NAME=f1