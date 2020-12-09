#!/usr/bin/env bash
protoc  -I ../databus/api/api.proto \
        --go_out=plugins=grpc:./databus/api
