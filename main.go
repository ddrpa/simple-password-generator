package main

import (
	"crypto/rand"
	"fmt"
	"github.com/akamensky/argparse"
	"math/big"
	"os"
	"strings"
)

const (
	Version                     = "1.4.0"
	LowerLetters                = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters                = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerLettersNoConfusingChar = "abcdefghijkmnpqrstuvwxyz"
	UpperLettersNoConfusingChar = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	Digits                      = "0123456789"
)

// Symbols 如果用户提供了自定义符号集合，会替换该变量
var Symbols = "_~!@#$%^&*()-=<>,.?;:|+{}[]/"

// 一些系统可能仅支持一部分特殊字符，可以使用如下所示的 JavaScript 代码获得过滤后的特殊字符集合
// [...Symbols].filter(ch => [...AllowedSymbols].includes(ch)).join('')

// SymbolsFlavorMySQL8 MySQL 8 支持在密码中使用的特殊符号
var SymbolsFlavorMySQL8 = "_~!@#$%^&*()-=<>,.?;:|"

// SymbolsFlavorRedis Redis 支持在密码中使用的特殊符号
var SymbolsFlavorRedis = "_~!@#$%^&*()-=<>,.?;|+{}[]"

func contains(elems []rune, v rune) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

func generate(letters []rune, length int) (string, error) {
	var sequence []rune
	collectionSize := big.NewInt(int64(len(letters)))
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, collectionSize)
		if err != nil {
			i--
			continue
		}

		char := letters[n.Int64()]

		// if sequence contains char, then i--
		if contains(sequence, char) {
			i--
			continue
		}

		sequence = append(sequence, char)
	}
	return string(sequence), nil
}

func validateLengthOfWantedPassword(collectionLength int, wantedLength int) error {
	if wantedLength < 16 {
		return fmt.Errorf("密码长度不能小于16")
	}
	if wantedLength > collectionLength {
		return fmt.Errorf("字符集不足以创建指定长度的密码")
	}
	return nil
}

func removeDuplicateSymbol(s string) (string, error) {
	var result string
	for _, c := range s {
		if strings.ContainsRune(LowerLetters, c) ||
			strings.ContainsRune(UpperLetters, c) ||
			strings.ContainsRune(Digits, c) ||
			strings.ContainsRune(result, c) {
			continue
		}
		result += string(c)
	}
	if len(result) == 0 {
		return "", fmt.Errorf("无法筛选出有效的自定义符号")
	} else {
		return result, nil
	}
}

/**
 * @description: 验证密码强度是否符合常见系统的强密码检测规则
 */
func verifyPasswordHasAllRequiredChar(password string) bool {
	return strings.ContainsAny(password, LowerLetters) &&
		strings.ContainsAny(password, UpperLetters) &&
		strings.ContainsAny(password, Digits) &&
		strings.ContainsAny(password, Symbols)
}

/**
 * @description: 验证密码强度是否符合常见系统的强密码检测规则，此外还可以套用一系列自定义规则
 */
func verifyPasswordHasAllRequiredCharWithCustomRules(password string, rules ...func(password string) bool) bool {
	if !verifyPasswordHasAllRequiredChar(password) {
		return false
	}
	for _, rule := range rules {
		if !rule(password) {
			return false
		}
	}
	return true
}

func main() {
	parser := argparse.NewParser("pg", "Generate strong password, version "+Version)
	number := parser.Int("n", "num", &argparse.Options{Required: false, Default: 5, Help: "生成数量"})
	length := parser.Int("l", "len", &argparse.Options{Required: false, Default: 20, Help: "密码长度"})
	allowConfusingElement := parser.Flag("c", "allow-confusing-element", &argparse.Options{Required: false, Help: "允许使用容易混淆的字符"})
	flavor := parser.String("f", "flavor", &argparse.Options{Required: false, Help: "使用特定系统支持的特殊符号集合，目前支持 mysql8 和 redis"})
	customSymbol := parser.String("s", "symbol", &argparse.Options{Required: false, Help: "自定义特殊符号集合"})
	showVersion := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "版本"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		return
	}

	if *showVersion {
		fmt.Println("password generator version " + Version)
		return
	}

	iter := *number
	if iter < 0 {
		iter = 5
	}

	letters := Digits
	if *flavor == "mysql8" {
		Symbols = SymbolsFlavorMySQL8
	} else if *flavor == "redis" {
		Symbols = SymbolsFlavorRedis
	} else if *customSymbol != "" {
		customSymbolChar, err := removeDuplicateSymbol(*customSymbol)
		Symbols = customSymbolChar
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if *allowConfusingElement {
		letters += Symbols + LowerLetters + UpperLetters
	} else {
		letters += Symbols + LowerLettersNoConfusingChar + UpperLettersNoConfusingChar
	}
	lettersAsRuneArray := []rune(letters)

	err = validateLengthOfWantedPassword(len(lettersAsRuneArray), *length)
	if err != nil {
		fmt.Println(err)
		return
	}

	var verifyPassword func(password string) bool
	if *flavor == "redis" {
		verifyPassword = func(password string) bool {
			return verifyPasswordHasAllRequiredCharWithCustomRules(password, func(password string) bool {
				// 在 Redis 6.2.12 之前，密码必须以大写或小写字母开头
				return strings.ContainsAny(password[:1], LowerLetters+UpperLetters)
			})
		}
	} else {
		verifyPassword = func(password string) bool {
			return verifyPasswordHasAllRequiredChar(password)
		}
	}

	for i := 0; i < iter; i++ {
		result, _ := generate(lettersAsRuneArray, *length)
		if verifyPassword(result) {
			fmt.Println(result)
		} else {
			i--
		}
	}
}
