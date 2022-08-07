#!/bin/bash
set -e
echo $YC_HOME
if [ -z "$YC_HOME" ]; then echo 'missing YC_HOME'; exit 255; fi
YC_DEV="go run ${YC_HOME}/cmd/yc"
eval "${YC_DEV}" service update
eval "${YC_DEV}" model update --go
eval "${YC_DEV}" client gen --go=. --internal
eval "${YC_DEV}" client gen --react-ts="${YC_HOME}/frontend/cos-backend/src/services/api"
