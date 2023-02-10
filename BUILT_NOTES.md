# Built Notes
These are notes on the steps needed to maintain a fork of `provider-aws`.

## Setup

```bash
# load git submodules (hydrates the build/ directory)
make submodules

# verify everything is working
make reviewable

# steps
make build.init
VERSION=v0.29.0 PROJECT_VERSION_TAG_GROUP=built make build
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
