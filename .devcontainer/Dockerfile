# debian version
# 10 (Buster)
# 11 (Bullseye)
# 12 (Bookworm)
# 13 (Trixie)
# 14 (Forky)

# for main image
ARG DOCKER_GO_VERSION=1.24.4-bookworm

# for builder images
ARG DOCKER_NODE_VERSION=22.16.0-bookworm
ARG DOCKER_JQ_VERSION=1.8.0
ARG DOCKER_YQ_VERSION=4.45.4
ARG DOCKER_GOLANGCI_LINT_VERSION=v2.1.6
ARG DOCKER_SQLC_VERSION=1.29.0
ARG DOCKER_GOLANG_MIGRATE_VERSION=v4.18.3
ARG DOCKER_DIVE_VERSION=v0.13.1

# for go install
ARG GO_DOCKERFMT_VERSION=v0.3.7
ARG GO_SHFMT_VERSION=v3.11.0
ARG GO_GO_TASK_VERSION=v3.44.0
ARG GO_MKCERT_VERSION=v1.4.4
ARG GO_DELVE_VERSION=v1.24.2
ARG GO_GOPLS_VERSION=v0.18.1
ARG GO_STATICCHECK_VERSION=v0.6.1
ARG GO_GO_LICENSES=v1.6.0
ARG GO_OAPI_CODEGEN_VERSION=v2.4.1
ARG GO_OGEN_VERSION=v1.14.0

# for npm tools
ARG NPM_PRETTIER_VERSION=3.5.3
ARG NPM_RENOVATE_VERSION=40.61.2
ARG NPM_OPENAPI_GENERATOR_CLI_VERSION=2.20.2
ARG NPM_IBM_OPENAPI_VALIDATOR_VERSION=1.35.2
ARG NPM_REDOCLY_CLI_VERSION=1.34.3
ARG NPM_SQL_FORMATTER_VERSION=15.6.2
ARG NPM_CLAUDE_CODE_VERSION=1.0.19
ARG NPM_CCUSAGE_VERSION=0.7.0

# for pipx tools
ARG PIPX_GITINGEST_VERSION=0.1.4

# ##########################################################################################
# builder images
# ##########################################################################################

# golang
# https://hub.docker.com/_/golang
FROM golang:${DOCKER_GO_VERSION} AS golang-builder

# dockerfmt
# https://github.com/reteps/dockerfmt
ARG GO_DOCKERFMT_VERSION
RUN go install github.com/reteps/dockerfmt@${GO_DOCKERFMT_VERSION}

# shfmt
# https://github.com/mvdan/sh
ARG GO_SHFMT_VERSION
RUN go install mvdan.cc/sh/v3/cmd/shfmt@${GO_SHFMT_VERSION}

# go-task
# https://github.com/go-task/task
# https://taskfile.dev/
ARG GO_GO_TASK_VERSION
RUN go install github.com/go-task/task/v3/cmd/task@${GO_GO_TASK_VERSION}

# mkcert
# https://github.com/FiloSottile/mkcert
ARG GO_MKCERT_VERSION
RUN go install filippo.io/mkcert@${GO_MKCERT_VERSION}

# delve
# https://github.com/go-delve/delve
ARG GO_DELVE_VERSION
RUN go install github.com/go-delve/delve/cmd/dlv@${GO_DELVE_VERSION}

# gopls
# https://pkg.go.dev/golang.org/x/tools/gopls
# https://cs.opensource.google/go/x/tools
ARG GO_GOPLS_VERSION
RUN go install golang.org/x/tools/gopls@${GO_GOPLS_VERSION}

# staticcheck
# https://godoc.org/honnef.co/go/tools
# https://github.com/dominikh/go-tools
ARG GO_STATICCHECK_VERSION
RUN go install honnef.co/go/tools/cmd/staticcheck@${GO_STATICCHECK_VERSION}

# go-licenses
# https://github.com/google/go-licenses
ARG GO_GO_LICENSES
RUN go install github.com/google/go-licenses@${GO_GO_LICENSES}

# oapi-codegen
# https://github.com/oapi-codegen/oapi-codegen
ARG GO_OAPI_CODEGEN_VERSION
RUN go install -v github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@${GO_OAPI_CODEGEN_VERSION}

