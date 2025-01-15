# Huggo

>[!WARNING]
> This is currently in development so things may break between a release and another without warning.

Huggo is a Go SDK for the [HuggingFace API](https://huggingface.co/docs/hub/en/api).

## Getting started

An example of usage

```go
package main

import (
    "fmt"

    "github.com/roushou/huggo"
)

func main() {
    hub := huggo.Hub("<api-key>")
    models := hub.Search.GetModels()
    fmt.Println(models)
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
