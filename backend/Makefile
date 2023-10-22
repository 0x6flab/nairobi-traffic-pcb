NTP_DOCKER_IMAGE_NAME_PREFIX ?= ghcr.io/0x6flab/nairobi-traffic-pcb
BUILD_DIR = build
DM_SERVICE = device-manager
SERVICES = $(DM_SERVICE)
CGO_ENABLED ?= 0
GOARCH ?= amd64
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.1.0")
COMMIT ?= $(shell git rev-parse HEAD)
TIME ?= $(shell date +%F_%T)

define compile_dm_service
	cd $(DM_SERVICE) && \
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) \
	go build -mod=vendor -ldflags "-s -w \
	-X 'github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager.BuildTime=$(TIME)' \
	-X 'github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager.Version=$(VERSION)' \
	-X 'github.com/0x6flab/nairobi-traffic-pcb/backend/device-manager.Commit=$(COMMIT)'" \
	-o ../${BUILD_DIR}/ntp-$(DM_SERVICE) cmd/main.go
endef

define make_dm_docker
	cd $(DM_SERVICE) && \
	docker build \
		--no-cache \
		--build-arg SVC=$(DM_SERVICE) \
		--tag=$(NTP_DOCKER_IMAGE_NAME_PREFIX)/$(DM_SERVICE):latest \
		--tag=$(NTP_DOCKER_IMAGE_NAME_PREFIX)/$(DM_SERVICE):$(VERSION) \
		-f docker/Dockerfile ../build
endef

define docker_push
	for svc in $(SERVICES); do \
		docker push $(NTP_DOCKER_IMAGE_NAME_PREFIX)/$$svc:$(VERSION); \
		docker push $(NTP_DOCKER_IMAGE_NAME_PREFIX)/$$svc:latest; \
	done
endef

$(DM_SERVICE):
	$(call compile_dm_service)

all: $(DM_SERVICE)

.PHONY: all $(SERVICES) dockers dockers_dev

dockers: $(DM_SERVICE)
	$(call make_dm_docker)

docker_push:
	$(call docker_push)

clean:
	rm -rf ${BUILD_DIR}

install:
	cp ${BUILD_DIR}/* $(GOBIN)

lint:
	golangci-lint run  --timeout 3m

version:
	standard-version
	goreleaser release --clean --release-notes CHANGELOG.md
	git push --follow-tags origin main

run:
	docker-compose -f docker/docker-compose.yaml --env-file ../.env up

logs:
	docker-compose -f docker/docker-compose.yaml --env-file ../.env logs -f

stop:
	docker-compose -f docker/docker-compose.yaml --env-file ../.env down -v
