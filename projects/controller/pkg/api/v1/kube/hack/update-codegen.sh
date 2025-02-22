
#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
ROOT_PKG=github.com/solo-io/gloo/projects/controller/pkg/api/v1
CLIENT_PKG=${ROOT_PKG}/kube/client
APIS_PKG=${ROOT_PKG}/kube/apis

# Grab code-generator version from go.sum.
CODEGEN_PKG=$(go list -f '{{ .Dir }}' -m k8s.io/code-generator)
GENGO_PKG=$(go list -f '{{ .Dir }}' -m k8s.io/gengo/v2)

echo ">> Using ${CODEGEN_PKG}"

# code-generator does work with go.mod but makes assumptions about
# the project living in $GOPATH/src. To work around this and support
# any location; create a temporary directory, use this as an output
# base, and copy everything back once generated.
TEMP_DIR=$(mktemp -d)
cleanup() {
    echo ">> Removing ${TEMP_DIR}"
    rm -rf ${TEMP_DIR}
}
trap "cleanup" EXIT SIGINT

echo ">> Temporary output directory ${TEMP_DIR}"

mkdir -p "${TEMP_DIR}/${ROOT_PKG}/pkg/client/informers" \
         "${TEMP_DIR}/${ROOT_PKG}/pkg/client/listers" \
         "${TEMP_DIR}/${ROOT_PKG}/pkg/client/clientset"

# Ensure we can execute.
chmod +x ${CODEGEN_PKG}/kube_codegen.sh

source ${CODEGEN_PKG}/kube_codegen.sh kube::codegen::gen_client \
    --output-dir "${TEMP_DIR}" \
    --output-pkg "${CLIENT_PKG}" \
    --with-watch \
    --boilerplate "${GENGO_PKG}/boilerplate/boilerplate.go.txt" \
    ${APIS_PKG}

ls -lha $TEMP_DIR

# Copy everything back.
cp -r "${TEMP_DIR}/${ROOT_PKG}/." "${SCRIPT_ROOT}/"

