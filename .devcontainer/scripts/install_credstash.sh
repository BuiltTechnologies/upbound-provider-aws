#!/usr/bin/bash
set -e -o pipefail

# DESCRIPTION: install terraform credstash provider plugin if not already
# present. This process is different if the latest version is needed.

CREDSTASH_VER="0.5.0"
CREDSTASH_DIR="${HOME}/.terraform.d/plugins/registry.terraform.io/hashicorp/credstash/${CREDSTASH_VER}/linux_amd64"
CREDSTASH_FILE="terraform-provider-credstash_v${CREDSTASH_VER}"
CREDSTASH_FILEPATH="${CREDSTASH_DIR}/${CREDSTASH_FILE}"

mkdir -p ${CREDSTASH_DIR}
if [ ! -f ${CREDSTASH_FILEPATH} ]; then
    echo "installing credstash provider plugin: ${CREDSTASH_FILEPATH}"
    curl -fsSL https://github.com/sspinc/terraform-provider-credstash/releases/download/v${CREDSTASH_VER}/terraform-provider-credstash_linux_amd64 -o ${CREDSTASH_FILE}
    mv ${CREDSTASH_FILE} ${CREDSTASH_DIR}
    chmod -R +x ${CREDSTASH_FILEPATH}
    ls -lh ${CREDSTASH_DIR}
fi
