FROM golang:1.22.1 as builder
WORKDIR /app
COPY . .
RUN go build ./src/main/main.go

FROM alpine:3.19.1
COPY --from=builder /app/main /app
ENTRYPOINT [ "./app" ]