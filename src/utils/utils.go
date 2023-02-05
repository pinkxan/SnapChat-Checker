package utils

import (
	"bufio"
	"math/rand"
	"os"
)

func RandomString(n int) string {
	var letters = []rune("0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func OpenFile(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func WriteFile(name, content string) bool {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
		return false
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		panic(err)
		return false
	}

	return true
}