# ogen
# https://github.com/ogen-go/ogen
ARG GO_OGEN_VERSION
RUN go install -v github.com/ogen-go/ogen/cmd/ogen@${GO_OGEN_VERSION}

# node
# https://hub.docker.com/_/node
# https://note.milldea.com/posts/two-ways-to-install-nodejs-with-fixed-version-in-dockerfile
FROM node:${DOCKER_NODE_VERSION} AS node-builder

# jq
# https://github.com/mikefarah/yq
# https://github.com/jqlang/jq
FROM ghcr.io/jqlang/jq:${DOCKER_JQ_VERSION} AS jq-builder

# yq
# https://github.com/mikefarah/yq
# https://hub.docker.com/r/mikefarah/yq
FROM mikefarah/yq:${DOCKER_YQ_VERSION} AS yq-builder

# golangci-lint
# https://golangci-lint.run/welcome/install/#local-installation
FROM golangci/golangci-lint:${DOCKER_GOLANGCI_LINT_VERSION} AS golangci-lint-builder

# sqlc
# https://github.com/sqlc-dev/sqlc
# https://sqlc.dev/
# https://hub.docker.com/r/sqlc/sqlc
FROM sqlc/sqlc:${DOCKER_SQLC_VERSION} AS sqlc-builder

# golang-migrate
# https://github.com/golang-migrate/migrate
# https://hub.docker.com/r/migrate/migrate
FROM migrate/migrate:${DOCKER_GOLANG_MIGRATE_VERSION} AS golang-migrate-builder

# dive
# https://github.com/wagoodman/dive
# https://hub.docker.com/r/wagoodman/dive
FROM wagoodman/dive:${DOCKER_DIVE_VERSION} AS dive-builder

# ##########################################################################################
# main image
# ##########################################################################################

# golang
# https://hub.docker.com/_/golang
# https://github.com/devcontainers/images/tree/main/src/go
# https://github.com/devcontainers/images/blob/main/src/go/.devcontainer/Dockerfile
# https://mcr.microsoft.com/v2/vscode/devcontainers/go/tags/list
FROM golang:${DOCKER_GO_VERSION}
ENV DEBIAN_FRONTEND=noninteractive
ENV DEBCONF_NOWARNINGS=yes

# add distribution packages
RUN apt-get -y update
RUN apt-get -y install \
    ca-certificates \
    curl \
    git \
    nano \
    less

# dockerfmt
# refs above
COPY --from=golang-builder /go/bin/dockerfmt /usr/local/bin/

# shfmt
# refs above
COPY --from=golang-builder /go/bin/shfmt /usr/local/bin/

# go-task
# refs above
COPY --from=golang-builder /go/bin/task /usr/local/bin/

# mkcert
# refs above
COPY --from=golang-builder /go/bin/mkcert /usr/local/bin/

# delve
# refs above
COPY --from=golang-builder /go/bin/dlv /usr/local/bin/

# gopls
# refs above
COPY --from=golang-builder /go/bin/gopls /usr/local/bin/

# staticcheck
# refs above
COPY --from=golang-builder /go/bin/staticcheck /usr/local/bin/

# go-licenses
# refs above
COPY --from=golang-builder /go/bin/go-licenses /usr/local/bin/

# oapi-codegen
# refs above
COPY --from=golang-builder /go/bin/oapi-codegen /usr/local/bin/

# ogen
# refs above
COPY --from=golang-builder /go/bin/ogen /usr/local/bin/

# node
# refs above
COPY --from=node-builder /usr/local/bin/node /usr/local/bin/node
COPY --from=node-builder /usr/local/lib/node_modules /usr/local/lib/node_modules
RUN ln -s /usr/local/bin/node /usr/local/bin/nodejs \
    && ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm \
    && ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npx
RUN npm install -g yarn

# jq
# refs above
COPY --from=jq-builder /jq /usr/local/bin/

# yq
# refs above
COPY --from=yq-builder /usr/bin/yq /usr/local/bin/

# golangci-lint
# refs above
COPY --from=golangci-lint-builder /usr/bin/golangci-lint /usr/local/bin/

# sqlc
# refs above
COPY --from=sqlc-builder /workspace/sqlc /usr/local/bin/

