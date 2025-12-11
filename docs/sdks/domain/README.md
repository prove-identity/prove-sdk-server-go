# Domain

## Overview

### Available Operations

* [V3DomainConfirmLinkRequest](#v3domainconfirmlinkrequest) - Confirm a domain link request
* [V3DomainIDRequest](#v3domainidrequest) - Get Domain Details
* [V3DomainLinkRequest](#v3domainlinkrequest) - Request a domain link
* [V3DomainLinkedRequest](#v3domainlinkedrequest) - Get the list of domains that are linked to this domain.
* [V3DomainUnlinkRequest](#v3domainunlinkrequest) - Remove a domain link or request

## V3DomainConfirmLinkRequest

Confirms a given domain link request by validating the PCID.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DomainConfirmLinkRequest" method="post" path="/v3/domain/confirm-link" -->
```go
package main

import(
	"context"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: provesdkservergo.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Domain.V3DomainConfirmLinkRequest(ctx, &components.V3DomainConfirmLinkRequest{
        Pcid: "pcid",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3DomainConfirmLinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [components.V3DomainConfirmLinkRequest](../../models/components/v3domainconfirmlinkrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.V3DomainConfirmLinkRequestResponse](../../models/operations/v3domainconfirmlinkrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3DomainIDRequest

Returns the domain details.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DomainIDRequest" method="get" path="/v3/domain/id" -->
```go
package main

import(
	"context"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: provesdkservergo.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Domain.V3DomainIDRequest(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3DomainIDResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.V3DomainIDRequestResponse](../../models/operations/v3domainidrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3DomainLinkRequest

Create a request to connect the requested domain to the domain the request is made from.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DomainLinkRequest" method="post" path="/v3/domain/link" -->
```go
package main

import(
	"context"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: provesdkservergo.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Domain.V3DomainLinkRequest(ctx, &components.V3DomainLinkRequest{
        Pcid: "pcid",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3DomainLinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.V3DomainLinkRequest](../../models/components/v3domainlinkrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.V3DomainLinkRequestResponse](../../models/operations/v3domainlinkrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3DomainLinkedRequest

Returns the accepted and pending links for this domain.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DomainLinkedRequest" method="get" path="/v3/domain/linked" -->
```go
package main

import(
	"context"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: provesdkservergo.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Domain.V3DomainLinkedRequest(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3DomainLinkedResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.V3DomainLinkedRequestResponse](../../models/operations/v3domainlinkedrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3DomainUnlinkRequest

Remove a domain link or request between the requested domain and the domain the request is made from.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DomainUnlinkRequest" method="post" path="/v3/domain/unlink" -->
```go
package main

import(
	"context"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: provesdkservergo.Pointer("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.Pointer("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Domain.V3DomainUnlinkRequest(ctx, &components.V3DomainUnlinkRequest{
        PcidFrom: provesdkservergo.Pointer("pcidFrom"),
        PcidTo: provesdkservergo.Pointer("pcidTo"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3DomainUnlinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.V3DomainUnlinkRequest](../../models/components/v3domainunlinkrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.V3DomainUnlinkRequestResponse](../../models/operations/v3domainunlinkrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |