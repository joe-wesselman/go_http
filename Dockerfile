FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN apk add --no-cache make && make 

FROM alpine:3.21

RUN adduser -D appuser 
USER appuser 
WORKDIR /app 

COPY --from=builder /app/basic_server.exe .
COPY --from=builder /app/assets ./assets

EXPOSE 8000
ENTRYPOINT ["./basic_server.exe"]