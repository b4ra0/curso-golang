package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites["Linkedin"])

	delete(websites, "Google")
	fmt.Println(websites)

	for _, value := range websites {
		fmt.Println(value)
	}
}
