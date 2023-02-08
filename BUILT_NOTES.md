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
