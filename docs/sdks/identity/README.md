# Identity

## Overview

### Available Operations

* [V3BatchGetIdentities](#v3batchgetidentities) - Batch Get Identities
* [V3EnrollIdentity](#v3enrollidentity) - Enroll Identity
* [V3BatchEnrollIdentities](#v3batchenrollidentities) - Batch Enroll Identities
* [V3DisenrollIdentity](#v3disenrollidentity) - Disenroll Identity
* [V3GetIdentity](#v3getidentity) - Get Identity
* [V3ActivateIdentity](#v3activateidentity) - Activate Identity
* [V3CrossDomainIdentity](#v3crossdomainidentity) - Cross Domain Identity
* [V3DeactivateIdentity](#v3deactivateidentity) - Deactivate Identity
* [V3GetIdentitiesByPhoneNumber](#v3getidentitiesbyphonenumber) - Get Identities By Phone Number

## V3BatchGetIdentities

Return a list of all identities you have enrolled in Identity Manager.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3BatchGetIdentities" method="get" path="/v3/identity" -->
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

    res, err := s.Identity.V3BatchGetIdentities(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3BatchGetIdentitiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                          |
| `clientRequestID`                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | A client-generated unique ID for a specific session. This can be used to identify specific requests. The format of this ID is defined by the client - Prove recommends using a GUID, but any format can be accepted. Do not include Personally Identifiable Information (PII) in this field. |
| `limit`                                                                                                                                                                                                                                                                                      | **int64*                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | The maximum number of identities to return per call. Default value is 100.                                                                                                                                                                                                                   |
| `startKey`                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | The pagination token for the GET /v3/identity API. Use this to retrieve the next page of results after a previous call to GET /v3/identity. This token is returned as lastKey in the GET /v3/identity API response - pass it in directly as startKey to get the next page of results.        |
| `showInactive`                                                                                                                                                                                                                                                                               | **bool*                                                                                                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | Whether to show identities associated with the current client that are currently marked as inactive. Default value is false.                                                                                                                                                                 |
| `opts`                                                                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                                                                |

### Response

**[*operations.V3BatchGetIdentitiesResponse](../../models/operations/v3batchgetidentitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3EnrollIdentity

Enrolls a single customer for monitoring using their phone number and unique identifier.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3EnrollIdentity" method="post" path="/v3/identity" -->
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

    res, err := s.Identity.V3EnrollIdentity(ctx, &components.V3EnrollIdentityRequest{
        ClientCustomerID: provesdkservergo.Pointer("e0f78bc2-f748-4eda-9d29-d756844507fc"),
        ClientRequestID: provesdkservergo.Pointer("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        DeviceID: provesdkservergo.Pointer("bf9ea15d-7dfa-4bb4-a64c-6c26b53472fc"),
        PhoneNumber: "2001001695",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3EnrollIdentityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.V3EnrollIdentityRequest](../../models/components/v3enrollidentityrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.V3EnrollIdentityResponse](../../models/operations/v3enrollidentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3BatchEnrollIdentities

Enrolls multiple customers in a single request for efficient bulk operations (up to 100).

### Example Usage

<!-- UsageSnippet language="go" operationID="V3BatchEnrollIdentities" method="post" path="/v3/identity/batch" -->
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

    res, err := s.Identity.V3BatchEnrollIdentities(ctx, &components.V3BatchEnrollIdentitiesRequest{
        ClientRequestID: provesdkservergo.Pointer("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        Items: []components.IdentityItem{
            components.IdentityItem{
                ClientName: "\"Client A\"",
                IdentityID: "\"e0f78bc2-f748-4eda-9d29-d756844507fc\"",
                Pcid: "\"12345\"",
            },
            components.IdentityItem{
                ClientName: "\"Client A\"",
                IdentityID: "\"e0f78bc2-f748-4eda-9d29-d756844507fc\"",
                Pcid: "\"12345\"",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3BatchEnrollIdentitiesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [components.V3BatchEnrollIdentitiesRequest](../../models/components/v3batchenrollidentitiesrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.V3BatchEnrollIdentitiesResponse](../../models/operations/v3batchenrollidentitiesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3DisenrollIdentity

Disenrolls an identity from Identity Manager. If you wish to monitor in future, re-enrollment of that identity is required.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DisenrollIdentity" method="delete" path="/v3/identity/{identityId}" -->
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

    res, err := s.Identity.V3DisenrollIdentity(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3DisenrollIdentityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                          |
| `identityID`                                                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                           | A Prove-generated unique ID for a specific identity.                                                                                                                                                                                                                                         |
| `clientRequestID`                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | A client-generated unique ID for a specific session. This can be used to identify specific requests. The format of this ID is defined by the client - Prove recommends using a GUID, but any format can be accepted. Do not include Personally Identifiable Information (PII) in this field. |
| `opts`                                                                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                                                                |

### Response

**[*operations.V3DisenrollIdentityResponse](../../models/operations/v3disenrollidentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3GetIdentity

Return details of an identity given the identity ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3GetIdentity" method="get" path="/v3/identity/{identityId}" -->
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

    res, err := s.Identity.V3GetIdentity(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetIdentityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                          |
| `identityID`                                                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                           | A unique Prove-generated identifier for the enrolled identity.                                                                                                                                                                                                                               |
| `clientRequestID`                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | A client-generated unique ID for a specific session. This can be used to identify specific requests. The format of this ID is defined by the client - Prove recommends using a GUID, but any format can be accepted. Do not include Personally Identifiable Information (PII) in this field. |
| `opts`                                                                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                                                                |

### Response

**[*operations.V3GetIdentityResponse](../../models/operations/v3getidentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3ActivateIdentity

Sets an identity as active for monitoring.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3ActivateIdentity" method="post" path="/v3/identity/{identityId}/activate" -->
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

    res, err := s.Identity.V3ActivateIdentity(ctx, "<id>", &components.V3ActivateIdentityRequest{
        ClientRequestID: provesdkservergo.Pointer("71010d88-d0e7-4a24-9297-d1be6fefde81"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ActivateIdentityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                     | Type                                                                                          | Required                                                                                      | Description                                                                                   | Example                                                                                       |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ctx`                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                         | :heavy_check_mark:                                                                            | The context to use for the request.                                                           |                                                                                               |
| `identityID`                                                                                  | *string*                                                                                      | :heavy_check_mark:                                                                            | A Prove-generated unique ID for a specific identity.                                          |                                                                                               |
| `v3ActivateIdentityRequest`                                                                   | [*components.V3ActivateIdentityRequest](../../models/components/v3activateidentityrequest.md) | :heavy_minus_sign:                                                                            | N/A                                                                                           | {<br/>"clientRequestId": "71010d88-d0e7-4a24-9297-d1be6fefde81"<br/>}                         |
| `opts`                                                                                        | [][operations.Option](../../models/operations/option.md)                                      | :heavy_minus_sign:                                                                            | The options for this request.                                                                 |                                                                                               |

### Response

**[*operations.V3ActivateIdentityResponse](../../models/operations/v3activateidentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3CrossDomainIdentity

Retreives the list of identities from other linked accounts.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3CrossDomainIdentity" method="post" path="/v3/identity/{identityId}/cross-domain" -->
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

    res, err := s.Identity.V3CrossDomainIdentity(ctx, "<id>", &components.V3CrossDomainIdentityRequest{
        ClientRequestID: provesdkservergo.Pointer("71010d88-d0e7-4a24-9297-d1be6fefde81"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3CrossDomainIdentityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         | Example                                                                                             |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |                                                                                                     |
| `identityID`                                                                                        | *string*                                                                                            | :heavy_check_mark:                                                                                  | A Prove-generated unique ID for a specific identity.                                                |                                                                                                     |
| `v3CrossDomainIdentityRequest`                                                                      | [*components.V3CrossDomainIdentityRequest](../../models/components/v3crossdomainidentityrequest.md) | :heavy_minus_sign:                                                                                  | N/A                                                                                                 | {<br/>"clientRequestId": "71010d88-d0e7-4a24-9297-d1be6fefde81"<br/>}                               |
| `opts`                                                                                              | [][operations.Option](../../models/operations/option.md)                                            | :heavy_minus_sign:                                                                                  | The options for this request.                                                                       |                                                                                                     |

### Response

**[*operations.V3CrossDomainIdentityResponse](../../models/operations/v3crossdomainidentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3DeactivateIdentity

Stops webhook notifications without disenrolling the identity.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3DeactivateIdentity" method="post" path="/v3/identity/{identityId}/deactivate" -->
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

    res, err := s.Identity.V3DeactivateIdentity(ctx, "<id>", &components.V3IdentityDeactivateRequest{
        ClientRequestID: provesdkservergo.Pointer("71010d88-d0e7-4a24-9297-d1be6fefde81"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3DeactivateIdentityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       | Example                                                                                           |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |                                                                                                   |
| `identityID`                                                                                      | *string*                                                                                          | :heavy_check_mark:                                                                                | A Prove-generated unique ID for a specific identity.                                              |                                                                                                   |
| `v3IdentityDeactivateRequest`                                                                     | [*components.V3IdentityDeactivateRequest](../../models/components/v3identitydeactivaterequest.md) | :heavy_minus_sign:                                                                                | N/A                                                                                               | {<br/>"clientRequestId": "71010d88-d0e7-4a24-9297-d1be6fefde81"<br/>}                             |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |                                                                                                   |

### Response

**[*operations.V3DeactivateIdentityResponse](../../models/operations/v3deactivateidentityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3GetIdentitiesByPhoneNumber

Return list of all identities you have enrolled that are associated with this phone number.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3GetIdentitiesByPhoneNumber" method="get" path="/v3/identity/{mobileNumber}/lookup" -->
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

    res, err := s.Identity.V3GetIdentitiesByPhoneNumber(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.V3GetIdentitiesByPhoneNumberResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                    | Type                                                                                                                                                                                                                                                                                         | Required                                                                                                                                                                                                                                                                                     | Description                                                                                                                                                                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The context to use for the request.                                                                                                                                                                                                                                                          |
| `mobileNumber`                                                                                                                                                                                                                                                                               | *string*                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                           | The phone number to find identities for. US phone numbers can be passed in with or without a leading +1. Acceptable characters are: alphanumeric with symbols '+'.                                                                                                                           |
| `clientRequestID`                                                                                                                                                                                                                                                                            | **string*                                                                                                                                                                                                                                                                                    | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | A client-generated unique ID for a specific session. This can be used to identify specific requests. The format of this ID is defined by the client - Prove recommends using a GUID, but any format can be accepted. Do not include Personally Identifiable Information (PII) in this field. |
| `opts`                                                                                                                                                                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                           | The options for this request.                                                                                                                                                                                                                                                                |

### Response

**[*operations.V3GetIdentitiesByPhoneNumberResponse](../../models/operations/v3getidentitiesbyphonenumberresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |