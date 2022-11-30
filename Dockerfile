FROM golang:1.19.3 AS builder

WORKDIR /app
COPY / .
RUN CGO_ENABLED=0 go build -o /my-test-app .

FROM scratch
WORKDIR /
COPY --from=builder /my-test-app .
CMD [ "/my-test-app" ]
