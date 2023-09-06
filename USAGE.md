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
                "corrupti",
            },
            Encoding: []string{
                "provident",
            },
            Files: [][]byte{
                []byte("distinctio"),
            },
            GzUncompressedContentType: ustest.String("quibusdam"),
            HiResModelName: []string{
                "unde",
            },
            OcrLanguages: []string{
                "nulla",
            },
            OutputFormat: ustest.String("corrupti"),
            PdfInferTableStructure: []string{
                "illum",
            },
            Strategy: []string{
                "vel",
            },
            XMLKeepTags: []string{
                "error",
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