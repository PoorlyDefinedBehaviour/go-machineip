# Get the host IP

Contains a function to get an ip that can be used like localhost.

# Motivation

If you have a rest api running in a machine executor on CircleCI
and try to make a http request to localhost, it won't work.

You can use this package to get an ip that you can call just like
localhost.

# Install

```terminal
go get github.com/poorlydefinedbehaviour/go-machineip
```

# Example

This example assumes there is a server running on port 5000.

```go
import (
  "testing"
  "http"
  "github.com/poorlydefinedbehaviour/go-machineip"
)

func Test_MyTest(t *testing.T) {
	t.Parallel()

	endpoint := fmt.Sprintf("%s:5000", machineip.IP())

	response, err := http.Get(endpoint)
	// ...
}
```
