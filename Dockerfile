FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod /app/
RUN go mod download
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go
CMD ["./main"]
