# build stage
FROM golang:latest AS build
WORKDIR /app

# speed up
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE=on

COPY . .
RUN go build -o ./server . && chmod +x ./server

# prod stage
FROM docker:dind AS prod
ENV Role=Judge
WORKDIR /app

COPY --from=build /app/server .
COPY . .

RUN chmod +x ./server
RUN mkdir -p /submit && mkdir -p /compile && mkdir -p /output && mkdir -p /case

ENTRYPOINT ["/app/server"]
