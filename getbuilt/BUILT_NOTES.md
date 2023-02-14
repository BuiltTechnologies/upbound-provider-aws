# Built Notes
These are notes on the steps needed to maintain a fork of `provider-aws`.

## Setup

```bash
# load git submodules (hydrates the build/ directory)
make submodules
```

1. Comment out unwanted resources for `ExternalNameConfigs` in `config/externalname.go`.
1. If you completely eliminated any groups:
    1. Manually delete the associated `config/RESOURCE` folder. You have to do this because `generate` does not delete files named `zz_generated.managed*.go`.
    1. Comment out code in `config/provider.go`:
        * the package import (top of file)
        * `Configure` reference (end of file)
1. Run `generate`.
    ```bash
    make generate
    ```

## Whitelisted Resources
|Group|Kind|
|---|---|
|iam.aws.upbound.io||

## Publish to ECR

```bash
# * must be on a branch named "main" or "release-*"
export AWS_REGION=us-east-2
export DOCKER_CLI_EXPERIMENTAL=enabled

# login to ECR
# NOTE: this assumes profile `eksdev-admin` for eks communications
AWS_PROFILE=eksdev-admin aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 325405374407.dkr.ecr.us-east-2.amazonaws.com

# create secret in crossplane-system to access ECR
kubectl create secret docker-registry package-pull-secret-ecr \
  --docker-server=325405374407.dkr.ecr.us-east-2.amazonaws.com \
  --docker-username=AWS \
  --docker-password=$(AWS_PROFILE=eksdev-admin aws ecr get-login-password) \
  --namespace=crossplane-system

# build/publish with up uxpkg
# NOTE: you want to use `do.build.xpkgs` because it has a --controller option
#   which sets the spec.controller.image in package.crossplane.yaml when
#   building the XPKG file.
export XPKG_REG_ORGS=325405374407.dkr.ecr.us-east-2.amazonaws.com
export VERSION=v0.29.0-built
make do.build.xpkgs PLATFORMS=linux_amd64
make publish.artifacts PLATFORMS=linux_amd64
```

## Publish to Upbound
This is NOT the recommended way to maintain the provider (use ECR).

```bash
# create an account at upbound
up login

# create repository
# NOTE:only allowed 1 free repository
up repository create provider-aws

# create token for access (expires in 30 days)
up ctp pull-secret create -n crossplane-system

# build / publish
export VERSION=v0.29.0-built
make PLATFORMS=linux_amd64 build
make PLATFORMS=linux_amd64 publish
```

Verify your Marketplace package (replace `builtroller` with org/user): [marketplace.upbound.io/providers/builtroller/provider-aws](https://marketplace.upbound.io/providers/builtroller/provider-aws).


## References
* [Pushing Provider Package to Registry Other Than Docker Hub #2475](https://github.com/crossplane/crossplane/discussions/2475)
* [ECR Image for provider-aws](https://us-east-2.console.aws.amazon.com/ecr/repositories/private/325405374407/provider-aws?region=us-east-2)
* [Build, Publish, and Install Crossplane Package](https://morningspace.medium.com/build-publish-and-install-crossplane-package-5e4b74a3ee37)

## GitHub CI Workflow

```bash
# set PAT
# devcontainer uses GH_TOKEN
export GH_TOKEN=xxx

# pull test image
docker pull nektos/act-environments-ubuntu:18.04

# nektos/act setup
#   --secret-file: store your token in a non-tracked location
#   -P: specifies how to modify the "runs-on" value
alias act='act --secret-file tmp/act-secrets.env -P ubuntu-22.04=catthehacker/ubuntu:act-22.04 '

# run job "publish" locally
act workflow_dispatch -j publish-artifacts

# run on GitHub
gh auth login
gh workflow run 'Publish Helm Charts' --ref roller/gha-charts
```

---

# Appendix

## Map make generate command targets
* generate
    * go.modules.verify
    * `$(TERRAFORM_PROVIDER_SCHEMA): $(TERRAFORM)`
    * go.generate

## Resources In Use
* ec2
* iam
* lambda
* s3
* sqs
* ssm
