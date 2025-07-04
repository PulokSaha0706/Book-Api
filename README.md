# BookApi

This project is a simple Go API server that can be run locally, containerized with Docker, and deployed on Kubernetes.

---

## Run Locally

To run the project locally:

```bash
go run main.go start -p 9090
```

## Build and Run Docker Image

Build the Docker image:

```bash
docker build -t puloksaha/bookapi:latest .
```
Run the Docker container exposing port 9090:

```bash
docker run -p 9090:9090 puloksaha/bookapi:latest start -p 9090
```


## Upload Docker Image to Docker Hub

Push your image to Docker Hub:

```bash
docker push puloksaha/bookapi:latest
```


## Deploy to Kubernetes

Apply the Deployment manifest:

```bash
kubectl apply -f bookapi-deployment.yaml
```
Apply the Service manifest

```bash
kubectl apply -f bookapi-server.yaml
```
Forward the service port to your local machine:
```bash
kubectl port-forward svc/bookapi-service 9090:9090
```

Now your API is accessible at http://localhost:9090.
