# Start by building the application.
FROM golang:alpine as builder

WORKDIR /go/src/app
ADD . /go/src/app

RUN go mod tidy

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app


# Now copy it into our base image.
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]
