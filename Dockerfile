# Download dependencies
FROM golang:1.13 AS modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

# Make a builder
FROM golang:1.13 AS builder

# add a non-privileged user
RUN useradd -u 10001 myapp

COPY --from=modules /go/pkg/mod /go/pkg/mod

RUN mkdir -p /hello
ADD . /hello
WORKDIR /hello

RUN CGO_ENABLED=0 go build -o bin/hello ./cmd/hello

# Run the binary
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
USER myapp

# certificates to interact with other services
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /hello/bin/hello /hello


ENV PORT 8080

EXPOSE $PORT

CMD ["/hello"]
