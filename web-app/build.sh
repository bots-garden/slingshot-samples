#!/bin/bash
tinygo build -scheduler=none --no-debug \
    -o webapp.wasm \
    -target wasi main.go

ls -lh *.wasm
