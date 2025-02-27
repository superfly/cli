package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cli/cli/v2/internal/ghcmd"
	"github.com/superfly/tokenizer"
)

func main() {
	url := os.Getenv("GH_TOKENIZER_URL")
	sealed := os.Getenv("GH_TOKENIZER_SEALED_TOKEN")
	auth := os.Getenv("GH_TOKENIZER_AUTH")

	if sealed != "" && auth != "" {
		if url == "" {
			url = "https://tokenizer.fly.io"
		}

		t, err := tokenizer.Transport(url, tokenizer.WithSecret(sealed, nil), tokenizer.WithAuth(auth))
		if err != nil {
			fmt.Println("Failed to inject tokenizer transport")
			fmt.Println(err)
			os.Exit(1)
		}

		http.DefaultTransport = t
	}

	code := ghcmd.Main()
	os.Exit(int(code))
}
