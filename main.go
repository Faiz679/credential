package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type input struct {
	Env string `json:"env"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing input string")
		os.Exit(1)
	}

	var in input
	if err := json.Unmarshal([]byte(os.Args[1]), &in); err != nil {
		fmt.Println("Error parsing input JSON:", err)
		os.Exit(1)
	}

	val := os.Getenv(in.Env)
	if val == "" {
		fmt.Printf(`{"error": "Environment variable %s is not set"}`, in.Env)
		os.Exit(1)
	}

	fmt.Printf(`{"env": {"%s": "%s"}}`, in.Env, val)
}
