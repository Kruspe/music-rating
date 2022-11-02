#!/usr/bin/env bash

set -eu -o pipefail

SCRIPT_PATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
BINARY_PATH="$SCRIPT_PATH/install"

DOWNLOAD_URL='https://s3.eu-central-1.amazonaws.com/dynamodb-local-frankfurt/dynamodb_local_latest.tar.gz'

if [ -d "$BINARY_PATH" ]; then
    echo "dynamodb is already downloaded to $BINARY_PATH"
    exit 0
fi

cd "$SCRIPT_PATH"
mkdir "$BINARY_PATH"
echo "Downloading DynamoDB. This may take a while if you're on a slow connection."
curl -fsSL "$DOWNLOAD_URL" | tar -xz --directory "$BINARY_PATH"

if [[ -z "${CODEBUILD_BUILD_IMAGE:-}" && -z "${CI:-}" ]]; then
    cd "$BINARY_PATH"
    curl -fsSL -o libsqlite4java-osx.dylib.arm64 'https://search.maven.org/remotecontent?filepath=io/github/ganadist/sqlite4java/libsqlite4java-osx-arm64/1.0.392/libsqlite4java-osx-arm64-1.0.392.dylib'
    mv DynamoDBLocal_lib/libsqlite4java-osx.dylib libsqlite4java-osx.dylib.x86_64
    lipo -create -output libsqlite4java-osx.dylib.fat libsqlite4java-osx.dylib.x86_64 libsqlite4java-osx.dylib.arm64
    mv libsqlite4java-osx.dylib.fat DynamoDBLocal_lib/libsqlite4java-osx.dylib
fi