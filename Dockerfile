FROM --platform=$BUILDPLATFORM golang:1.21-alpine3.18 AS builder

ARG TARGETARCH

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /serve

FROM scratch

COPY --from=builder /serve /serve

CMD ["/serve"]