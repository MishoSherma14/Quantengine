FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o worker ./cmd/worker

FROM gcr.io/distroless/static
COPY --from=builder /app/worker /worker

CMD ["/worker"]

