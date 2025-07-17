
# IoT Telemetry Ingestion System

This project simulates and ingests high-frequency telemetry data from IoT devices (e.g., GPS data with IMEI, latitude, longitude, timestamp) and processes it via Kafka for downstream systems.

##  📡Project Structure

- telemetry-ingestor: Go service that consumes telemetry events, calculates cumulative distance per IMEI, and stores the results in Redis.
- iot-simulator: Simulates GPS events and sends them to the ingestor via HTTP POST requests.
- deployment: Contains all configuration files for deploying the entire stack using Docker Compose, Docker Swarm, or Kubernetes.

---

## 📦 Docker Images

Prebuilt Docker images are available on Docker Hub:

- `sshewalkar4094/telemetry-ingestor`
- `sshewalkar4094/iot-simulator`
---

## Prerequisites

Please choose your environment:

### ✅ Local Docker
- Use `docker-compose.yml`.
- Requires only Docker.

### ✅ Docker Swarm
- Use `docker-compose-swarm.yml`.
- Requires Docker and `docker swarm init`.

### ✅ Kubernetes
- Use `telemetry-manifests.yaml`.
- Requires Minikube, CRC (CodeReady Containers), or a Kubernetes cluster.

---

## 🚀 Deployment Options

### 🐳 Option 1: Docker Compose

```bash
cd deployment
docker-compose up --build
````

Create Kafka topic manually (optional if auto.create.topics is disabled):

```bash
docker exec -it <kafka-container-id> kafka-topics.sh --bootstrap-server localhost:9092 --create --topic telemetry --partitions 3
```

### 🐝 Option 2: Docker Swarm

```bash
cd deployment
docker swarm init
docker stack deploy -c docker-compose-swarm.yml telemetry_stack
# To remove
docker stack rm telemetry_stack
```

### ☸️ Option 3: Kubernetes (Minikube/CRC/Cluster)

```bash
kubectl create namespace telemetry-system
# or with oc
oc new-project telemetry-system

cd deployment
kubectl apply -f telemetry-manifests.yaml

# Port-forward to access API locally
oc port-forward -n telemetry-system pod/<telemetry-gateway-pod-name> 8080:8080
```

---
## Documentation
Swagger UI available at: http://localhost:8080/swagger/index.html.

---
## 📤 Sample Payload

**Endpoint:** `POST /v1/telemetrybatch`

```json
{
  "imei": "121212121212121",
  "events": [
    {
      "latitude": 18.5204,
      "longitude": 73.8567,
      "device_time": 1721202000000
    },
    {
      "latitude": 19.5210,
      "longitude": 85.8572,
      "device_time": 2021202060000
    }
  ]
}
```


**Endpoint:** `GET v1/devices/121212121212121`

Expected Output:

```json
{
  "imei": "121212121212121",
  "data": {
    "dist": "1266168.2113725634",
    "lat": "19.521",
    "lon": "85.8572"
  }
}
```