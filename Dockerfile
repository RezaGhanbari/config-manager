FROM golang:1.12 AS builder

COPY . /go/src/config-manager/
WORKDIR /go/src/config-manager/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
COPY --from=builder /go/src/config-manager//main /app/
WORKDIR /app
CMD ["./main"]
