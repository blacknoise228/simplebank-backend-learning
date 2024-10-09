# Build stage
FROM golang:1.22.8-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.20 AS runner
WORKDIR /app
COPY --from=builder /app/app.env .
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]