#!/bin/bash

slingshot run --wasm=./hello-js.wasm \
    --handler=handle \
    --input="Bob 🤓"
