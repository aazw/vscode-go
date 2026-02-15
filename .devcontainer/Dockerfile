FROM jetpackio/devbox:0.16.0 AS builder

WORKDIR /code/vscode-go
USER root:root
RUN mkdir -p /code/vscode-go && chown ${DEVBOX_USER}:${DEVBOX_USER} /code/vscode-go
USER ${DEVBOX_USER}:${DEVBOX_USER}

COPY --chown=${DEVBOX_USER}:${DEVBOX_USER} devbox.json devbox.json
COPY --chown=${DEVBOX_USER}:${DEVBOX_USER} devbox.lock devbox.lock
RUN devbox install


FROM aazw/devcontainers-devbox-go:10 AS runtime

USER root
COPY --from=builder /nix/store /nix/store
COPY --from=builder /code/vscode-go/.devbox /code/vscode-go/.devbox
USER vscode

ENV PATH="/code/vscode-go/.devbox/nix/profile/default/bin:${PATH}"
