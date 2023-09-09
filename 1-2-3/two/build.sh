#!/bin/bash
tinygo build -scheduler=none --no-debug \
    -o two.wasm \
    -target wasi main.go

ls -lh *.wasm
