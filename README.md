# decodestream
decode a big JSON to a stream

Example:

```go
package main

import (
	"log"
	"os"
    "github.com/doITmagic/decodestream"
)

func main() {
	stream := decodestream.NewJSONStream()
	go func() {
		for data := range stream.Watch() {
			if data.Error != nil {
				log.Println(data.Error)
			}
			print(data.Data.(map[string]interface{}))
		}
	}()

	// Open file to read.
	file, err := os.Open("ports.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stream.Start(file)
}

```
To run test use 
```bash
go test -v
```

To get test coverage use
```bash
 ginkgo -coverpkg=./... -r
```

Run example
```bash
cd example
go run .
```
