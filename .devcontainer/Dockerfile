# debian version
# 10 (Buster)
# 11 (Bullseye)
# 12 (Bookworm)
# 13 (Trixie)
# 14 (Forky)

# for main image
ARG DOCKER_DEBIAN_VERSION=12.11

# for builder images
ARG DOCKER_GOLANG_VERSION=1.24.4-bookworm
ARG DOCKER_NODE_VERSION=22.17.0-bookworm
ARG DOCKER_PYTHON_VERSION=3.13-bookworm

# for go install
ARG GO_DOCKERFMT_VERSION=v0.3.7
ARG GO_SHFMT_VERSION=v3.11.0
ARG GO_GO_TASK_VERSION=v3.44.0
ARG GO_MKCERT_VERSION=v1.4.4
ARG GO_DELVE_VERSION=v1.25.0
ARG GO_GOPLS_VERSION=v0.19.1
ARG GO_STATICCHECK_VERSION=v0.6.1
ARG GO_GO_LICENSES=v1.6.0
ARG GO_OAPI_CODEGEN_VERSION=v2.4.1
ARG GO_OGEN_VERSION=v1.14.0
ARG GO_GOLANGCI_LINT_VERSION=v2.2.0
ARG GO_SQLC_VERSION=v1.29.0
ARG GO_GOLANG_MIGRATE_VERSION=v4.18.3
ARG GO_DIVE_VERSION=v0.13.1

# for npm tools
ARG NPM_NPM_VERSION=11.4.2
ARG NPM_PRETTIER_VERSION=3.6.2
ARG NPM_RENOVATE_VERSION=41.15.0
ARG NPM_CLAUDE_CODE_VERSION=1.0.35
ARG NPM_CCUSAGE_VERSION=15.2.0
ARG NPM_REPOMIX_VERSION=0.3.9
ARG NPM_OPENAPI_GENERATOR_CLI_VERSION=2.21.0
ARG NPM_IBM_OPENAPI_VALIDATOR_VERSION=1.35.2
ARG NPM_REDOCLY_CLI_VERSION=1.34.4
ARG NPM_SQL_FORMATTER_VERSION=15.6.6

# for pip tools
ARG PIP_PIP_VERSION=25.1.1
ARG PIP_PIPX_VERSION=1.7.1
ARG PIP_GITINGEST_VERSION=0.1.5

# for other tools (github-tags)
ARG OTHER_JQ_VERSION=1.8.0
ARG OTHER_YQ_VERSION=4.45.4
ARG OTHER_XQ_VERSION=1.3.0
ARG OTHER_TAPLO_VERSION=0.10.0
ARG OTHER_SHELLCHECK_VERSION=v0.10.0
ARG OTHER_RIPGREP_VERSION=14.1.1

# ##########################################################################################
# builder images
# ##########################################################################################

# golang
# https://hub.docker.com/_/golang
FROM golang:${DOCKER_GOLANG_VERSION} AS golang-builder

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

# golangci-lint
# https://golangci-lint.run/welcome/install/#local-installation
# https://golangci-lint.run/welcome/install/#install-from-sources
ARG GO_GOLANGCI_LINT_VERSION
RUN go install -v github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GO_GOLANGCI_LINT_VERSION}

# sqlc
# https://github.com/sqlc-dev/sqlc
# https://sqlc.dev/
# https://hub.docker.com/r/sqlc/sqlc
ARG GO_SQLC_VERSION
RUN go install -v github.com/sqlc-dev/sqlc/cmd/sqlc@${GO_SQLC_VERSION}

# golang-migrate
# https://github.com/golang-migrate/migrate
# https://hub.docker.com/r/migrate/migrate
# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#with-go-toolchain
ARG GO_GOLANG_MIGRATE_VERSION
RUN go install -v -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@${GO_GOLANG_MIGRATE_VERSION}

# dive
# https://github.com/wagoodman/dive
# https://hub.docker.com/r/wagoodman/dive
ARG GO_DIVE_VERSION
RUN go install -v github.com/wagoodman/dive@${GO_DIVE_VERSION}

# node
# https://hub.docker.com/_/node
# https://note.milldea.com/posts/two-ways-to-install-nodejs-with-fixed-version-in-dockerfile
FROM node:${DOCKER_NODE_VERSION} AS node-builder

