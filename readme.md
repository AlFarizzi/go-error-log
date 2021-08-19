```go
package main

import errjson "github.com/AlFarizzi/go-error-log"

func main() {
	errjson.WriteError("err.json", "Ini Error")
}

```