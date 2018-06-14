FROM alpine:3.5

COPY bin/binary-alpine /sql-exporter

# https://serverfault.com/questions/772227/chmod-not-working-correctly-in-docker
RUN mv /sql-exporter /usr/local/bin/sql-exporter && \
    chmod +x /usr/local/bin/sql-exporter

ENTRYPOINT []
