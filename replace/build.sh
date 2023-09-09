#!/bin/bash
tinygo build -scheduler=none --no-debug \
    -o replace.wasm \
    -target wasi main.go

ls -lh *.wasm
