# Cloud Platform Engineering

A cloud-native platform engineering project demonstrating how to build, containerize, deploy, observe, secure, and operate a Go-based REST API on Kubernetes.

The project focuses on practical DevOps/SRE capabilities including:

- Kubernetes deployment and service management
- Containerization with Docker
- Automated CI/CD with GitHub Actions
- Container vulnerability scanning with Trivy
- GitHub Container Registry (GHCR)
- Kubernetes health checks and resource management
- Prometheus application metrics
- Grafana dashboards and alerting
- Automated Kubernetes rollouts
- SRE-oriented observability and troubleshooting

---

## Project Architecture

```text
                              Developer
                                  |
                                  | git push
                                  v
                         GitHub Repository
                                  |
                                  v
                         GitHub Actions CI
                                  |
                 +----------------+----------------+
                 |                |                |
              Go Tests       Docker Build       Trivy
                                                    |
                                             Security Scan
                                                    |
                                                    v
                                             GitHub Container
                                                Registry
                                                    |
                                                    | image SHA
                                                    v
                                      Self-Hosted GitHub Runner
                                                    |
                                                  kubectl
                                                    |
                                                    v
                                             Kubernetes
                                              (Minikube)
                                                    |
                                  +-----------------+----------------+
                                  |                                  |
                                  v                                  v
                           Go API Pods                         Kubernetes Service
                           (2 replicas)                              |
                                  |                                   |
                                  +---------------+-------------------+
                                                  |
                                               /metrics
                                                  |
                                                  v
                                              Prometheus
                                                  |
                                                  v
                                               Grafana
                                                  |
                                  +---------------+---------------+
                                  |                               |
                                  v                               v
                           SRE Dashboard                    Alert Rules
                                                                  |
                                                                  v
                                                               Email
