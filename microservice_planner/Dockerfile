FROM golang:1.18.2-alpine3.15 AS base_builder

WORKDIR /myapp/

COPY ["go.mod", "go.sum", "./"]

RUN apk add build-base && go mod download

FROM base_builder AS builder

WORKDIR /myapp/

COPY . .

RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-sql-driver/mysql
RUN go build -ldflags "-s -w" -o app cmd/main.go

FROM alpine:latest

WORKDIR /home/appuser/app/

COPY --from=builder ["/myapp/app", "myapp/env.json", "./"]

RUN echo "**** Install Dependencies ****" \
  && adduser --uid 1000 --gecos 'App User' --disabled-password appuser \
  && chown -R appuser.appuser /home/appuser/app/ && rm -Rf /bin/ls /bin/cp /bin/mv /bin/rm /bin/cat /bin/ln /bin/sed /bin/pwd

EXPOSE 8086

USER appuser

ENTRYPOINT ["./app"]