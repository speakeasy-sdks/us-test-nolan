<!-- Start SDK Example Usage -->


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
                "provident",
                "distinctio",
                "quibusdam",
            },
            Encoding: []string{
                "nulla",
                "corrupti",
                "illum",
            },
            Files: [][]byte{
                []byte("error"),
                []byte("deserunt"),
            },
            GzUncompressedContentType: ustest.String("suscipit"),
            HiResModelName: []string{
                "magnam",
                "debitis",
            },
            OcrLanguages: []string{
                "delectus",
            },
            OutputFormat: ustest.String("tempora"),
            PdfInferTableStructure: []string{
                "molestiae",
                "minus",
            },
            Strategy: []string{
                "voluptatum",
                "iusto",
                "excepturi",
                "nisi",
            },
            XMLKeepTags: []string{
                "temporibus",
                "ab",
                "quis",
                "veritatis",
            },
        },
        UnstructuredAPIKey: ustest.String("deserunt"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->