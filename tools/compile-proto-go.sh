#!/usr/bin/env bash

# ------ GENERATE DATA-STORE ------

# LIB
LIB_OUT_DIR="./pb"
DATASTORE_IN_DIR="./proto"

protoc \
    -I="$DATASTORE_IN_DIR" \
    --go_out=$LIB_OUT_DIR \
    --go_opt=paths=import \
    --go-grpc_opt=paths=import \
    --go-grpc_out=$LIB_OUT_DIR \
    "$DATASTORE_IN_DIR"/spaces.proto
