FROM golang:1.22-alpine AS build
ENV GOPATH=/root/app/gopath
WORKDIR /root/app
COPY . .
ARG APP_NAME
RUN --mount=type=cache,id=$APP_NAME-gopath,target=$GOPATH go build -o /app .

FROM debian:12-slim
RUN apt update && apt install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=build app /app
CMD ["/app"]

