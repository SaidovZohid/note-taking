FROM golang:1.19.1-alpine3.16 as builder

WORKDIR /note

COPY . .

RUN apk add curl
RUN go build -o main cmd/main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz


FROM alpine:3.16

WORKDIR /note
RUN mkdir media

COPY --from=builder /note/main .
COPY --from=builder /note/migrate ./migrate
COPY migrations ./migrations
COPY templates ./templates

EXPOSE 8080

CMD [ "/note/main" ]

