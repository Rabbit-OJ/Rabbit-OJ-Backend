FROM alpine:latest AS prod
WORKDIR /app

COPY ./tester /app/server

COPY ./files /app/files
COPY ./config.json /app/config.json
COPY ./files/certs /app/files/certs

EXPOSE 8090
EXPOSE 8888

RUN chmod +x ./server
RUN mkdir -p /submit && mkdir -p /compile && mkdir -p /output && mkdir -p /case && mkdir -p /result

ENTRYPOINT ["/app/server"]
