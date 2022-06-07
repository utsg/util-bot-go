FROM golang:1.18-bullseye as builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o ./util-bot-go

FROM scratch
WORKDIR /
COPY --from=builder  /app/util-bot-go /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080

CMD ["/util-bot-go"]
