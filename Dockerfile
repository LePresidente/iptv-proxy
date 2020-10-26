FROM golang:1.14.4-alpine

RUN apk add ca-certificates

# Timezone (TZ)
RUN apk update && apk add --no-cache tzdata
ENV TZ=America/Chicago
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /go/src/github.com/pierre-emmanuelJ/iptv-proxy
COPY . .
RUN GO111MODULE=off CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o iptv-proxy .

# Expose Port
EXPOSE 9000

FROM alpine:3.11.6
COPY --from=0  /go/src/github.com/pierre-emmanuelJ/iptv-proxy/iptv-proxy /
ENTRYPOINT ["/iptv-proxy"]
