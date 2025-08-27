FROM golang:1.25 AS builder
WORKDIR /app

COPY go.mod ./

COPY desculpa-ai ./desculpa-ai

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app/desculpaai ./desculpa-ai/cmd

FROM gcr.io/distroless/base-debian12 AS runner
WORKDIR /app

COPY --from=builder /app/desculpaai .
COPY --from=builder /app/desculpa-ai/web ./web

EXPOSE 8080
CMD ["./desculpaai"]

