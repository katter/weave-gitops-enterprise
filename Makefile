.PHONY: all install clean images lint unit-tests check
.DEFAULT_GOAL := all

# Boiler plate for bulding Docker containers.
# All this must go at top of file I'm afraid.
IMAGE_PREFIX := docker.io/weaveworks/
IMAGE_TAG := $(shell tools/image-tag)
GIT_REVISION := $(shell git rev-parse HEAD)
VERSION=$(shell git describe --always)
CURRENT_DIR := $(shell pwd)
UPTODATE := .uptodate
# The GOOS to use for local binaries that we `make install`
LOCAL_BINARIES_GOOS ?= $(GOOS)

# Every directory with a Dockerfile in it builds an image called
# $(IMAGE_PREFIX)<dirname>. Dependencies (i.e. things that go in the image)
# still need to be explicitly declared.
%/$(UPTODATE): %/Dockerfile %/*
	$(SUDO) docker build --build-arg=revision=$(GIT_REVISION) -t $(IMAGE_PREFIX)$(shell basename $(@D)) $(@D)/
	$(SUDO) docker tag $(IMAGE_PREFIX)$(shell basename $(@D)) $(IMAGE_PREFIX)$(shell basename $(@D)):$(IMAGE_TAG)
	touch $@

# Get a list of directories containing Dockerfiles
DOCKERFILES := $(shell find . -name tools -prune -o -name vendor -prune -o -name rpm -prune -o -name build -prune -o -name environments -prune -o -name test -prune -o -name examples -prune -o -type f -name 'Dockerfile' -print)
UPTODATE_FILES := $(patsubst %/Dockerfile,%/$(UPTODATE),$(DOCKERFILES))
DOCKER_IMAGE_DIRS := $(patsubst %/Dockerfile,%,$(DOCKERFILES))
IMAGE_NAMES := $(foreach dir,$(DOCKER_IMAGE_DIRS),$(patsubst %,$(IMAGE_PREFIX)%,$(shell basename $(dir))))
images:
	$(info $(IMAGE_NAMES))
	@echo > /dev/null


# Define imagetag-golang, etc, for each image, which parses the dockerfile and
# prints an image tag. For example:
#     FROM golang:1.8.1-stretch
# in the "foo/Dockerfile" becomes:
#     $ make imagetag-foo
#     1.8.1-stretch
define imagetag_dep
.PHONY: imagetag-$(1)
$(patsubst $(IMAGE_PREFIX)%,imagetag-%,$(1)): $(patsubst $(IMAGE_PREFIX)%,%,$(1))/Dockerfile
	@cat $$< | grep "^FROM " | head -n1 | sed 's/FROM \(.*\):\(.*\)/\2/'
endef
$(foreach image, $(IMAGE_NAMES), $(eval $(call imagetag_dep, $(image))))

all: $(UPTODATE_FILES) binaries

check: all lint unit-tests container-tests

LOCAL_BINARIES = \
	cmd/wk/wk \
	cmd/wks-entitle/wks-entitle \
	cmd/update-manifest/update-manifest \
	cmd/wks-ci/wks-ci

BINARIES = \
	$(LOCAL_BINARIES) \
	cmd/github-service/github-service \
	cmd/gitops-repo-broker/gitops-repo-broker \
	cmd/mock-authz-server/server \
	cmd/mock-https-authz-server/server \
	cmd/ui-server/ui-server \
	cmd/wks-ci/checks/policy/policy \
	kerberos/cmd/k8s-krb5-server/server \
	kerberos/cmd/wk-kerberos/wk-kerberos \
	$(NULL)

binaries: $(BINARIES)

godeps=$(shell go list -f '{{join .Deps "\n"}}' $1 | \
	   xargs go list -f \
	   '{{if not .Standard}}{{ $$dep := . }}{{range .GoFiles}}{{$$dep.Dir}}/{{.}} {{end}}{{end}}')

DEPS=$(call godeps,./cmd/wk)

USER_GUIDE_SOURCES=$(shell find user-guide/ -name public -prune -o -print)
user-guide/public: $(USER_GUIDE_SOURCES)
	cd user-guide && ./make-static.sh

pkg/guide/assets_vfsdata.go: user-guide/public
	go generate ./pkg/guide

POLICIES=$(shell find pkg/opa/policy/rego -name '*.rego' -print)
pkg/opa/policy/policy_vfsdata.go: $(POLICIES)
	go generate ./pkg/opa/policy

SETUP=$(shell find setup -name bin -prune -o -print)
pkg/setup/setup_vfsdata.go: $(SETUP)
	RELEASE_GOOS=$(LOCAL_BINARIES_GOOS) ./tools/build/setup/build-release.sh $(CURRENT_DIR)/setup $(CURRENT_DIR)/setup/wk-quickstart/setup/dependencies.toml
	go generate ./pkg/setup
	@rm -rf $(CURRENT_DIR)/setup/wk-quickstart/.git

GENERATED = pkg/guide/assets_vfsdata.go pkg/opa/policy/policy_vfsdata.go pkg/setup/setup_vfsdata.go

cmd/wk/wk: $(DEPS) $(GENERATED)
cmd/wk/wk: cmd/wk/*.go
	CGO_ENABLED=0 GOOS=$(LOCAL_BINARIES_GOOS) GOARCH=amd64 go build -ldflags "-X github.com/weaveworks/wks/pkg/version.Version=$(VERSION) -X github.com/weaveworks/wks/pkg/version.ImageTag=$(IMAGE_TAG)" -o $@ cmd/wk/*.go

cmd/wks-ci/checks/policy/.uptodate: cmd/policy/policy
cmd/wks-ci/checks/policy/policy: cmd/wks-ci/checks/policy/*.go $(GENERATED)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ cmd/wks-ci/checks/policy/*.go

ENTITLE_DEPS=$(call godeps,./cmd/wks-entitle)
cmd/wks-entitle/wks-entitle: $(ENTITLE_DEPS)
	CGO_ENABLED=0 GOOS=$(LOCAL_BINARIES_GOOS) GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ cmd/wks-entitle/*.go

CI_DEPS=$(call godeps,./cmd/wks-ci)

cmd/wks-ci/.uptodate: cmd/wks-ci/wks-ci cmd/wks-ci/checks/policy/policy cmd/wks-ci/Dockerfile
cmd/wks-ci/wks-ci: $(CI_DEPS) cmd/wks-ci/*.go
	CGO_ENABLED=0 GOOS=$(LOCAL_BINARIES_GOOS) GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ cmd/wks-ci/*.go

UPDATE_MANIFEST_DEPS=$(call godeps,./cmd/update-manifest)
cmd/update-manifest/update-manifest: $(UPDATE_MANIFEST_DEPS) cmd/update-manifest/*.go
	CGO_ENABLED=0 GOOS=$(LOCAL_BINARIES_GOOS) GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ cmd/update-manifest/*.go

kerberos/cmd/wk-kerberos/wk-kerberos: $(shell find kerberos/cmd/wk-kerberos/ -type f -name '*.go')
	CGO_ENABLED=0 GOARCH=amd64 go build -o $@ ./kerberos/cmd/wk-kerberos

kerberos/cmd/k8s-krb5-server/.uptodate: kerberos/cmd/k8s-krb5-server/server kerberos/cmd/k8s-krb5-server/Dockerfile
kerberos/cmd/k8s-krb5-server/server: kerberos/cmd/k8s-krb5-server/*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ ./kerberos/cmd/k8s-krb5-server

cmd/mock-authz-server/.uptodate: cmd/mock-authz-server/server cmd/mock-authz-server/Dockerfile
cmd/mock-authz-server/server: cmd/mock-authz-server/*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ cmd/mock-authz-server/*.go

cmd/mock-https-authz-server/.uptodate: cmd/mock-https-authz-server/server cmd/mock-https-authz-server/Dockerfile
cmd/mock-https-authz-server/server: cmd/mock-https-authz-server/*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ cmd/mock-https-authz-server/*.go

cmd/github-service/.uptodate: cmd/github-service/github-service cmd/github-service/Dockerfile
cmd/github-service/github-service: cmd/github-service/*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ ./cmd/github-service

cmd/gitops-repo-broker/.uptodate: cmd/gitops-repo-broker/gitops-repo-broker cmd/gitops-repo-broker/Dockerfile
cmd/gitops-repo-broker/gitops-repo-broker: cmd/gitops-repo-broker/*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $@ ./cmd/gitops-repo-broker

# UI
UI_CODE_DEPS = $(shell find ui/src -name '*.jsx' -or -name '*.json')
UI_BUILD_DEPS = \
	ui/.babelrc.js \
	ui/.eslintrc.js \
	ui/server.js \
	ui/yarn.lock \
	ui/webpack.common.js \
	ui/webpack.production.js
UI_DEPS = $(UI_CODE_DEPS) $(UI_BUILD_DEPS)
ui/build: $(UI_DEPS) user-guide/public
	cd ui && yarn install --frozen-lockfile && yarn lint && yarn build
	cp -r user-guide/public ui/build/docs

ui/.uptodate: ui-server/ui-server ui/Dockerfile ui/build
ui-server/ui-server: $(SERVER_DEPS)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o $@ ui-server/*.go

install: $(LOCAL_BINARIES)
	cp $(LOCAL_BINARIES) `go env GOPATH`/bin

EMBEDMD_FILES = \
	docs/entitlements.md \
	$(NULL)

lint:
	bin/go-lint
	bin/check-embedmd.sh $(EMBEDMD_FILES)

clean:
	$(SUDO) docker rmi $(IMAGE_NAMES) >/dev/null 2>&1 || true
	$(SUDO) docker rmi $(patsubst %, %:$(IMAGE_TAG), $(IMAGE_NAMES)) >/dev/null 2>&1 || true
	rm -rf $(UPTODATE_FILES)
	rm -f $(BINARIES)
	rm -f $(GENERATED)
	rm -rf ui/build

push:
	for IMAGE_NAME in $(IMAGE_NAMES); do \
		docker push $$IMAGE_NAME:$(IMAGE_TAG); \
	done

generate-manifests:
	cd wkp-cluster-components && yarn && yarn generate-manifests

# We select which directory we want to descend into to not execute integration
# tests here.
unit-tests: $(GENERATED)
	go test -v ./cmd/... ./pkg/...

container-tests:  test/container/images/centos7/.uptodate
	go test -count=1 ./test/container/...

FORCE:
