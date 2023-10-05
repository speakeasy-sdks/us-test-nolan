<!-- Start SDK Example Usage -->


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
                "Practical",
            },
            Encoding: []string{
                "Gasoline",
            },
            Files: [][]byte{
                []byte("[eAhpDJhn'"),
            },
            GzUncompressedContentType: ustestnolan.String("Shoes"),
            HiResModelName: []string{
                "what",
            },
            OcrLanguages: []string{
                "whether",
            },
            OutputFormat: ustestnolan.String("Direct"),
            PdfInferTableStructure: []string{
                "gee",
            },
            Strategy: []string{
                "Metal",
            },
            XMLKeepTags: []string{
                "Avon",
            },
        },
        UnstructuredAPIKey: ustestnolan.String("Intranet synthesizing HTTP"),
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