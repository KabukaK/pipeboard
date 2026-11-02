FROM golang:1.27-alpine AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /bin/api /bin/api

EXPOSE 8080
CMD ["/bin/api"]