## slzl - sl Zerolog Implementation

A [Zerolog](https://github.com/rs/zerolog) implementation of [sl](https://github.com/hashamali/sl).


## Usage

```
import (
    "github.com/hashamali/slzl"
    "github.com/rs/zerolog"
)

func main() {
    l := slzl.New(zerolog.New(os.Stdout))
    l.Info("Ready!")
}
```
