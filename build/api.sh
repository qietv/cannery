#!/usr/bin/env bash
protoPath=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && cd .. && pwd )
echo ${protoPath}
protoc  -I ${protoPath}/broker/api \
        --go_out=plugins=grpc:${protoPath}/broker/api \
        ${protoPath}/broker/api/api.proto
