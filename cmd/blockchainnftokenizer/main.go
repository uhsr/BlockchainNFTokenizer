// cmd/blockchainnftokenizer/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftokenizer/internal/blockchainnftokenizer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftokenizer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
