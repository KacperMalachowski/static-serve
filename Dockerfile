FROM golang:1.20-alpine3.18 AS builder

ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /serve

FROM scratch

COPY --from=builder /serve /serve

CMD ["/serve"]