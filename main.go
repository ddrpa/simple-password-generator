package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"math/big"
	"os"
	"strconv"
	"strings"
)

const (
	LowerLetters                = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters                = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerLettersNoConfusingChar = "abcdefghijkmnpqrstuvwxyz"
	UpperLettersNoConfusingChar = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	Digits                      = "0123456789"
	Symbols                     = "~!@#$%^&*()_+-={}[]:<>?,./"
	Version                     = "20230524"
)

func generate(letters string, length int) (string, error) {
	var sequence string
	for i := 0; i < length; i++ {
		char, err := randomElement(letters)
		if err != nil {
			return "", err
		}

		if strings.Contains(sequence, char) {
			i--
			continue
		}
		sequence += char
	}
	return sequence, nil
}

func randomElement(s string) (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(s))))
	if err != nil {
		return "", err
	}
	return string(s[n.Int64()]), nil
}

func validateLengthOfWantedPassword(args []string) error {
	length, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	if length < 16 {
		return errors.New("密码长度不能小于16")
	} else if length > 41 {
		return errors.New("别开玩笑了，KDF 不一定能转换到更大的集合里面去")
	}
	return nil
}

func main() {
	parser := argparse.NewParser("pg", "Generate strong password, version "+Version)
	number := parser.Int("n", "num", &argparse.Options{Required: false, Default: 5, Help: "生成数量"})
	length := parser.Int("l", "len", &argparse.Options{Required: false, Default: 20, Help: "密码长度，在 16 到 41 之间选择一个数", Validate: validateLengthOfWantedPassword})
	allowConfusingElement := parser.Flag("c", "allow-confusing-element", &argparse.Options{Required: false, Help: "允许使用容易混淆的字符"})
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

	letters := Digits + Symbols
	if *allowConfusingElement {
		letters += LowerLetters + UpperLetters
	} else {
		letters += LowerLettersNoConfusingChar + UpperLettersNoConfusingChar
	}

	for i := 0; i < iter; i++ {
		result, _ := generate(letters, *length)
		fmt.Println(result)
	}
}
