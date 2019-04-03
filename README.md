# Prometheus Lab

Prometheus operator stuff.

## Prerequisites

- a Kubernetes cluster

## How To

Deploy prometheus operator

    $ kubectl create namespace monitoring
    $ kubectl apply -f prometheus-operator.yaml

Deploy go-server

    $ kubectl create ns go-server
    $ kubectl apply -f go-server/go-server.yaml

Create prometheus instance

    $ kubectl apply -f go-server/go-server-prom.yaml
    $ kubectl port-forward -n go-server svc/go-server-prom 9090:9090

Create service monitor

    $ kubectl apply -f go-server/go-server-svcmon.yaml
    $ kubectl port-forward -n go-server svc/go-server-prom 9090:9090

