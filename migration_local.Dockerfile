FROM alpine:3.13

RUN apk update
RUN apk upgrade
RUN apk add bash
RUN rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.14.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /root

ADD migrations/*.sql migrations/
ADD migration_local.sh .
ADD local.env .

RUN chmod +x migration_local.sh

ENTRYPOINT ["bash", "migration_local.sh"]