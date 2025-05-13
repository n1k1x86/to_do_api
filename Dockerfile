FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o to_do ./cmd/
FROM alpine:latest

COPY --from=builder /app/to_do /app/to_do

WORKDIR /app
CMD [ "./to_do" ]