FROM golang:alpine

ENV OTEL_SERVICE_NAME="go-example"
ENV OTEL_EXPORTER_OTLP_PROTOCOL=http/protobuf
ENV OTEL_EXPORTER_OTLP_ENDPOINT="https://api.honeycomb.io"
ENV OTEL_EXPORTER_OTLP_HEADERS="x-honeycomb-team=9P1O6c95bc44KB7n5bthjA"


RUN go install github.com/honeycombio/otel-config-go/otelconfig@latest 

ADD . .
RUN go build main.go

EXPOSE 8080

CMD ./main
