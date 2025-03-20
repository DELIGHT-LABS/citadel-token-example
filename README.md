# Citadel token example

## Key generate

### Summary
Using the ed25519 key pair, the public key is registered in Citadel each port, and a token is created with the private key and used when storing in Citadel.

### Process
1. generate key
2. Issuer name(service name) and public key are share to all Citadel operators(through DELIGHT LABS).

### How to run
```
$ go run ./cmd/key-generate

private key =  49bd37f2afb2d0c71c7be5db6081144afd68e40d4a16b8b5108bbdffcd594cd3e6a76529c6fae76cdd51f2f13692fc2b1bd2cddfe29945415f64fa382796d41d
public key =  e6a76529c6fae76cdd51f2f13692fc2b1bd2cddfe29945415f64fa382796d41d
```

## Token generate

### Summary

### Process
1. Key generate
2. Input issuer, private key, expires and uuid

### How to run
```
$ go run ./cmd/token-generate test-issuer 49bd37f2afb2d0c71c7be5db6081144afd68e40d4a16b8b5108bbdffcd594cd3e6a76529c6fae76cdd51f2f13692fc2b1bd2cddfe29945415f64fa382796d41d 1h test-uuid
token :  eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0LWlzc3VlciIsImV4cCI6MTc0MjQ1Njc3MSwianRpIjoidGVzdC11dWlkIn0.s6kl_GbPHV-m1nAsFMdjgBKykJAz6t4csikwWfBZdTODv41aNa626YV83GiN5nahRlhemHgW9OlTTBecnmhYDA
```

## For test token
Issuer : test-issuer

Private key : `49bd37f2afb2d0c71c7be5db6081144afd68e40d4a16b8b5108bbdffcd594cd3e6a76529c6fae76cdd51f2f13692fc2b1bd2cddfe29945415f64fa382796d41d`

Public key : `e6a76529c6fae76cdd51f2f13692fc2b1bd2cddfe29945415f64fa382796d41d`

