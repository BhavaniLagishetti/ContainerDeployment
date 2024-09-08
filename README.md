# Key Server Application



## Features

- Exposes an endpoint (/key) to generate random keys with the specified length
- Prometheus metrics are exposed at the /metrics endpoint
- Dynamic configuration via command-line arguments for server port and maximum key size
- Unit tests included for both positive and negative cases
- Packaged into a lightweight Docker container
- Deployable to Kubernetes using Helm

## Usage

### Command Line Arguments

- --max-size (default: 1024): Defines the maximum length of the key
- --srv-port (default: 1123): Specifies the port the server will listen on

### Example HTTP Request


GET http://localhost:1123/key?length=512

You can replace localhost with the IP address of the server where you are running.

This request generates a random key of 512 bytes, assuming the maximum key size is 1024 bytes.

### Prometheus Metrics

- Key length distribution histogram (with 20 linear buckets)
- Counter for HTTP status codes

## Unit Tests

To run unit tests, use the following command:


go test ./...


## Building and Running

1. Clone the repository:
   
   git clone 
   cd key-server
   

2. Build the Go binary:
   
   go build -o key-server .
   

3. Run the application:
   
   ./key-server --max-size 1024 --srv-port 1123
   

4. Access the key generation endpoint:
   
   curl "http://localhost:1123/key?length=100"
   

5. Access the Prometheus metrics:
   
   curl "http://localhost:1123/metrics"
   

## Docker

### Build the Docker Image


docker build -t your-docker-repo/key-server .


### Run the Docker Container


docker run -p 1123:1123 your-docker-repo/key-server --max-size 1024 --srv-port 1123


## Kubernetes

### Helm Deployment

You can deploy the application to Kubernetes using the provided Helm chart.

1. Build and push the Docker image to your Docker registry.
2. Install the Helm chart:


helm install key-server ./key-server-chart


Note: You can replace localhost with the IP address of the server where you are running.