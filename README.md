# github.com/speakeasy-sdks/us-test-nolan

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/us-test-nolan
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	ustestnolan "github.com/speakeasy-sdks/us-test-nolan/v2"
	"github.com/speakeasy-sdks/us-test-nolan/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/us-test-nolan/v2/pkg/models/shared"
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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [PipelineV0031](docs/sdks/pipelinev0031/README.md)

* [Build](docs/sdks/pipelinev0031/README.md#build) - Pipeline 1

### [PipelineV0](docs/sdks/pipelinev0/README.md)

* [Build](docs/sdks/pipelinev0/README.md#build) - Pipeline 1
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                  | Status Code                   | Content Type                  |
| ----------------------------- | ----------------------------- | ----------------------------- |
| sdkerrors.HTTPValidationError | 422                           | application/json              |
| sdkerrors.SDKError            | 400-600                       | */*                           |

### Example

```go
package main

import (
	"context"
	"errors"
	ustestnolan "github.com/speakeasy-sdks/us-test-nolan/v2"
	"github.com/speakeasy-sdks/us-test-nolan/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/us-test-nolan/v2/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/us-test-nolan/v2/pkg/models/shared"
	"log"
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

		var e *sdkerrors.HTTPValidationError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.unstructured.io` | None |

#### Example

```go
package main

import (
	"context"
	ustestnolan "github.com/speakeasy-sdks/us-test-nolan/v2"
	"github.com/speakeasy-sdks/us-test-nolan/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/us-test-nolan/v2/pkg/models/shared"
	"log"
	"net/http"
)

func main() {
	s := ustestnolan.New(
		ustestnolan.WithServerIndex(0),
	)

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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	ustestnolan "github.com/speakeasy-sdks/us-test-nolan/v2"
	"github.com/speakeasy-sdks/us-test-nolan/v2/pkg/models/operations"
	"github.com/speakeasy-sdks/us-test-nolan/v2/pkg/models/shared"
	"log"
	"net/http"
)

func main() {
	s := ustestnolan.New(
		ustestnolan.WithServerURL("https://api.unstructured.io"),
	)

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
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
