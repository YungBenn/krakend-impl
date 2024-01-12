FROM golang:1.20-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go build -o user-service ./main.go

FROM alpine:3.18.5
WORKDIR /app
COPY --from=builder /app/user-service /app

CMD [ "./user-service" ]