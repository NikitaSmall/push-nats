FROM golang:1.13-alpine AS build-env

RUN apk add --no-cache curl git openssl bzr \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && apk del curl

ENV G=/go/src/github.com/nikitasmall/push-nats/publisher
COPY . $G

RUN cd $G && dep ensure -vendor-only && go build -o /tmp/pub

# final stage
FROM alpine
COPY --from=build-env /tmp/pub /
ENTRYPOINT ["/pub"]
