# Prometheus Lab

Instrument an application to expose Prometheus metrics and collect those metrics using Prometheus deployed with the Prometheus Operator.

## Prerequisites

- a Kubernetes cluster

## How To

Build a 0.1 version of go-server:

    $ cd go-server
    $ export IMAGE_REPO=[your image repo]
    $ export IMAGE_TAG=0.1
    $ make release

Edit `go-server.yaml`, add your image repo/tag and deploy:

    $ kubectl create ns go-server
    $ kubectl apply -f go-server.yaml

Deploy the crashcart to test your go server:

    $ kubectl apply -f ../crashcart-po.yaml
    $ kubectl exec -it -n go-server crashcart bash
    # curl go-server:8080

Expose some prometheus metrics on go-server:

    $ patch main.go expose_metrics.patch
    $ export IMAGE_TAG=0.2
    $ make release

Edit the `go-server.yaml` manifest and update the tag on the deployment's image, then update the running deployment:

    $ kubectl apply -f go-server.yaml

Test the update:

    $ kubectl exec -it -n go-server crashcart bash
    # curl go-server:8080
    # curl go-server:8080/metrics

Deploy prometheus operator:

    $ kubectl create namespace monitoring
    $ kubectl apply -f ../prometheus-operator.yaml

Create prometheus instance:

    $ kubectl apply -f go-server-prom.yaml
    $ kubectl port-forward -n go-server svc/go-server-prom 9090:9090

Create service monitor:

    $ kubectl apply -f go-server-svcmon.yaml
    $ kubectl port-forward -n go-server svc/go-server-prom 9090:9090

Add some metadata to metrics:

    $ patch main.go add_metadata.patch
    $ export IMAGE_TAG=0.3
    $ make release

Once again, edit the `go-server.yaml` manifest and update the tag on the deployment's image, then update the running deployment:

    $ kubectl apply -f go-server.yaml

Test the update:

    $ kubectl exec -it -n go-server crashcart bash
    # curl go-server:8080
    # curl go-server:8080/metrics

Create rules:

    $ promtool check rules go-server-rules-spec.yaml
    $ cat go-server-rules.yaml go-server-rules-spec.yaml | kubectl apply -f -

