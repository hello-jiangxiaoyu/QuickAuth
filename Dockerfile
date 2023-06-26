##
## bulid web
##
FROM node:19-alpine AS build-web
WORKDIR /app
COPY ./web /app

RUN npm run install
RUN npm run export


##
## bulid backend
##
FROM golang:1.20-alpine as build-back
WORKDIR /app
COPY . .

RUN go env -w GO111MODULE=on \
        && go env -w GOPROXY=https://goproxy.cn,direct \
        && go env -w CGO_ENABLED=0

RUN go build -o QuickAuth .


##
## deploy
##
FROM alpine:latest
RUN mkdir -p /app/web
WORKDIR /app

COPY deploy/dev.yaml ./system.yaml

COPY --from=build-web    /app/.next/static       ./web/_next/static
COPY --from=build-web    /app/.next/server/pages ./web

COPY --from=build-back   /app/QuickAuth    ./
COPY --from=build-back   /app/system.yaml  ./

EXPOSE 80
ENTRYPOINT ["/app/QuickAuth"]
