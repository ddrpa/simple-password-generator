# Password Generator

```shell
$ pg -h
usage: pg [-h|--help] [-n|--num <integer>] [-l|--len <integer>]
          [-c|--allow-confusing-element]

          生成安全的密码

Arguments:

  -h  --help                     Print help information
  -n  --num                      生成数量. Default: 5
  -l  --len                      密码长度，在 16 到 41 之间选择一个数. Default: 16
  -c  --allow-confusing-element  允许使用容易混淆的字符
```

## 特性

- 默认生成5个16字符长的密码，且无重复字符
- 使用大小写字母数字和特殊符号，默认剔除易混淆符号，如小写字母 `l`，`o`，大写字母 `I`，`O`
- 使用 `-h` 参数可查看详细选项，如 `./pg -h`

## 怎样编译

因为我只在一台 Macbook 上有 golang 环境，以下使用了交叉编译，如果你不需要跨平台，可以直接使用 `go build`

for Windows user

```shell
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o target/pg.exe main.go
```

for macOS user

```shell
# x86_64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o target/pg main.go
# apple silicon
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o target/pg main.go
```

for Linux user

```shell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o target/pg main.go
```
