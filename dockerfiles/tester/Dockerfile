# build stage
FROM golang:latest AS build
WORKDIR /app

# speed up
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE=on

COPY . .
RUN go build -o ./tester . && chmod +x ./tester

# prod stage
FROM alpine:latest AS prod
ENV Role=Tester

WORKDIR /app

COPY --from=build /app/tester ./tester

RUN chmod +x ./tester
RUN mkdir -p /submit && mkdir -p /compile && mkdir -p /output && mkdir -p /case

ENTRYPOINT ["/app/tester"]
