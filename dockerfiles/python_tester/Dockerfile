FROM python:3

ENV Role="Tester"

WORKDIR /app
COPY ./tester /app/server
RUN chmod +x /app/server

ENTRYPOINT ["/app/server"]