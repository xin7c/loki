echo "交叉编译linux loki"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
echo "复制到odin-loki"
cp loki ../odin-loki