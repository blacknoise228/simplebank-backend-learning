# Build stage
FROM golang:1.22.8-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L "https://github.com/pressly/goose/releases/download/v3.22.1/goose_linux_x86_64" --output goose
RUN chmod 755 goose

# Run stage
FROM alpine:3.20 AS runner
WORKDIR /app
COPY --from=builder /app/app.env .
COPY --from=builder /app/main .
COPY --from=builder /app/goose .
COPY /db/migration ./migration
COPY start.sh .
COPY wait-for.sh .

EXPOSE 8080
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]