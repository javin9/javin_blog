package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 示例字符串
	str := "Hello, 世界 🌍"

	fmt.Println("原始字符串:", str)
	fmt.Println("字节长度 (len):", len(str))
	fmt.Println("字符长度 (len([]rune)):", len([]rune(str)))
	fmt.Println("UTF8 字符数 (utf8.RuneCountInString):", utf8.RuneCountInString(str))

	fmt.Println("\n逐个字符遍历:")
	for i, r := range str {
		fmt.Printf("位置 %d: %c (Unicode: %U)\n", i, r, r)
	}

	fmt.Println("\n逐个字节遍历:")
	for i, b := range []byte(str) {
		fmt.Printf("字节 %d: %d (0x%02x)\n", i, b, b)
	}

	// 演示中文字符
	chinese := "你好世界"
	fmt.Println("\n中文字符串:", chinese)
	fmt.Println("字节长度:", len(chinese))
	fmt.Println("字符长度:", len([]rune(chinese)))

	// 演示 emoji
	emoji := "🚀🎉💻"
	fmt.Println("\nEmoji 字符串:", emoji)
	fmt.Println("字节长度:", len(emoji))
	fmt.Println("字符长度:", len([]rune(emoji)))
}
