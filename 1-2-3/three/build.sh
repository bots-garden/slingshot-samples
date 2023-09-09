#!/bin/bash
tinygo build -scheduler=none --no-debug \
    -o three.wasm \
    -target wasi main.go

ls -lh *.wasm
