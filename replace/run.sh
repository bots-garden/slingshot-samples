# Replace "hello" by "hey"

text64=$(cat source.txt| base64)

# Slingshot Arguments
args=$(cat << EOF
{
	"find": "hello",
	"replace": "hey",
	"text": "${text64}"
}
EOF
)

result=$(slingshot run \
  --wasm=./replace.wasm \
  --handler=callHandler \
  --input="${args}"
)
echo "${result}" > target.txt
# Exit code
echo $?
