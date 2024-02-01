FROM golang:1.21.6-alpine3.19 as build

WORKDIR /go/src

COPY . .

RUN go build -o /go/bin/find-my-ip

FROM golang:1.21.6-alpine3.19 as prod

RUN adduser -S non-root

COPY --from=build /go/bin/find-my-ip /usr/local/bin/find-my-ip

USER non-root

EXPOSE 80

ENTRYPOINT ["find-my-ip"]