This folder contains the solution for Problem 3
# 🐳 Problem 3 – GoLang App Deployment on Kubernetes

This folder contains the solution for Problem Statement 3 of the Product Manager assignment.

## 📌 Task Overview
Create a GoLang web app that displays the current date and time. Containerize it with Docker and deploy it to Kubernetes using a declarative approach (YAML files).

## 📄 Included Files
- `main.go` – GoLang code for the date/time server
- `Dockerfile` – Docker instructions to containerize the app
- `deployment.yaml` – Kubernetes Deployment with 2 replicas
- `service.yaml` – Kubernetes Service to expose the app

## 🔧 Tools Used
- GoLang
- Docker
- Kubernetes (Minikube, Play with K8s, or GCP)

## ✅ Goal
Deploy a minimal GoLang app to a K8s cluster and expose it to the internet using a declarative setup.
