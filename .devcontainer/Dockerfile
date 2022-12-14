# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.245.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.19, 1.18, 1-bullseye, 1.19-bullseye, 1.18-bullseye, 1-buster, 1.19-buster, 1.18-buster
ARG VARIANT="1.19-bullseye"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 18, 16, 14
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
USER vscode
ENV HOME /workspaces/go-cli
RUN go get -x github.com/uudashr/gopkgs/v2/cmd/gopkgs \
	github.com/ramya-rao-a/go-outline \
	github.com/nsf/gocode \
	github.com/acroca/go-symbols \
	github.com/fatih/gomodifytags \
	github.com/josharian/impl \
	github.com/haya14busa/goplay/cmd/goplay \
	github.com/go-delve/delve/cmd/dlv \
	golang.org/x/lint/golint \
	golang.org/x/tools/gopls

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
