#!/usr/bin/env bash

set -eo pipefail

echo "Generating gogo proto code"
cd proto

buf generate --template buf.gen.gogo.yaml

cd ..

# move proto files to the right places
cp -r github.com/cosmos/interchain-security/* ./
rm -rf github.com

# FIXME with -compat=1.18 it's not working, check why and if it's still needed
go mod tidy #-compat=1.18
