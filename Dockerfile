# Build stage
FROM golang:1.19-alpine3.17 AS builder
WORKDIR /app/postr
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.17
WORKDIR /app/postr
COPY --from=builder /app/postr/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/postr/main" ]