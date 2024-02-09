FROM golang:1.21-alpine

WORKDIR /DebTour-backend

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]

