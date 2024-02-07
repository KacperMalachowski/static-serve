FROM golang:1.22-alpine3.18 AS builder

RUN apk update && \
  apk add --no-cache git ca-certificates && \
  update-ca-certificates

ENV USER=appuser
ENV UID=10001 

RUN adduser \    
  --disabled-password \    
  --gecos "" \    
  --home "/nonexistent" \    
  --shell "/sbin/nologin" \    
  --no-create-home \    
  --uid "${UID}" \    
  "${USER}"

WORKDIR $GOPATH/src/KacperMalachowski/static-serve/

COPY . ./
RUN go mod download
RUN go mod verify


RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/serve

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /go/bin/serve /go/bin/serve

USER appuser:appuser
#checkov:skip=CKV_DOCKER_2: No need to setup Health Check
CMD ["/go/bin/serve"]