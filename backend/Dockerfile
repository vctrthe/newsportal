FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o newsportal-backend .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/newsportal-backend .

COPY ./docs /app/docs

COPY .env /app/.env

EXPOSE 8081

CMD ["./newsportal-backend"]