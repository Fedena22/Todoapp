FROM golang:1.19.0 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o todo

FROM scratch
COPY --from=builder /app/todo /todo

ENTRYPOINT ["/todo"]