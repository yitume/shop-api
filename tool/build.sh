#!/bin/bash

VERBOSE=${VERBOSE:-"0"}
V=""
if [[ "${VERBOSE}" == "1" ]];then
    V="-x"
    set -x
fi

ROOTDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

OUT=${1:?"output path"}
VERSION_PACKAGE=${2:?"version go package"} # istio.io/istio/pkg/version
BUILDPATH=${3:?"path to build"}

set -e

BUILD_GOOS=${GOOS:-linux}
BUILD_GOARCH=${GOARCH:-amd64}
GOBINARY=${GOBINARY:-go}
GOPKG="$GOPATH/pkg"
BUILDINFO=${BUILDINFO:-""}
STATIC=${STATIC:-1}
LDFLAGS="-extldflags -static"
GOBUILDFLAGS=${GOBUILDFLAGS:-""}
# Split GOBUILDFLAGS by spaces into an array called GOBUILDFLAGS_ARRAY.
IFS=' ' read -r -a GOBUILDFLAGS_ARRAY <<< "$GOBUILDFLAGS"

GCFLAGS=${GCFLAGS:-}
export CGO_ENABLED=0

if [[ "${STATIC}" !=  "1" ]];then
    LDFLAGS=""
fi

# gather buildinfo if not already provided
# For a release build BUILDINFO should be produced
# at the beginning of the build and used throughout
if [[ -z ${BUILDINFO} ]];then
    BUILDINFO=$(mktemp)
    "${ROOTDIR}/tool/version.sh" > "${BUILDINFO}"
fi

# BUILD LD_EXTRAFLAGS
LD_EXTRAFLAGS=""
while read -r line; do
    read SYMBOL VALUE < <(echo $line)
    LD_EXTRAFLAGS=${LD_EXTRAFLAGS}" -X ${VERSION_PACKAGE}.${SYMBOL}=${VALUE}"
done < "${BUILDINFO}"

set -x
# forgoing -i (incremental build) because it will be deprecated by tool chain.
time GOOS=${BUILD_GOOS} GOARCH=${BUILD_GOARCH} ${GOBINARY} build \
        ${V} "${GOBUILDFLAGS_ARRAY[@]}" ${GCFLAGS:+-gcflags "${GCFLAGS}"} \
        -o "${OUT}" \
        -pkgdir="${GOPKG}/${BUILD_GOOS}_${BUILD_GOARCH}" \
        -ldflags "${LDFLAGS} ${LD_EXTRAFLAGS}" "${BUILDPATH}"
