FROM golang:1.15-alpine AS build

ENV CGO_ENABLED 0

RUN apk add --no-cache ca-certificates curl && \
  apk add --no-cache --virtual .build-deps git

WORKDIR /src/
COPY main.go /src/
RUN go get -d -v && go build -o /bin/micro_prometheus_zendesk_exporter

FROM alpine
LABEL maintainer="Peter Savichev <psavichev@gmail.com>"

RUN apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=build /bin/micro_prometheus_zendesk_exporter /bin/micro_prometheus_zendesk_exporter

EXPOSE 9803
ENTRYPOINT ["/bin/micro_prometheus_zendesk_exporter"]
