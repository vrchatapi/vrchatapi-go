#!/bin/bash

npm install @openapitools/openapi-generator-cli

rm *.ts -rf

./node_modules/\@openapitools/openapi-generator-cli/main.js generate \
-g go \
--additional-properties=isGoSubmodule=true,packageName=vrchatapi \
--git-user-id=vrchatapi \
--git-repo-id=vrchatapi-go \
-o . \
-i https://raw.githubusercontent.com/vrchatapi/specification/gh-pages/openapi.yaml \
--http-user-agent="vrchatapi-go"

# Remove messily pasted markdown at top of every file
for i in *.go; do
    sed -i '/VRChat API Banner/d' $i
done

go get
go build