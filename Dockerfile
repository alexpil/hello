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

RUN GOOS=linux GOARCH=amd64 make build

# Run the binary
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
USER myapp

# certificates to interact with other services
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /hello/bin/hello /hello

ENV PORT 8080
ENV DIAG_PORT 8089
ENV DATABASE_URL postgres://user:pass@db/postgres?sslmode=disable

EXPOSE $PORT
EXPOSE $DIAG_PORT

CMD ["/hello"]
