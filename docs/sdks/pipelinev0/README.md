# PipelineV0

### Available Operations

* [Build](#build) - Pipeline 1

## Build

Pipeline 1

### Example Usage

```go
package main

import(
	"context"
	"log"
	ustestnolan "github.com/speakeasy-sdks/us-test-nolan"
	"github.com/speakeasy-sdks/us-test-nolan/pkg/models/operations"
	"github.com/speakeasy-sdks/us-test-nolan/pkg/models/shared"
)

func main() {
    s := ustestnolan.New()

    ctx := context.Background()
    res, err := s.PipelineV0.Build(ctx, operations.Pipeline1GeneralV0GeneralPostRequest{
        PipelineBodyV0: &shared.PipelineBodyV0{
            Coordinates: []string{
                "voluptatum",
            },
            Encoding: []string{
                "iusto",
            },
            Files: [][]byte{
                []byte("excepturi"),
            },
            GzUncompressedContentType: ustestnolan.String("nisi"),
            HiResModelName: []string{
                "recusandae",
            },
            OcrLanguages: []string{
                "temporibus",
            },
            OutputFormat: ustestnolan.String("ab"),
            PdfInferTableStructure: []string{
                "quis",
            },
            Strategy: []string{
                "veritatis",
            },
            XMLKeepTags: []string{
                "deserunt",
            },
        },
        UnstructuredAPIKey: ustestnolan.String("perferendis"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.Pipeline1GeneralV0GeneralPostRequest](../../models/operations/pipeline1generalv0generalpostrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.Pipeline1GeneralV0GeneralPostResponse](../../models/operations/pipeline1generalv0generalpostresponse.md), error**

