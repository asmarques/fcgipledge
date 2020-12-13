# fcgipledge

[![Go Reference](https://pkg.go.dev/badge/github.com/asmarques/fcgipledge.svg)](https://pkg.go.dev/github.com/asmarques/fcgipledge)
![Build Status](https://github.com/asmarques/fcgipledge/workflows/CI/badge.svg)

Create a FastCGI server in Go restricted using OpenBSD's [`pledge`](https://man.openbsd.org/pledge.2).

## Requirements

- OpenBSD 5.9 or later

No restrictions will be applied to the created server when run on any other OS.

## Installation

```bash
go get github.com/asmarques/fcgipledge
```

## Usage

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/asmarques/fcgipledge"
)

func main() {
	path := "/run/fcgi.sock"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world")
	})

	fcgipledge.ListenAndServe(path, nil, nil)
}
```

To forward requests from OpenBSD's [`httpd`](https://man.openbsd.org/httpd.8) to a running server just add a reference to the socket in [`httpd.conf`](https://man.openbsd.org/httpd.conf.5):

```
...

server "www.example.com" {
        ...
        fastcgi socket "/run/fcgi.sock"
}

...
```

## License

[MIT](LICENSE)
