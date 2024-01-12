FROM golang:1.20-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o hello-service ./main.go

FROM alpine:3.18.5
WORKDIR /app
COPY --from=builder /app/hello-service /app

CMD [ "./hello-service" ]