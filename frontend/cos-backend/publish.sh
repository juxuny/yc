#!/bin/bash
set -e
if [ -z "$@" ]; then echo 'missing version'; exit 255; fi
VERSION="$@"
docker build -t registry.cn-guangzhou.aliyuncs.com/juxuny-public/cos-backend:"${VERSION}" . -f deploy/backend.dockerfile
docker push registry.cn-guangzhou.aliyuncs.com/juxuny-public/cos-backend:"${VERSION}"
