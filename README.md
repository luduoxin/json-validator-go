# json-validator-go

A json validator for golang

## Getting started
### Prerequisites
- **[Go](https://go.dev/)**: any one of the **three latest major** [releases](https://go.dev/doc/devel/release).

### Getting json-validator-go
With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/luduoxin/json-validator-go/validator"
```

Otherwise, run the following Go command to install the `json-validator-go` package:

```sh
$ go get -u github.com/luduoxin/json-validator-go
```

### using json-validator-go

```go
package main

import (
	"fmt"
	"github.com/luduoxin/json-validator-go/validator"
)

func main() {
	dataStr := `{"foo":"bar"}`
	err := validator.Valid([]byte(dataStr))
	fmt.Println(err)
}
```

## Contributing

We appreciate your help!