# golang-migrate
# refs above
COPY --from=golang-migrate-builder /usr/local/bin/migrate /usr/local/bin/

# dive
# refs above
COPY --from=dive-builder /usr/local/bin/dive /usr/local/bin/

# prettier
# https://prettier.io/docs/install
ARG NPM_PRETTIER_VERSION
RUN yarn global add prettier@${NPM_PRETTIER_VERSION}

# renovate
# https://github.com/renovatebot/renovate
# https://www.npmjs.com/package/renovate
ARG NPM_RENOVATE_VERSION
RUN yarn global add renovate@${NPM_RENOVATE_VERSION}

# openapi-generator-cli
# https://github.com/OpenAPITools/openapi-generator
# https://github.com/OpenAPITools/openapi-generator-cli
ARG NPM_OPENAPI_GENERATOR_CLI_VERSION
RUN yarn global add @openapitools/openapi-generator-cli@${NPM_OPENAPI_GENERATOR_CLI_VERSION}

# lint-openapi
# https://github.com/IBM/openapi-validator
ARG NPM_IBM_OPENAPI_VALIDATOR_VERSION
RUN yarn global add ibm-openapi-validator@${NPM_IBM_OPENAPI_VALIDATOR_VERSION}

# redocly-cli
# https://github.com/Redocly/redocly-cli
ARG NPM_REDOCLY_CLI_VERSION
RUN yarn global add @redocly/cli@${NPM_REDOCLY_CLI_VERSION}

# sql-formatter
# https://github.com/sql-formatter-org/sql-formatter
ARG NPM_SQL_FORMATTER_VERSION
RUN yarn global add sql-formatter@${NPM_SQL_FORMATTER_VERSION}

# pipx
# https://github.com/pypa/pipx
RUN apt-get -y install pipx
RUN pipx ensurepath

# gitingest
# https://github.com/cyclotruc/gitingest
ARG PIPX_GITINGEST_VERSION
RUN pipx install gitingest==${PIPX_GITINGEST_VERSION}

# claude-code
# https://github.com/anthropics/claude-code
# https://github.com/anthropics/claude-code/blob/main/.devcontainer/devcontainer.json
# https://github.com/anthropics/claude-code/blob/main/.devcontainer/Dockerfile
# https://github.com/anthropics/claude-code/blob/main/.devcontainer/init-firewall.sh
RUN apt update -y && apt install -y \
    less \
    git \
    procps \
    sudo \
    fzf \
    zsh \
    man-db \
    unzip \
    gnupg2 \
    gh \
    iptables \
    ipset \
    iproute2 \
    dnsutils \
    aggregate
ARG NPM_CLAUDE_CODE_VERSION
RUN npm install -g @anthropic-ai/claude-code@${NPM_CLAUDE_CODE_VERSION}
ENV DEVCONTAINER=true

# init-firewall.sh
# https://docs.docker.com/reference/dockerfile/#adding-files-from-a-git-repository
# https://docs.docker.com/build/concepts/context/#url-fragments
# ADD https://github.com/anthropics/claude-code.git#main:.devcontainer/init-firewall.sh /usr/local/bin/
ARG CLAUDE_CODE_COMMIT_HASH=bd5ca708adf82c4b81857abf40fe36d9d9cc3d1c
RUN curl -fsSL -o /init-firewall.sh https://raw.githubusercontent.com/anthropics/claude-code/${CLAUDE_CODE_COMMIT_HASH}/.devcontainer/init-firewall.sh \
    && chmod 644 /init-firewall.sh
RUN chmod +x /init-firewall.sh \
    && echo "node ALL=(root) NOPASSWD: /usr/local/bin/init-firewall.sh" >/etc/sudoers.d/node-firewall \
    && chmod 0440 /etc/sudoers.d/node-firewall

## ccusage
# https://github.com/ryoppippi/ccusage
# https://www.npmjs.com/package/ccusage
ARG NPM_CCUSAGE_VERSION
RUN npm install -g ccusage@${NPM_CCUSAGE_VERSION}

# add vscode user
# https://code.visualstudio.com/remote/advancedcontainers/add-nonroot-user#_creating-a-nonroot-user
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update \
    && apt-get install -y sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL >/etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME
USER $USERNAME
