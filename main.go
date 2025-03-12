package main

import (
	"fmt"
	"url-shortnere/internal/utils"
)

func main() {

	result := utils.ShortenURLHash("https://chatgpt.com/c/67d0f70c-f330-8002-ab1b-316725e436c4")
	fmt.Println(result)

}
