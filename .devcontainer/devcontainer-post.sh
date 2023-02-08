#!/bin/bash
set -e

# k8s
# host's ~/.kube/ directory is readonly; config is copied as a starting point
mkdir -p ~/.kube
cp ~/.kube-host/config ~/.kube

# pipenv
# NOTE: .devcontainer/Dockerfile must make the correct version of Python
# the default to allow for the sync without prompting the user to install the
# correct version of Python.
# pipenv sync --dev

# git
# pre-commit install -t commit-msg -t pre-commit

# terraform
mkdir -p tmp
touch tmp/tf.log

# personalizations: place in tmp/devcontainer-post.sh
if [ -f tmp/devcontainer-post.sh ]; then
    # run interactive to give access to aliases
    bash -i tmp/devcontainer-post.sh
fi
