FROM guoxf/golang-build:1.17.0-alpine3.14 as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

WORKDIR /workspace
COPY . .
RUN go env && go build -gcflags=-G=3 -o telegraf-sidecar main.go

FROM guoxf/golang-run:alpine-3.13.5
LABEL MAINTAINER="eojan_support@126.com"

WORKDIR /workspace
COPY --from=builder /workspace/telegraf-sidecar ./
COPY --from=builder /workspace/docs/swagger.json ./docs
COPY --from=builder /workspace/docs/swagger.yaml ./docs

EXPOSE 48087

ENTRYPOINT ./telegraf-sidecar
