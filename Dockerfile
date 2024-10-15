# FROM golang:1.22-alpine AS build
# ARG WORK_DIR=/root/app
# ENV GOPATH=/root/app/gopath
# WORKDIR $WORK_DIR
# COPY . .
# ARG BUILD_ID
# RUN --mount=type=cache,id=$BUILD_ID-gopath,target=$GOPATH go build -o /app .

FROM debian:12-slim
RUN apt update && apt install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY dailydo /app
CMD ["/app"]

