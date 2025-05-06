# グローバルスコープ (最初のFROMより前に書かれたARGのこと)
ARG GOLANG_VERSION=1.24
ARG JQ_VERSION=1.7.1
ARG DOCKERFMT_VERSION=latest
ARG GOLANGCI_LINT_VERSION=v2.1.6

# ##############################
# builder images
# ##############################

# jq image
# https://github.com/jqlang/jq
FROM ghcr.io/jqlang/jq:${JQ_VERSION} AS jq-builder

# golangci-lint
# https://golangci-lint.run/welcome/install/#local-installation
FROM golangci/golangci-lint:${GOLANGCI_LINT_VERSION} AS golangci-lint-builder

# ##############################
# main image
# ##############################

# https://mcr.microsoft.com/v2/vscode/devcontainers/go/tags/list
FROM mcr.microsoft.com/devcontainers/go:${GOLANG_VERSION}

RUN apt-get -y update
RUN apt-get -y install \
    apt-transport-https \
    build-essential \
    ca-certificates \
    gnupg \
    gnupg-agent \
    lsb-release \
    software-properties-common \
    curl \
    git

# Install Docker
# https://matsuand.github.io/docs.docker.jp.onthefly/engine/install/debian/
RUN for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt-get remove $pkg; done
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
RUN echo \
    "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian \
    $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list >/dev/null
RUN apt-get -y update
RUN apt-get -y install docker-ce docker-ce-cli containerd.io docker-compose-plugin

# dockerfmt
# https://github.com/reteps/dockerfmt
ARG DOCKERFMT_VERSION
RUN go install github.com/reteps/dockerfmt@${DOCKERFMT_VERSION}

# jq (latest)
# https://github.com/jqlang/jq
COPY --from=jq-builder /jq /usr/local/bin/

# golangci-lint
COPY --from=golangci-lint-builder /usr/bin/golangci-lint /usr/local/bin/