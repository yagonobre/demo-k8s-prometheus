build-docker-alpine:
	docker build . -t yago/prometheus-demo

build-docker-scratch:
	env GOOS=linux go build -o app .
	docker build -f Dockerfile-scratch . -t yago/prometheus-demo:scratch

stress-test:
	ab -k -c 100 -n 10000 $(minikube service hello-go --url)/
