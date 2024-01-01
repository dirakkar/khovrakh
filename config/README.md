# `khovrakh/config`

## Install

```sh
go get -u github.com/dirakkar/khovrakh/config
```

## Example

```go
import (
    "time"
    /* ... */

    "github.com/dirakkar/khovrakh/config"
)

var configPath = "config.json"

type Config struct {
    Address   string `cfg:"address"`
    BodyLimit int    `cfg:"body_limit"`

    ShutdownTimeout time.Duration `cfg:"shutdown_timeout"`
}

func main() {
    cfg, err := config.New[Config](configPath)
    if err != nil { /* ... */ }

    // ...
}
```
```json
{
  "address": "localhost:8000",
  "body_limit": "5kb",
  "shutdown_timeout": "5s"
}
```
