FROM node:13-alpine

ENV Role="Tester"

WORKDIR /app
COPY ./tester /app/server
RUN chmod +x /app/server

ENTRYPOINT ["/app/server"]