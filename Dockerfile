FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o storage .

FROM scratch

COPY --from=builder ["/app/storage", "/"]

EXPOSE 4009

ENTRYPOINT ["/storage"]