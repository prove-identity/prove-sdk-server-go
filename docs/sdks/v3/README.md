# V3
(*V3*)

## Overview

### Available Operations

* [V3TokenRequest](#v3tokenrequest) - Request OAuth Token
* [V3ChallengeRequest](#v3challengerequest) - Submit Challenge
* [V3CompleteRequest](#v3completerequest) - Complete Flow
* [V3StartRequest](#v3startrequest) - Start Flow
* [V3UnifyRequest](#v3unifyrequest) - Initiate Possession Check
* [V3UnifyBindRequest](#v3unifybindrequest) - Bind Prove Key
* [V3UnifyStatusRequest](#v3unifystatusrequest) - Check Status
* [V3ValidateRequest](#v3validaterequest) - Validate Phone Number
* [V3VerifyRequest](#v3verifyrequest) - Initiate Verified Users Session
* [V3VerifyBatchRequest](#v3verifybatchrequest) - Batch Verify Users

## V3TokenRequest

This endpoint allows you to request an OAuth token.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3TokenRequest" method="post" path="/token" -->
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

    s := provesdkservergo.New()

    res, err := s.V3.V3TokenRequest(ctx, &components.V3TokenRequest{
        ClientID: "customer_id",
        ClientSecret: "secret",
        GrantType: "client_credentials",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3TokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.V3TokenRequest](../../models/components/v3tokenrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.V3TokenRequestResponse](../../models/operations/v3tokenrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3ChallengeRequest

This endpoint allows you to submit challenge information.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3ChallengeRequest" method="post" path="/v3/challenge" -->
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

    res, err := s.V3.V3ChallengeRequest(ctx, &components.V3ChallengeRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
        Dob: provesdkservergo.Pointer("1981-01"),
        Ssn: provesdkservergo.Pointer("0596"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ChallengeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.V3ChallengeRequest](../../models/components/v3challengerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.V3ChallengeRequestResponse](../../models/operations/v3challengerequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3CompleteRequest

This endpoint allows you to verify the user and complete the flow.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3CompleteRequest" method="post" path="/v3/complete" -->
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

    res, err := s.V3.V3CompleteRequest(ctx, &components.V3CompleteRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
        Individual: components.V3CompleteIndividualRequest{
            Addresses: []components.V3CompleteAddressEntryRequest{
                components.V3CompleteAddressEntryRequest{
                    Address: provesdkservergo.Pointer("39 South Trail"),
                    City: provesdkservergo.Pointer("San Antonio"),
                    ExtendedAddress: provesdkservergo.Pointer("Apt 23"),
                    PostalCode: provesdkservergo.Pointer("78285"),
                    Region: provesdkservergo.Pointer("TX"),
                },
            },
            Dob: provesdkservergo.Pointer("1981-01"),
            EmailAddresses: []string{
                "jdoe@example.com",
            },
            FirstName: provesdkservergo.Pointer("Tod"),
            LastName: provesdkservergo.Pointer("Weedall"),
            Ssn: provesdkservergo.Pointer("265228370"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3CompleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.V3CompleteRequest](../../models/components/v3completerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.V3CompleteRequestResponse](../../models/operations/v3completerequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3StartRequest

This endpoint allows you to start the solution flow.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3StartRequest" method="post" path="/v3/start" -->
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

    res, err := s.V3.V3StartRequest(ctx, &components.V3StartRequest{
        AllowOTPRetry: provesdkservergo.Pointer(true),
        Dob: provesdkservergo.Pointer("1981-01"),
        EmailAddress: provesdkservergo.Pointer("mpinsonm@dyndns.org"),
        FinalTargetURL: provesdkservergo.Pointer("https://www.example.com/landing-page"),
        FlowType: "mobile",
        IPAddress: provesdkservergo.Pointer("10.0.0.1"),
        PhoneNumber: provesdkservergo.Pointer("2001001695"),
        SmsMessage: provesdkservergo.Pointer("#### is your temporary code to continue your application. Caution: for your security, don't share this code with anyone."),
        Ssn: provesdkservergo.Pointer("0596"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3StartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.V3StartRequest](../../models/components/v3startrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.V3StartRequestResponse](../../models/operations/v3startrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3UnifyRequest

This endpoint allows you to initiate the possession check.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3UnifyRequest" method="post" path="/v3/unify" -->
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

    res, err := s.V3.V3UnifyRequest(ctx, &components.V3UnifyRequest{
        AllowOTPRetry: provesdkservergo.Pointer(true),
        CheckReputation: provesdkservergo.Pointer(true),
        ClientCustomerID: provesdkservergo.Pointer("e0f78bc2-f748-4eda-9d29-d756844507fc"),
        ClientHumanID: provesdkservergo.Pointer("7bfbb91d-9df8-4ec0-99a6-de05ecc23a9e"),
        ClientRequestID: "71010d88-d0e7-4a24-9297-d1be6fefde81",
        DeviceID: provesdkservergo.Pointer("bf9ea15d-7dfa-4bb4-a64c-6c26b53472fc"),
        EmailAddress: provesdkservergo.Pointer("sbutrimovichb@who.int"),
        FinalTargetURL: provesdkservergo.Pointer("https://www.example.com/landing-page"),
        IPAddress: provesdkservergo.Pointer("192.168.0.1"),
        PhoneNumber: provesdkservergo.Pointer("2001004011"),
        PossessionType: "mobile",
        ProveID: provesdkservergo.Pointer("a07b94ce-218c-461f-beda-d92480e40f61"),
        Rebind: provesdkservergo.Pointer(true),
        SmsMessage: provesdkservergo.Pointer("#### is your verification code."),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3UnifyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.V3UnifyRequest](../../models/components/v3unifyrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.V3UnifyRequestResponse](../../models/operations/v3unifyrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3UnifyBindRequest

This endpoint allows you to bind a Prove Key to a phone number of a Unify session and get the possession result.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3UnifyBindRequest" method="post" path="/v3/unify-bind" -->
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

    res, err := s.V3.V3UnifyBindRequest(ctx, &components.V3UnifyBindRequest{
        ClientRequestID: provesdkservergo.Pointer("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        CorrelationID: provesdkservergo.Pointer("713189b8-5555-4b08-83ba-75d08780aebd"),
        PhoneNumber: provesdkservergo.Pointer("2001004011"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3UnifyBindResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.V3UnifyBindRequest](../../models/components/v3unifybindrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.V3UnifyBindRequestResponse](../../models/operations/v3unifybindrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3UnifyStatusRequest

This endpoint allows you to check the status of a Unify session and get the possession result.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3UnifyStatusRequest" method="post" path="/v3/unify-status" -->
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

    res, err := s.V3.V3UnifyStatusRequest(ctx, &components.V3UnifyStatusRequest{
        ClientRequestID: provesdkservergo.Pointer("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        CorrelationID: provesdkservergo.Pointer("713189b8-5555-4b08-83ba-75d08780aebd"),
        PhoneNumber: provesdkservergo.Pointer("2001004011"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3UnifyStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.V3UnifyStatusRequest](../../models/components/v3unifystatusrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.V3UnifyStatusRequestResponse](../../models/operations/v3unifystatusrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3ValidateRequest

This endpoint allows you to check if the phone number entered/discovered earlier in the flow is validated.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3ValidateRequest" method="post" path="/v3/validate" -->
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

    res, err := s.V3.V3ValidateRequest(ctx, &components.V3ValidateRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3ValidateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.V3ValidateRequest](../../models/components/v3validaterequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.V3ValidateRequestResponse](../../models/operations/v3validaterequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3VerifyRequest

This endpoint allows you to verify a user depending on your particular use case.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3VerifyRequest" method="post" path="/v3/verify" -->
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

    res, err := s.V3.V3VerifyRequest(ctx, &components.V3VerifyRequest{
        ClientCustomerID: provesdkservergo.Pointer("e0f78bc2-f748-4eda-9d29-d756844507fc"),
        ClientHumanID: provesdkservergo.Pointer("aad25769-23bb-458c-80db-50296a82c91b"),
        ClientRequestID: provesdkservergo.Pointer("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        EmailAddress: provesdkservergo.Pointer("ecoldman1h@storify.com"),
        FirstName: provesdkservergo.Pointer("Elena"),
        IPAddress: provesdkservergo.Pointer("192.168.1.1"),
        LastName: provesdkservergo.Pointer("Coldman"),
        PhoneNumber: "2001004053",
        UserAgent: provesdkservergo.Pointer("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0"),
        VerificationType: components.VerificationTypeVerifiedUser,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3VerifyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.V3VerifyRequest](../../models/components/v3verifyrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.V3VerifyRequestResponse](../../models/operations/v3verifyrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3VerifyBatchRequest

This endpoint allows you to batch verify and enroll users.

### Example Usage

<!-- UsageSnippet language="go" operationID="V3VerifyBatchRequest" method="post" path="/v3/verify/batch" -->
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

    res, err := s.V3.V3VerifyBatchRequest(ctx, &components.V3VerifyBatchRequest{
        ClientRequestID: provesdkservergo.Pointer("3d1215f7-ec3f-4fd2-9894-7b46f00e31a6"),
        Items: []components.VerifyBatchRequestItem{
            components.VerifyBatchRequestItem{
                ClientCustomerID: provesdkservergo.Pointer("e0f78bc2-f748-4eda-9d29-d756844507fc"),
                ClientHumanID: provesdkservergo.Pointer("clientHumanId"),
                EmailAddress: provesdkservergo.Pointer("ecoldman1h@storify.com"),
                FirstName: "Elena",
                IPAddress: provesdkservergo.Pointer("192.168.1.1"),
                LastName: "Coldman",
                PhoneNumber: "2001004053",
                ProveID: provesdkservergo.Pointer("e0f78bc2-f748-4eda-9d29-d756844507fc"),
                UserAgent: provesdkservergo.Pointer("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0"),
                VerificationType: "verifiedUser",
            },
            components.VerifyBatchRequestItem{
                ClientCustomerID: provesdkservergo.Pointer("e0f78bc2-f748-4eda-9d29-d756844507fc"),
                ClientHumanID: provesdkservergo.Pointer("clientHumanId"),
                EmailAddress: provesdkservergo.Pointer("ecoldman1h@storify.com"),
                FirstName: "Elena",
                IPAddress: provesdkservergo.Pointer("192.168.1.1"),
                LastName: "Coldman",
                PhoneNumber: "2001004053",
                ProveID: provesdkservergo.Pointer("e0f78bc2-f748-4eda-9d29-d756844507fc"),
                UserAgent: provesdkservergo.Pointer("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0"),
                VerificationType: "verifiedUser",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3VerifyBatchResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.V3VerifyBatchRequest](../../models/components/v3verifybatchrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.V3VerifyBatchRequestResponse](../../models/operations/v3verifybatchrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |