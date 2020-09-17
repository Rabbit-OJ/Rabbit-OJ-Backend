FROM node:14-alpine

ENV Role="Tester"

WORKDIR /app
COPY ./tester /app/server
RUN chmod +x /app/server

RUN mkdir -p /result && mkdir -p /case && mkdir -p /output

ENTRYPOINT ["/app/server"]