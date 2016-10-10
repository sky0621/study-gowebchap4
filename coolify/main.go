package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	duplicateVowel bool = true
	removeVowel    bool = false
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := []byte(s.Text())
		if randBool() {
			var vI int = -1
			// 標準入力から渡された文字列に母音が含まれていたら、50%の確率でその母音のある位置を保持
			for i, char := range word {
				switch char {
				case 'a', 'e', 'i', 'u', 'o', 'A', 'E', 'I', 'U', 'O':
					if randBool() {
						vI = i
					}
				}
			}
			// 母音が含まれており、かつ、50%の確率に合致した場合
			if vI >= 0 {
				// さらに50%の確率で、重複指定ありか除去指定ありかによって、文字列の母音を重複させるか除去するか決定
				switch randBool() {
				case duplicateVowel:
					word = append(word[:vI+1], word[vI:]...)
				case removeVowel:
					word = append(word[:vI], word[vI+1:]...)
				}
			}
		}
		fmt.Println(string(word))
	}
}
