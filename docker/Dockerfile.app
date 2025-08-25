FROM golang:1.25 AS builder

WORKDIR /app

# copiar go.mod e go.sum primeiro (cache eficiente)
COPY go.mod ./
RUN go mod download

# copiar código
COPY . .

# build
RUN go build -o desculpaai ./cmd/desculpaai

# imagem final
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/desculpaai .
COPY web ./web

EXPOSE 8080
CMD ["./desculpaai"]

