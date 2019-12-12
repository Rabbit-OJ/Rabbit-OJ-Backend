FROM alpine:latest

ENV Role="Tester"

WORKDIR /app
COPY ./tester /app/server
RUN chmod +x /app/server

ENTRYPOINT ["/app/server"]