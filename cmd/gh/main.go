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

	if sealed != "" {
		if url == "" {
			url = "https://tokenizer.fly.io"
		}

		opts := []tokenizer.ClientOption{tokenizer.WithSecret(sealed, nil)}
		if auth != "" {
			opts = append(opts, tokenizer.WithAuth(auth))
		}

		t, err := tokenizer.Transport(url, opts...)
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
