<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	ustestnolan "github.com/speakeasy-sdks/us-test-nolan/v3"
	"github.com/speakeasy-sdks/us-test-nolan/v3/pkg/models/operations"
	"github.com/speakeasy-sdks/us-test-nolan/v3/pkg/models/shared"
	"log"
	"net/http"
)

func main() {
	s := ustestnolan.New()

	ctx := context.Background()
	res, err := s.PipelineV0031.Build(ctx, operations.Pipeline1GeneralV0031GeneralPostRequest{
		PipelineBodyV0031: &shared.PipelineBodyV0031{
			Coordinates: []string{
				"string",
			},
			Encoding: []string{
				"string",
			},
			Files: []shared.PipelineBodyV0031Files{
				shared.PipelineBodyV0031Files{
					Content:  []byte("0x591E0BfdA7"),
					FileName: "cab_touring_henry.mpg4",
				},
			},
			HiResModelName: []string{
				"string",
			},
			OcrLanguages: []string{
				"string",
			},
			PdfInferTableStructure: []string{
				"string",
			},
			Strategy: []string{
				"string",
			},
			XMLKeepTags: []string{
				"string",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->