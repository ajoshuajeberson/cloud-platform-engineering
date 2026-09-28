# Cloud Platform Engineering

A cloud-native platform engineering project demonstrating how to build, containerize, deploy, observe, and operate a Go-based API on Kubernetes.

The project focuses on practical DevOps/SRE capabilities including Kubernetes deployment, application health checks, Prometheus metrics, Grafana observability, persistent monitoring, and automated alerting.

## Project Overview

The platform takes a Go REST API through the following workflow:

Application → Docker → Kubernetes → Prometheus → Grafana → Alerting

The project demonstrates an operational workflow rather than only application development.

## Architecture

```text
                    Developer
                        |
                        | git push
                        v
                 GitHub Repository
                        |
                        v
                  Go Application
                        |
                        v
                  Docker Image
                        |
                        v
                  Kubernetes
                cloud-platform
                        |
             +----------+----------+
             |                     |
             v                     v
       Go API Pods             Kubernetes Service
       (2 replicas)                  |
             |                       |
             | /metrics              |
             +-----------+-----------+
                         |
                         v
                    Prometheus
                         |
                         v
                      Grafana
                         |
             +-----------+-----------+
             |                       |
             v                       v
       SRE Dashboard            Alert Rules
                                     |
                                     v
                                SRE Email
