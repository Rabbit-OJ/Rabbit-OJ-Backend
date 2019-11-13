FROM alpine:latest AS prod
ENV Role=Tester

WORKDIR /app

COPY ./tester ./tester

RUN chmod +x ./tester
RUN mkdir -p /submit && mkdir -p /compile && mkdir -p /output && mkdir -p /case

ENTRYPOINT ["/app/tester"]
