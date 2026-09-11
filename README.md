# Cloud Platform Engineering

A production-style cloud platform engineering project demonstrating how to build, containerize, test, deploy, and automate a Go-based application using modern DevOps and platform engineering practices.

## Project Overview

This project starts with a simple Go REST API and progressively builds a complete cloud-native delivery platform around it.

The goal is to demonstrate an end-to-end workflow:

Application → Docker → Kubernetes → CI/CD → Infrastructure as Code → Cloud

## Architecture

```text
Developer
   |
   | git push
   v
GitHub Repository
   |
   v
CI/CD Pipeline
   |
   +---- Run tests
   |
   +---- Build application
   |
   +---- Build Docker image
   |
   +---- Security checks
   |
   v
Container Registry
   |
   v
Kubernetes
   |
   +---- Go API Pods
   |
   +---- Service
   |
   +---- Health Checks
   |
   v
Application