# python
# https://hub.docker.com/_/python
FROM python:${DOCKER_PYTHON_VERSION} AS python-builder
RUN mkdir -p /python-install \
    && cp -r /usr/local/bin /python-install/ \
    && cp -r /usr/local/lib /python-install/ \
    && cp -r /usr/local/include /python-install/ \
    && rm -rf /python-install/lib/pkgconfig

# ##########################################################################################
# main image
# ##########################################################################################

# debian
# https://hub.docker.com/_/debian/
FROM debian:${DOCKER_DEBIAN_VERSION}
ENV DEBIAN_FRONTEND=noninteractive
ENV DEBCONF_NOWARNINGS=yes

# add distribution packages
RUN apt-get update -y && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    wget \
    git \
    nano \
    less \
    gzip \
    tar \
    unzip \
    xz-utils \
    locales-all

# golang
# refs above
COPY --from=golang-builder /usr/local/go /usr/local/go
ENV PATH=/usr/local/go/bin/:$PATH

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

# golangci-lint
# refs above
COPY --from=golang-builder /go/bin/golangci-lint /usr/local/bin/

# sqlc
# refs above
COPY --from=golang-builder /go/bin/sqlc /usr/local/bin/

# golang-migrate
# refs above
COPY --from=golang-builder /go/bin/migrate /usr/local/bin/

# dive
# refs above
COPY --from=golang-builder /go/bin/dive /usr/local/bin/

# node
# refs above
COPY --from=node-builder /usr/local/bin/node /usr/local/bin/node
COPY --from=node-builder /usr/local/lib/node_modules /usr/local/lib/node_modules
RUN ln -s /usr/local/bin/node /usr/local/bin/nodejs \
    && ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm \
    && ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npx

# python
# refs above
RUN apt-get update -y && apt-get install -y expat
COPY --from=python-builder /python-install/bin/ /usr/local/bin/
COPY --from=python-builder /python-install/lib/ /usr/local/lib/
COPY --from=python-builder /python-install/include/ /usr/local/include/
ENV PIP_ROOT_USER_ACTION=ignore

# npm
# https://github.com/npm/cli
# https://www.npmjs.com/package/npm
ARG NPM_NPM_VERSION
RUN npm install -g npm@${NPM_NPM_VERSION}

# prettier
# https://prettier.io/docs/install
ARG NPM_PRETTIER_VERSION
RUN npm install -g prettier@${NPM_PRETTIER_VERSION}

# renovate
# https://github.com/renovatebot/renovate
# https://www.npmjs.com/package/renovate
ARG NPM_RENOVATE_VERSION
RUN npm install -g renovate@${NPM_RENOVATE_VERSION}

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

# ccusage
# https://github.com/ryoppippi/ccusage
# https://www.npmjs.com/package/ccusage
ARG NPM_CCUSAGE_VERSION
RUN npm install -g ccusage@${NPM_CCUSAGE_VERSION}

# repomix
# https://github.com/yamadashy/repomix
# https://www.npmjs.com/package/repomix
ARG NPM_REPOMIX_VERSION
RUN npm install -g repomix@${NPM_REPOMIX_VERSION}

# openapi-generator-cli
# https://github.com/OpenAPITools/openapi-generator
# https://github.com/OpenAPITools/openapi-generator-cli
ARG NPM_OPENAPI_GENERATOR_CLI_VERSION
RUN npm install -g @openapitools/openapi-generator-cli@${NPM_OPENAPI_GENERATOR_CLI_VERSION}

# lint-openapi
# https://github.com/IBM/openapi-validator
ARG NPM_IBM_OPENAPI_VALIDATOR_VERSION
RUN npm install -g ibm-openapi-validator@${NPM_IBM_OPENAPI_VALIDATOR_VERSION}

# redocly-cli
# https://github.com/Redocly/redocly-cli
ARG NPM_REDOCLY_CLI_VERSION
RUN npm install -g @redocly/cli@${NPM_REDOCLY_CLI_VERSION}

# sql-formatter
# https://github.com/sql-formatter-org/sql-formatter
ARG NPM_SQL_FORMATTER_VERSION
RUN npm install -g sql-formatter@${NPM_SQL_FORMATTER_VERSION}

# pip
# https://pypi.org/project/pip/
ARG PIP_PIP_VERSION
RUN pip install --upgrade pip==${PIP_PIP_VERSION}

