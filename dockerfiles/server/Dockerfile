# build stage
FROM golang:latest AS build
WORKDIR /app

# speed up
ENV GOPROXY="https://goproxy.io"
ENV GO111MODULE=on

COPY . .
RUN go build -o ./server . && chmod +x ./server

# prod stage
FROM alpine:latest AS prod
ENV Role=Server
WORKDIR /app

COPY --from=build /app/server .
COPY --from=build /app/files .
COPY --from=build /app/statics .

RUN chmod +x ./server

ENTRYPOINT ["/app/server"]
