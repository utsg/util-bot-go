FROM golang:1.18 as builder

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go mod download

RUN go build -o /util-bot-go

EXPOSE 8080

CMD ["/util-bot-go"]
