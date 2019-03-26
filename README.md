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

