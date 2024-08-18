FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /httpserv
FROM alpine:3.20.2
COPY --from=builder /httpserv /httpserv
EXPOSE 9989
CMD [ "/httpserv","-ms","-s:9989", "-chttp://localhost:9980"]