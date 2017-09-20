FROM golang:1.9-alpine3.6

RUN mkdir -p /go/src/github.com/yagonobre/demo-k8s-prometheus
WORKDIR /go/src/github.com/yagonobre/demo-k8s-prometheus
COPY . /go/src/github.com/yagonobre/demo-k8s-prometheus

RUN go build -o app .

ENTRYPOINT ["./app"]
EXPOSE 8080
