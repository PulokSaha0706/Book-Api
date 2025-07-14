run:
	go run main.go start -p 9090
dockerRun:
	docker build -t puloksaha/bookapi:latest .
	docker run -p 9090:9090 puloksaha/bookapi:latest start -p 9090
yamlRun:
	kubectl apply -f bookapi-deployment.yaml
	kubectl apply -f bookapi-server.yaml
	kubectl port-forward svc/bookapi-service 9090:9090
