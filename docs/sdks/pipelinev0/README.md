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
	"github.com/speakeasy-sdks/us-test-nolan"
	"github.com/speakeasy-sdks/us-test-nolan/pkg/models/operations"
	"github.com/speakeasy-sdks/us-test-nolan/pkg/models/shared"
)

func main() {
    s := ustest.New()

    ctx := context.Background()
    res, err := s.PipelineV0.Build(ctx, operations.Pipeline1GeneralV0GeneralPostRequest{
        PipelineBodyV0: &shared.PipelineBodyV0{
            Coordinates: []string{
                "suscipit",
            },
            Encoding: []string{
                "iure",
            },
            Files: [][]byte{
                []byte("magnam"),
            },
            GzUncompressedContentType: ustest.String("debitis"),
            HiResModelName: []string{
                "ipsa",
            },
            OcrLanguages: []string{
                "delectus",
            },
            OutputFormat: ustest.String("tempora"),
            PdfInferTableStructure: []string{
                "suscipit",
            },
            Strategy: []string{
                "molestiae",
            },
            XMLKeepTags: []string{
                "minus",
            },
        },
        UnstructuredAPIKey: ustest.String("placeat"),
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

