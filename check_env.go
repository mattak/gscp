package main

import (
	"fmt"
	"os"
)

func CheckEnvironmentValues() {
	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		fmt.Println("The GOOGLE_APPLICATION_CREDENTIALS environment variable is not set")
		os.Exit(1)
	}
}
