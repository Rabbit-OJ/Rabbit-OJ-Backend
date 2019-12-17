FROM swift:5.1

ENV Role="Tester"

WORKDIR /app
COPY ./tester /app/server
RUN chmod +x /app/server

ENTRYPOINT ["/app/server"]