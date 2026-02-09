FROM golang:1.25-alpine3.23 AS builder

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

#RUN groupadd -r appuser && useradd -r -g appuser -m appuser

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# runtime stage
FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/public ./public

EXPOSE 8080
CMD ["./app"]


