## stlog - Zerolog Implementation

A [Zerolog](https://github.com/rs/zerolog) implementation of [stlog](https://github.com/hashamali/stlog).


## Usage

```
import (
    "github.com/hashamali/stlogzl"
    "github.com/rs/zerolog"
)

func main() {
    l := stlogzl.New(zerolog.New(os.Stdout))
    l.Info("Ready!")
}
```