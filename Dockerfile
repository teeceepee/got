# builder image
FROM golang:1.11-alpine AS builder

WORKDIR /got-build/
COPY *.go .
RUN go build -o got-bin


# release image
FROM alpine:3.8 AS release

WORKDIR /gooooot
COPY --from=builder /got-build/got-bin .
CMD ["./got-bin"]
