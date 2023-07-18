# Password Generator

```shell
$ pg -h
usage: pg [-h|--help] [-n|--num <integer>] [-l|--len <integer>]
          [-c|--allow-confusing-element] [-s|--symbol "<value>"] [-v|--version]

          Generate strong password, version 1.2.0

Arguments:

  -h  --help                     Print help information
  -n  --num                      生成数量. Default: 5
  -l  --len                      密码长度
  -c  --allow-confusing-element  允许使用容易混淆的字符
  -s  --symbol                   自定义符号
  -v  --version                  显示版本
```

## Changelog

**Version 1.2.0**

我们注意到一些系统无法识别特殊符号集合中的部分字符，因此允许通过 `-s` 指定一个特殊符号集合。注意特殊符号集合中的重复字符会被剔除。

```shell
$ pg -s "一定要这样玩是吧，好好好"
 
yxK要w0s一n是zGcD1d定WH9
u1好ZBe玩F这C3WzU6是s吧E4
7kmaAv一j样RdtQz好玩2g吧D
pqGi8r定k玩E4X9W0yRd样e
玩WYKxmD9RMEytZAGiuP7
```

## 特性

- 默认生成 5 个 20 字符长的密码
    - 密码长度至少为 16，至多不超过字符集全长
- 单个密码中每个字符只会出现一次
- 确保数字、大小写字母和特殊符号都有出现
- 使用大小写字母数字和特殊符号，默认剔除易混淆符号，如小写字母 `l`，`o`，大写字母 `I`，`O`
- 支持用户使用自己的特殊符号集合

## 怎样编译

详见 `release.sh`
