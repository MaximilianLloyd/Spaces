
#!/usr/bin/env bash

# ------ GENERATE DATA-STORE ------

# LIB
LIB_OUT_DIR="client/src/lib/proto/"
DATASTORE_IN_DIR="./proto"

PROTOC="$(npm bin)/grpc_tools_node_protoc"
PROTOC_GEN_TS_PATH="$(npm bin)/protoc-gen-ts"
PROTOC_GEN_GRPC_PATH="$(npm bin)/grpc_tools_node_protoc_plugin"

# rm -rf $LIB_OUT_DIR/libs/proto

$PROTOC \
    -I="$DATASTORE_IN_DIR" \
    --plugin=./node_modules/.bin/protoc-gen-ts_proto \
    --ts_proto_opt=nestJs=true,snakeToCamel=false,lowerCaseServiceMethods=false \
    --ts_proto_out="$LIB_OUT_DIR" \
    "$DATASTORE_IN_DIR"/*.proto
