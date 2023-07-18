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
	Version                     = "1.2.0"
	LowerLetters                = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters                = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerLettersNoConfusingChar = "abcdefghijkmnpqrstuvwxyz"
	UpperLettersNoConfusingChar = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	Digits                      = "0123456789"
)

// Symbols 如果用户提供了自定义符号集合，会替换该变量
var Symbols = "~!@#$%^&*()_+-={}[]:<>?,./"

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

func main() {
	parser := argparse.NewParser("pg", "Generate strong password, version "+Version)
	number := parser.Int("n", "num", &argparse.Options{Required: false, Default: 5, Help: "生成数量"})
	length := parser.Int("l", "len", &argparse.Options{Required: false, Default: 20, Help: "密码长度"})
	allowConfusingElement := parser.Flag("c", "allow-confusing-element", &argparse.Options{Required: false, Help: "允许使用容易混淆的字符"})
	customSymbol := parser.String("s", "symbol", &argparse.Options{Required: false, Help: "自定义符号"})
	showVersion := parser.Flag("v", "version", &argparse.Options{Required: false, Help: "显示版本"})

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
	if *customSymbol != "" {
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

	for i := 0; i < iter; i++ {
		result, _ := generate(lettersAsRuneArray, *length)
		if verifyPasswordHasAllRequiredChar(result) {
			fmt.Println(result)
		} else {
			i--
		}
	}
}
