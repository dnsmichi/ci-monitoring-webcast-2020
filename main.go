package main
/*
import (
	"fmt"
	"os"
)*/

func main() {
	hello_from := os.Getenv("HELLO_FROM")

	if hello_from == "" {
		hello_from = "our GitLab meetup"
	}

	fmt.Println("Hello from " + hello_from + "!")
	fmt.Println("Today we learn about GitLab Best Practices and CI/CD magic :)")

	fmt.Println("");

	fmt.Print(GetTanuki(true))

	fmt.Println("");
	fmt.Println("Join us at https://www.everyonecancontribute.com")
}
