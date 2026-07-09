FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o program-search main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/program-search /app

CMD ["/app/program-search"]