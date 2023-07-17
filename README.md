# Password Generator

```shell
$ pg -h
usage: pg [-h|--help] [-n|--num <integer>] [-l|--len <integer>]
          [-c|--allow-confusing-element]

          生成安全的密码

Arguments:

  -h  --help                     Print help information
  -n  --num                      生成数量. Default: 5
  -l  --len                      密码长度，在 16 到 41 之间选择一个数. Default: 20
  -c  --allow-confusing-element  允许使用容易混淆的字符
```

## 特性

- 默认生成 5 个 20 字符长的密码
- 单个密码中每个字符只会出现一次
- 确保数字、大小写字母和特殊符号都有出现
- 使用大小写字母数字和特殊符号，默认剔除易混淆符号，如小写字母 `l`，`o`，大写字母 `I`，`O`

## 怎样编译

详见 `release.sh`
