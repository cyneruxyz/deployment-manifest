# builder container
FROM golang:1.19-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN go build -v -o server

# runner container
FROM alpine:latest

COPY --from=builder /app/server /app/server

CMD ["/app/server"]
