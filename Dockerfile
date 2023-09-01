FROM golang:1.20-alpine AS builder

RUN apk update
RUN apk add git
RUN apk add ca-certificates

WORKDIR /workspace
COPY . .

ENV GO111MODULE=on

RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o build stateManager/cmd/main.go

FROM alpine

COPY --from=builder /workspace/build ./main

EXPOSE 3002
COPY .config.yaml /
CMD ["./main"]