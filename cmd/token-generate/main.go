package main

import (
	"os"

	"github.com/DELIGHT-LABS/citadel-token-example/token"
)

func main() {
	if len(os.Args) != 5 {
		panic("Usage: token-generate <issuer> <private_key> <expires> <uuid>")
	}

	service := token.NewTokenService(os.Args[1], os.Args[2], os.Args[3])
	token, err := service.GenerateToken(os.Args[4])
	if err != nil {
		panic(err)
	}

	println(`token : `, token)
}
