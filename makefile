OWNER=inloopx
IMAGE_NAME=sql-exporter
BINARY_NAME=sql-exporter
QNAME=$(OWNER)/$(IMAGE_NAME)

GIT_TAG=$(QNAME):$(TRAVIS_COMMIT)
BUILD_TAG=$(QNAME):0.1.$(TRAVIS_BUILD_NUMBER)
LATEST_TAG=$(QNAME):latest

lint:
	docker run -it --rm -v "$(PWD)/Dockerfile:/Dockerfile:ro" redcoolbeans/dockerlint

build:
	go get ./...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/binary-alpine
	docker build -t $(GIT_TAG) .

tag:
	docker tag $(GIT_TAG) $(BUILD_TAG)
	docker tag $(GIT_TAG) $(LATEST_TAG)
	docker tag $(GIT_TAG) $(OWNER)/$(IMAGE_NAME)

login:
	@docker login -u "$(DOCKER_USER)" -p "$(DOCKER_PASS)"
push: login
	# docker push $(GIT_TAG)
	# docker push $(BUILD_TAG)
	docker push $(LATEST_TAG)
	docker push $(OWNER)/$(IMAGE_NAME)


build-local:
	go get ./...
	go build -o $(BINARY_NAME)

deploy-local:
	make build-local
	mv $(BINARY_NAME) /usr/local/bin/
