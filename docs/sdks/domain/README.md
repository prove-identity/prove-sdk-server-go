# Domain
(*Domain*)

## Overview

### Available Operations

* [V3DomainID](#v3domainid) - Get Domain Details
* [V3DomainLink](#v3domainlink) - # Create a request to connect the requested domain to the domain the request is made from.
* [V3DomainLinked](#v3domainlinked) - Get the list of domains that are linked to this domain.
* [V3DomainUnlink](#v3domainunlink) - # Remove a domain link or request.

## V3DomainID

Returns the domain details.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DomainID" method="post" path="/v3/domain/id" -->
```go
package main

import(
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Domain.V3DomainID(ctx, nil)
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
| `request`                                                | [string](../../.md)                                      | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.V3DomainIDResponse](../../models/operations/v3domainidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3DomainLink

# Create a request to connect the requested domain to the domain the request is made from.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DomainLink" method="post" path="/v3/domain/link" -->
```go
package main

import(
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Domain.V3DomainLink(ctx, &components.V3DomainLinkRequest{
        Pcid: provesdkservergo.String("pcid"),
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

**[*operations.V3DomainLinkResponse](../../models/operations/v3domainlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3DomainLinked

Returns the accepted and pending links for this domain.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DomainLinked" method="get" path="/v3/domain/linked" -->
```go
package main

import(
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Domain.V3DomainLinked(ctx)
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

**[*operations.V3DomainLinkedResponse](../../models/operations/v3domainlinkedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3DomainUnlink

# Remove a domain link or request.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DomainUnlink" method="post" path="/v3/domain/unlink" -->
```go
package main

import(
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := provesdkservergo.New(
        provesdkservergo.WithSecurity(components.Security{
            ClientID: provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
            ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
        }),
    )

    res, err := s.Domain.V3DomainUnlink(ctx, &components.V3DomainUnlinkRequest{
        PcidFrom: provesdkservergo.String("pcidFrom"),
        PcidTo: provesdkservergo.String("pcidTo"),
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

**[*operations.V3DomainUnlinkResponse](../../models/operations/v3domainunlinkresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |