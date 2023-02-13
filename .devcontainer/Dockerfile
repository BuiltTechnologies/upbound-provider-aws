FROM mcr.microsoft.com/vscode/devcontainers/base:0.203.9-ubuntu-22.04

# GitHub API rate limits unauthenticated requests to 60/hour
# https://docs.github.com/en/rest/overview/resources-in-the-rest-api#rate-limiting
ARG GH_TOKEN
ARG CURL_GH="curl -H 'Authorization: Bearer ${GH_TOKEN}'"

ARG DEBIAN_FRONTEND=noninteractive
ARG TZ=America/Los_Angeles
ARG USER=vscode
ARG PYTHON_VER=3.7

RUN apt update \
    && apt install -y --no-install-recommends --no-install-suggests \
    ack \
    build-essential \
    fd-find \
    fzf \
    gnupg2 \
    pass \
    python3-dev \
    python3-pip \
    software-properties-common \
    tmux \
    tzdata \
    && apt -y --purge autoremove

# default python to PYTHON_VER
RUN echo "deb http://ppa.launchpad.net/deadsnakes/ppa/$(lsb_release -is | tr '[:upper:]' '[:lower:]') $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/python.list \
    && apt-key adv --keyserver keyserver.ubuntu.com/ --recv-keys BA6932366A755776 \
    && apt update \
    && apt install -y \
        python${PYTHON_VER} \
        python${PYTHON_VER}-dev \
        python${PYTHON_VER}-distutils \
    && python${PYTHON_VER} --version
RUN update-alternatives --install /usr/bin/python3 python /usr/bin/python3.10 1 \
    && update-alternatives --install /usr/bin/python3 python /usr/bin/python${PYTHON_VER} 2 \
    && python3 --version
# install correct pip version
RUN curl -fsSL https://bootstrap.pypa.io/get-pip.py | python${PYTHON_VER}

WORKDIR /tmp

# aws
# https://github.com/aws/aws-cli/blob/v2/CHANGELOG.rst?plain=1
ARG AWS_VER=2.7.34
RUN cd /tmp \
    && curl -sSL https://awscli.amazonaws.com/awscli-exe-linux-x86_64-${AWS_VER}.zip -o awscliv2.zip \
    && unzip awscliv2.zip \
    && ./aws/install \
    && rm -rf awscliv2.zip aws \
    && aws --version

# docker
# ------------------------------------------------------------------------------
# https://github.com/microsoft/vscode-dev-containers/tree/master/containers/docker-from-docker

# docker cli
ARG DOCKER_VER=5:20.10.18~3-0~ubuntu-jammy
ENV DOCKER_BUILDKIT=1
RUN curl -fsSL https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]')/gpg | apt-key add - 2>/dev/null \
    && echo "deb [arch=amd64] https://download.docker.com/linux/$(lsb_release -is | tr '[:upper:]' '[:lower:]') $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list \
    && apt update \
    && apt install -y docker-ce-cli=${DOCKER_VER} \
    && groupadd docker \
    && usermod -aG docker ${USER}

RUN echo ${GH_TOKEN}
# docker-compose
ENV COMPOSE_DOCKER_CLI_BUILD=1
RUN LATEST_COMPOSE_VERSION=$(bash -c "$CURL_GH -sSL https://api.github.com/repos/docker/compose/releases/latest" | jq -r .tag_name) \
    && bash -c "${CURL_GH} -sSL https://github.com/docker/compose/releases/download/${LATEST_COMPOSE_VERSION}/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose" \
    && chmod +x /usr/local/bin/docker-compose \
    && docker-compose version

# Default to root only access to the Docker socket, set up non-root init script
RUN touch /var/run/docker-host.sock \
    && ln -s /var/run/docker-host.sock /var/run/docker.sock \
    && apt update \
    && apt -y install socat

# Create docker-init.sh to spin up socat
RUN echo "#!/bin/sh\n\
    sudoIf() { if [ \"\$(id -u)\" -ne 0 ]; then sudo \"\$@\"; else \"\$@\"; fi }\n\
    sudoIf rm -rf /var/run/docker.sock\n\
    ((sudoIf socat UNIX-LISTEN:/var/run/docker.sock,fork,mode=660,user=${USER} UNIX-CONNECT:/var/run/docker-host.sock) 2>&1 >> /tmp/vscr-docker-from-docker.log) & > /dev/null\n\
    \"\$@\"" >> /usr/local/share/docker-init.sh \
    && chmod +x /usr/local/share/docker-init.sh

# VS Code overrides ENTRYPOINT and CMD when executing `docker run` by default.
# Setting the ENTRYPOINT to docker-init.sh will configure non-root access to
# the Docker socket if "overrideCommand": false is set in devcontainer.json.
# The script will also execute CMD if you need to alter startup behaviors.
ENTRYPOINT [ "/usr/local/share/docker-init.sh" ]
CMD [ "sleep", "infinity" ]
# ------------------------------------------------------------------------------

# github cli
ARG GH_CLI_VER=2.16.0
RUN bash -c "${CURL_GH} -sSL https://github.com/cli/cli/releases/download/v${GH_CLI_VER}/gh_${GH_CLI_VER}_linux_amd64.deb -o gh_cli.deb" \
    && dpkg -i gh_cli.deb \
    && gh --version

# tfswitch (terraform)
# https://github.com/warrensbox/terraform-switcher/releases
ARG TFSWITCH_VER=0.13.1308
ENV PATH=${PATH}:/home/${USER}/bin
RUN curl -fsSL https://raw.githubusercontent.com/warrensbox/terraform-switcher/release/install.sh | bash -s -- ${TFSWITCH_VER}

# dive
ARG DIVE_VER=0.9.2
RUN bash -c "${CURL_GH} -fsSL https://github.com/wagoodman/dive/releases/download/v${DIVE_VER}/dive_${DIVE_VER}_linux_amd64.deb -o dive.deb" \
    && apt install ./dive.deb \
    && dive --version

# eksctl
ARG APP=eksctl
RUN bash -c "${CURL_GH} -fsSL https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp \
    && mv /tmp/eksctl /usr/local/bin \
    && eksctl version

# local github actions runner
# using 0.2.21 because of git clone authentication issue
# https://github.com/nektos/act/issues/1343
ARG ACT_VER=v0.2.34
RUN curl -fsSL https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash -s -- ${ACT_VER} \
    && mv bin/act /usr/local/bin \
    && act --version

# kubectl
# https://kubernetes.io/releases/
ARG K8S_VER=1.24.10
RUN curl -fsSLO https://dl.k8s.io/release/v${K8S_VER}/bin/linux/amd64/kubectl \
    && install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl \
    && kubectl version --client

# helm
# https://github.com/helm/helm/releases
ARG HELM_VER=3.11.0
RUN curl -fsSL https://get.helm.sh/helm-v${HELM_VER}-linux-amd64.tar.gz -o helm.tgz \
    && mkdir -p helm-${HELM_VER} \
    && tar -zxvf helm.tgz -C helm-${HELM_VER} \
    && mv helm-${HELM_VER}/linux-amd64/helm /usr/local/bin \
    && helm version

# helmfile
ARG HELMFILE_VER=0.150.0
RUN bash -c "${CURL_GH} -fsSL https://github.com/helmfile/helmfile/releases/download/v${HELMFILE_VER}/helmfile_${HELMFILE_VER}_linux_amd64.tar.gz | tar -xvz" \
    && mv helmfile /usr/local/bin \
    && helmfile version

# k9s
ARG K9S_VER=0.26.7
RUN bash -c "${CURL_GH} -fsSL https://github.com/derailed/k9s/releases/download/v${K9S_VER}/k9s_Linux_x86_64.tar.gz | tar -xvz" \
    && mv k9s /usr/local/bin \
    && k9s version

# go
# https://go.dev/doc/devel/release
ARG GO_VER=1.19.5
ENV PATH=$PATH:/usr/local/go/bin:/home/${USER}/go/bin
RUN curl -sSL https://golang.org/dl/go${GO_VER}.linux-amd64.tar.gz -o go.tgz \
    && tar -C /usr/local -xzf go.tgz \
    && rm -rf go.tgz \
    && go version


# crossplane
ARG CROSSPLANE_VER=v1.11.0
RUN curl -sL https://raw.githubusercontent.com/crossplane/crossplane/master/install.sh | VERSION=${CROSSPLANE_VER} sh \
    && mv kubectl-crossplane /usr/local/bin \
    && kubectl crossplane --version

# upbound cli
ARG UP_VER=v0.15.0
RUN curl -fsSL "https://cli.upbound.io" | VERSION="${UP_VER}" sh \
    && mv up /usr/local/bin \
    && up --version

# helper commands
COPY bin/* /usr/local/bin/

# timezone
RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime \
    && dpkg-reconfigure tzdata


## user-dependent configurations
# ==============================================================================
USER ${USER}

# pyenv
ENV PATH=/home/${USER}/.pyenv/bin:${PATH}
RUN bash -c "${CURL_GH} -fsSL https://github.com/pyenv/pyenv-installer/raw/master/bin/pyenv-installer | bash" \
    && echo 'eval "$(pyenv virtualenv-init -)"' >> ~/.bashrc
# https://github.com/pyenv/pyenv/wiki#suggested-build-environment
RUN sudo apt update \
    && sudo apt install -y --no-install-recommends --no-install-suggests \
        build-essential \
        curl \
        libbz2-dev \
        libffi-dev \
        liblzma-dev \
        libncursesw5-dev \
        libreadline-dev \
        libsqlite3-dev \
        libssl-dev \
        libxml2-dev \
        libxmlsec1-dev \
        llvm \
        make \
        tk-dev \
        wget \
        xz-utils \
        zlib1g-dev \
    && sudo apt -y --purge autoremove

# python
ARG PIP_OPTS="--timeout 1000"
ENV PATH=/home/${USER}/.local/bin:${PATH}
RUN python${PYTHON_VER} -m pip $PIP_OPTS install --upgrade --user \
    pip \
    setuptools \
    wheel
RUN python${PYTHON_VER} -m pip $PIP_OPTS install --user \
    black \
    commitizen \
    credstash \
    isort \
    # kopf[full-auth] \
    pipenv \
    pre-commit \
    jmespath \
    yamllint

# aws-iam-authenticator
ARG AUTH_VER=1.21.2
RUN curl -fsSL -o aws-iam-authenticator https://amazon-eks.s3.us-west-2.amazonaws.com/${AUTH_VER}/2021-07-05/bin/linux/amd64/aws-iam-authenticator \
    && curl -fsSL -o aws-iam-authenticator.sha256 https://amazon-eks.s3.us-west-2.amazonaws.com/${AUTH_VER}/2021-07-05/bin/linux/amd64/aws-iam-authenticator.sha256 \
    && sha256sum --check --status aws-iam-authenticator.sha256 \
    && chmod +x ./aws-iam-authenticator \
    && mv aws-iam-authenticator /home/$USER/.local/bin \
    && aws-iam-authenticator version

# tmux plugin manager
RUN git clone https://github.com/tmux-plugins/tpm /home/${USER}/.tmux/plugins/tpm

# fd
RUN ln -s $(which fdfind) ~/.local/bin/fd

# kube-ps1
RUN git clone https://github.com/jonmosco/kube-ps1.git /home/${USER}/kube-ps1

# go
ENV GO111MODULE=on
ENV GOPATH=/home/${USER}/
# https://go.googlesource.com/vscode-go/+/HEAD/docs/tools.md
RUN go install golang.org/x/tools/gopls@latest \
    && go install golang.org/x/tools/cmd/goimports@latest \
    && go install github.com/josharian/impl@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go install github.com/cweill/gotests/gotests@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/norwoodj/helm-docs/cmd/helm-docs@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest

# node / nvm
# https://nodejs.org/en/download/releases/
# https://github.com/nvm-sh/nvm/releases
ENV NODE_VER=16.17.0
ARG NVM_VER=0.39.1
ENV NVM_DIR=/usr/local/share/nvm
RUN sudo mkdir -p ${NVM_DIR} \
    && sudo chown ${USER}:${USER} ${NVM_DIR}
RUN curl -fsSL -o- https://raw.githubusercontent.com/nvm-sh/nvm/v${NVM_VER}/install.sh | bash \
    && chmod a+x ${NVM_DIR}/nvm.sh \
    && ${NVM_DIR}/nvm.sh

# tfenv
ARG TFENV_VER=v3.0.0
ENV PATH=/home/${USER}/.tfenv/bin:${PATH}
RUN git clone -b ${TFENV_VER} --depth 1 https://github.com/tfutils/tfenv.git /home/${USER}/.tfenv \
    && tfenv --version

# krew
# https://github.com/kubernetes-sigs/krew/releases
ARG KREW_VER=0.4.3
ENV PATH=/home/${USER}/.krew/bin:${PATH}
RUN set -x; cd "$(mktemp -d)" \
    && OS="$(uname | tr '[:upper:]' '[:lower:]')" \
    && ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" \
    && KREW="krew-${OS}_${ARCH}" \
    && bash -c "${CURL_GH} -fsSLO https://github.com/kubernetes-sigs/krew/releases/download/v${KREW_VER}/${KREW}.tar.gz" \
    && tar zxvf "${KREW}.tar.gz" \
    && ./"${KREW}" install krew \
    && kubectl krew version

# aliases / functions
COPY --chown=${USER}:${USER} system/* /tmp/system/
ENV PATH=$PATH:/home/${USER}/scripts
RUN mv /tmp/system/.alias ~ \
    && echo 'source ~/.alias' >> ~/.bashrc \
    && mv /tmp/system/.function ~ \
    && echo 'source ~/.function' >> ~/.bashrc
