# build stage
FROM golang:latest AS build
WORKDIR /app

# speed up
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY . .
RUN go build -o ./server . && chmod +x ./server

# prod stage
FROM docker:dind AS prod
WORKDIR /app

COPY --from=build /app/server /app/server
COPY --from=build /app/files /app/files

RUN chmod +x ./server
RUN mkdir -p /submit && mkdir -p /compile && mkdir -p /output && mkdir -p /case && mkdir -p /result
RUN echo "[]" >> /app/files/storage.json

ENTRYPOINT ["/app/server"]
