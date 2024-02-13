<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	ustestnolan "github.com/speakeasy-sdks/us-test-nolan/v3"
	"github.com/speakeasy-sdks/us-test-nolan/v3/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
	s := ustestnolan.New()

	ctx := context.Background()
	res, err := s.PipelineV0031.Build(ctx, operations.Pipeline1GeneralV0031GeneralPostRequest{})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->