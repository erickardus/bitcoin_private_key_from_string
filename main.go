package main

import (
  "fmt"
  "crypto/sha256"
  "encoding/hex"
  "os"
  "strings"
)

func main() {
  
    seed := os.Args[1:]
    seed_str := strings.Join(seed, " ")
    seed_bytes := []byte(seed_str)

    hash := sha256.New()
    hash.Write(seed_bytes)
    md := hash.Sum(nil)
    mdStr := hex.EncodeToString(md)

    fmt.Printf("Bitcoin Private Key: %s\n", mdStr)
}