# pipx
# https://github.com/pypa/pipx
ARG PIP_PIPX_VERSION
RUN pip install pipx==${PIP_PIPX_VERSION}
RUN pipx ensurepath

# gitingest
# https://github.com/cyclotruc/gitingest
ARG PIP_GITINGEST_VERSION
RUN pip install gitingest==${PIP_GITINGEST_VERSION}

# jq
# https://github.com/jqlang/jq
ARG OTHER_JQ_VERSION
RUN curl -fsSL -o /jq https://github.com/jqlang/jq/releases/download/jq-${OTHER_JQ_VERSION}/jq-linux-$(uname -m | sed 's/aarch64/arm64/' | sed 's/x86_64/amd64/') \
    && chmod +x /jq \
    && mv /jq /usr/local/bin/ \
    && jq --version

# yq
# https://github.com/mikefarah/yq
# https://hub.docker.com/r/mikefarah/yq
ARG OTHER_YQ_VERSION
RUN curl -fsSL -o /yq https://github.com/mikefarah/yq/releases/download/v${OTHER_YQ_VERSION}/yq_linux_$(uname -m | sed 's/aarch64/arm64/' | sed 's/x86_64/amd64/') \
    && chmod +x /yq \
    && mv /yq /usr/local/bin/ \
    && yq --version

# xq
# https://github.com/sibprogrammer/xq
ARG OTHER_XQ_VERSION
RUN curl -fsSL -o /xq.tar.gz https://github.com/sibprogrammer/xq/releases/download/v${OTHER_XQ_VERSION}/xq_${OTHER_XQ_VERSION}_linux_$(uname -m | sed 's/aarch64/arm64/' | sed 's/x86_64/amd64/').tar.gz \
    && mkdir -p /xq \
    && tar -xvzf /xq.tar.gz -C /xq/ \
    && chmod +x /xq/xq \
    && mv /xq/xq /usr/local/bin/ \
    && rm /xq.tar.gz \
    && rm -rf /xq \
    && xq --version

# taplo
# https://github.com/tamasfe/taplo
# https://taplo.tamasfe.dev/
# https://hub.docker.com/r/tamasfe/taplo
ARG OTHER_TAPLO_VERSION
RUN curl -fsSL -o /taplo.gz https://github.com/tamasfe/taplo/releases/download/${OTHER_TAPLO_VERSION}/taplo-linux-$(uname -m | sed 's/arm64/aarch64/' | sed 's/amd64/x86_64/').gz \
    && gzip -d /taplo.gz \
    && chmod +x /taplo \
    && mv /taplo /usr/local/bin/ \
    && taplo --version

# shellcheck
# https://github.com/koalaman/shellcheck
# https://hub.docker.com/r/koalaman/shellcheck
ARG OTHER_SHELLCHECK_VERSION
RUN curl -fsSL -o /shellcheck.tar.xz https://github.com/koalaman/shellcheck/releases/download/${OTHER_SHELLCHECK_VERSION}/shellcheck-${OTHER_SHELLCHECK_VERSION}.linux.$(uname -m | sed 's/arm64/aarch64/' | sed 's/amd64/x86_64/').tar.xz \
    && tar -xvf /shellcheck.tar.xz -C / \
    && chmod +x /shellcheck-${OTHER_SHELLCHECK_VERSION}/shellcheck \
    && mv /shellcheck-${OTHER_SHELLCHECK_VERSION}/shellcheck /usr/local/bin/ \
    && rm /shellcheck.tar.xz \
    && rm -rf /shellcheck-${OTHER_SHELLCHECK_VERSION} \
    && shellcheck --version

# ripgrep
# https://github.com/BurntSushi/ripgrep
ARG OTHER_RIPGREP_VERSION
RUN mkdir -p /ripgrep \
    && curl -fsSL -o /ripgrep/ripgrep.tar.gz -L https://github.com/BurntSushi/ripgrep/releases/download/${OTHER_RIPGREP_VERSION}/ripgrep-${OTHER_RIPGREP_VERSION}-$(uname -m)-unknown-linux-gnu.tar.gz \
    && cd /ripgrep \
    && tar -xvzf ripgrep.tar.gz \
    && mv ripgrep-${OTHER_RIPGREP_VERSION}-$(uname -m)-unknown-linux-gnu/rg /usr/local/bin/ \
    && chmod +x /usr/local/bin/rg \
    && rm -rf /ripgrep

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
