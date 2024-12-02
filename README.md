# Go Notes Application

A modern, cloud-native Notes Application built with Go, featuring a complete infrastructure setup with Terraform and Kubernetes deployment configuration.

## 🚀 Features

- RESTful API for note management
- Containerized application with Docker
- Infrastructure as Code with Terraform
- Kubernetes deployment with auto-scaling
- Load balanced with AWS ALB
- Secure configuration management

## 🏗️ Architecture

### Backend (Go)
- Clean architecture design
- RESTful API endpoints
- Middleware for security and logging
- Containerized with minimal Docker image

### Infrastructure (Terraform)
- EKS cluster setup
- VPC with public/private subnets
- Security groups and IAM roles
- Jenkins server for CI/CD

### Kubernetes
- Namespace isolation
- Horizontal Pod Autoscaling
- Rolling updates
- Health checks and monitoring
- AWS Load Balancer integration

## 📁 Project Structure
```
.
├── go-backend/          # Go application source code
│   ├── cmd/            # Application entrypoints
│   ├── internal/       # Private application code
│   ├── Dockerfile      # Container configuration
│   └── go.mod          # Go modules file
│
├── k8s/                # Kubernetes manifests
│   ├── 00-namespace.yaml
│   ├── 01-configmap.yaml
│   ├── 02-deployment.yaml
│   ├── 03-service.yaml
│   ├── 04-hpa.yaml
│   ├── 05-ingress.yaml
│   └── kustomization.yaml
│
└── Terraform/          # Infrastructure as Code
    ├── modules/        # Terraform modules
    │   ├── eks/       # EKS cluster configuration
    │   ├── vpc/       # Network configuration
    │   └── jenkins/   # Jenkins server setup
    ├── main.tf
    └── variables.tf
```

## 🛠️ Setup and Installation

### Prerequisites
- Go 1.21 or later
- Docker
- kubectl
- Terraform
- AWS CLI configured
- Access to an AWS account

### Local Development
1. Clone the repository:
   ```bash
   git clone https://github.com/AhmedZeins/Go_NoteApp.git
   cd Go_NotesApp
   ```

2. Run the Go application locally:
   ```bash
   cd go-backend
   go mod download
   go run cmd/server/main.go
   ```

3. Build Docker image:
   ```bash
   docker build -t notes-app:latest .
   ```

### Infrastructure Deployment

1. Initialize Terraform:
   ```bash
   cd Terraform
   terraform init
   ```

2. Deploy infrastructure:
   ```bash
   terraform plan
   terraform apply
   ```

3. Configure kubectl:
   ```bash
   aws eks update-kubeconfig --name your-cluster-name --region your-region
   ```

### Kubernetes Deployment

1. Deploy the application:
   ```bash
   kubectl apply -k k8s/
   ```

2. Verify deployment:
   ```bash
   kubectl -n notes-app get pods
   kubectl -n notes-app get svc
   kubectl -n notes-app get ingress
   ```

## 🔒 Security Features

- Non-root container execution
- Network policy isolation
- Resource limits and requests
- Secure configuration management
- TLS termination at ALB

## 🔍 Monitoring and Scaling

- Horizontal Pod Autoscaling based on CPU/Memory
- Readiness and liveness probes
- Prometheus-ready metrics
- Rolling update strategy

## 🧪 Testing

Run the tests:
```bash
cd go-backend
go test ./...
```
