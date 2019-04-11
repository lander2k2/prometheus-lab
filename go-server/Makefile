IMAGE_REPO?=quay.io/myrepo/go-server
IMAGE_TAG?=0.1

build:
	GOOS=linux go build -o foster-server main.go

image:
	docker build -t $(IMAGE_REPO):$(IMAGE_TAG) .

push:
	docker push $(IMAGE_REPO):$(IMAGE_TAG)

clean:
	rm foster-server

release: build image push clean

