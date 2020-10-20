FROM golang:1.13.5-alpine3.11 AS builder
WORKDIR /opt/simple-https-listener
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GO111MODULE=on
COPY go.mod .
COPY main.go .
RUN go build -ldflags="-w -s" -o simple-https-listener .

FROM scratch
COPY --from=builder /opt/simple-https-listener/simple-https-listener /simple-https-listener
ENTRYPOINT ["/simple-https-listener"]