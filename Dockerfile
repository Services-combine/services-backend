# Build stage go
FROM golang:1.18-alpine3.17 AS builder-go

WORKDIR /app
COPY . .
RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

# Build stage python
FROM python:3.9 as builder

COPY requirements.txt .
RUN pip3 install --user -r requirements.txt

# Run stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder-go /app/.bin/app .
COPY --from=builder-go /app/configs configs/
COPY --from=builder-go /app/.env .
COPY --from=builder-go /app/accounts accounts/
COPY --from=builder-go /app/channels channels/
COPY --from=builder-go /app/python/account_telethon python/
COPY --from=builder /root/.local /root/.local

EXPOSE 8000
CMD ["./app"]