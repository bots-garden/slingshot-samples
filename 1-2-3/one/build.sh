#!/bin/bash
tinygo build -scheduler=none --no-debug \
    -o one.wasm \
    -target wasi main.go

ls -lh *.wasm
