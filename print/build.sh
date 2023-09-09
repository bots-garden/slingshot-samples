#!/bin/bash
tinygo build -scheduler=none --no-debug \
    -o print.wasm \
    -target wasi main.go

ls -lh *.wasm
