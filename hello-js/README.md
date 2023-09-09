# Hello-JS

> WIP 🚧 (it's a preview of the Slingshot JS-PDK)

## Install dependencies

```bash
npm install

export TAG="v0.5.0"
export ARCH="x86_64"
export  OS="linux"
curl -L -O "https://github.com/extism/js-pdk/releases/download/$TAG/extism-js-$ARCH-$OS-$TAG.gz"
gunzip extism-js*.gz
sudo mv extism-js-* /usr/local/bin/extism-js
chmod +x /usr/local/bin/extism-js
```

## Build 

```bash
npm run build
```

## Run 

```bash
slingshot run --wasm=./hello-js.wasm \
    --handler=handle \
    --input="Bob 🤓"
```
