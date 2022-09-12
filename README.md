# PVTA-go

Go API wrapper for the PVTA (Pioneer Valley Transit Authority) API.

## Getting Started

Install the package with:

```bash
go get github.com/jackmerrill/pvta-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/jackmerrill/pvta-go"
)

func main() {
    // Create a new PVTA client
    client := pvta.NewClient()

    // Get all routes
    routes, err := client.GetRoutes()
    if err != nil {
        panic(err)
    }

    // Print the first route
    fmt.Println(routes[0])
}
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MPL-2.0](https://choosealicense.com/licenses/mpl-2.0/)