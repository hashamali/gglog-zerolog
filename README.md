## gslzl: gsl Zerolog Implementation
[![godoc](https://godoc.org/github.com/hashamali/gslzl?status.svg)](http://godoc.org/github.com/hashamali/gslzl)
[![sec](https://img.shields.io/github/workflow/status/hashamali/gslzl/security?label=security&style=flat-square)](https://github.com/hashamali/gslzl/actions?query=workflow%3Asecurity)
[![go-report](https://goreportcard.com/badge/github.com/hashamali/gslzl)](https://goreportcard.com/report/github.com/hashamali/gslzl)
[![license](https://badgen.net/github/license/hashamali/gslzl)](https://opensource.org/licenses/MIT)

A [Zerolog](https://github.com/rs/zerolog) implementation of [gsl](https://github.com/hashamali/gsl).


## Usage

```
import (
    "github.com/hashamali/gslzl"
    "github.com/rs/zerolog"
)

func main() {
    l := gslzl.New(zerolog.New(os.Stdout))
    l.Info("Ready!")
}
```
