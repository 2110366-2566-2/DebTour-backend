FROM golang:1.21-alpine

WORKDIR /DebTour-backend

COPY . .

CMD ["./main"]

