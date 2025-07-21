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
* [V3VerifyStatusRequest](#v3verifystatusrequest) - Perform Checks for Verified Users Session

## V3TokenRequest

This endpoint allows you to request an OAuth token.

### Example Usage

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

    res, err := s.V3.V3ChallengeRequest(ctx, &components.V3ChallengeRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
        Dob: provesdkservergo.String("1981-01"),
        Ssn: provesdkservergo.String("0596"),
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

    res, err := s.V3.V3CompleteRequest(ctx, &components.V3CompleteRequest{
        CorrelationID: "713189b8-5555-4b08-83ba-75d08780aebd",
        Individual: components.V3CompleteIndividualRequest{
            Addresses: []components.V3CompleteAddressEntryRequest{
                components.V3CompleteAddressEntryRequest{
                    Address: provesdkservergo.String("39 South Trail"),
                    City: provesdkservergo.String("San Antonio"),
                    ExtendedAddress: provesdkservergo.String("Apt 23"),
                    PostalCode: provesdkservergo.String("78285"),
                    Region: provesdkservergo.String("TX"),
                },
            },
            Dob: provesdkservergo.String("1981-01"),
            EmailAddresses: []string{
                "jdoe@example.com",
            },
            FirstName: provesdkservergo.String("Tod"),
            LastName: provesdkservergo.String("Weedall"),
            Ssn: provesdkservergo.String("265228370"),
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

    res, err := s.V3.V3StartRequest(ctx, &components.V3StartRequest{
        AllowOTPRetry: provesdkservergo.Bool(true),
        Dob: provesdkservergo.String("1981-01"),
        EmailAddress: provesdkservergo.String("mpinsonm@dyndns.org"),
        FinalTargetURL: provesdkservergo.String("https://www.example.com/landing-page"),
        FlowType: "mobile",
        IPAddress: provesdkservergo.String("10.0.0.1"),
        PhoneNumber: provesdkservergo.String("2001001695"),
        SmsMessage: provesdkservergo.String("#### is your temporary code to continue your application. Caution: for your security, don't share this code with anyone."),
        Ssn: provesdkservergo.String("0596"),
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

    res, err := s.V3.V3UnifyRequest(ctx, &components.V3UnifyRequest{
        AllowOTPRetry: provesdkservergo.Bool(true),
        ClientCustomerID: provesdkservergo.String("e0f78bc2-f748-4eda-9d29-d756844507fc"),
        ClientRequestID: provesdkservergo.String("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        FinalTargetURL: provesdkservergo.String("https://www.example.com/landing-page"),
        PhoneNumber: provesdkservergo.String("2001004011"),
        PossessionType: "mobile",
        Rebind: provesdkservergo.Bool(true),
        SmsMessage: provesdkservergo.String("#### is your verification code."),
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

    res, err := s.V3.V3UnifyBindRequest(ctx, &components.V3UnifyBindRequest{
        ClientRequestID: provesdkservergo.String("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        CorrelationID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
        PhoneNumber: provesdkservergo.String("2001004011"),
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

    res, err := s.V3.V3UnifyStatusRequest(ctx, &components.V3UnifyStatusRequest{
        ClientRequestID: provesdkservergo.String("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        CorrelationID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
        PhoneNumber: provesdkservergo.String("2001004011"),
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

This endpoint allows you to initiate a Verified Users session.

### Example Usage

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

    res, err := s.V3.V3VerifyRequest(ctx, &components.V3VerifyRequest{
        AllowOTPRetry: provesdkservergo.Bool(true),
        ClientCustomerID: provesdkservergo.String("e0f78bc2-f748-4eda-9d29-d756844507fc"),
        ClientRequestID: provesdkservergo.String("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        EmailAddress: provesdkservergo.String("sbutrimovichb@who.int"),
        FinalTargetURL: provesdkservergo.String("https://www.example.com/landing-page"),
        FirstName: "Sheilakathryn",
        LastName: "Butrimovich",
        PhoneNumber: "2001004011",
        PossessionType: "mobile",
        SmsMessage: provesdkservergo.String("#### is your temporary code to continue your application. Caution: for your security, don't share this code with anyone."),
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

## V3VerifyStatusRequest

This endpoint allows you to perform the necessary checks for a Verified Users session.

### Example Usage

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

    res, err := s.V3.V3VerifyStatusRequest(ctx, &components.V3VerifyStatusRequest{
        ClientRequestID: provesdkservergo.String("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        CorrelationID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3VerifyStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [components.V3VerifyStatusRequest](../../models/components/v3verifystatusrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.V3VerifyStatusRequestResponse](../../models/operations/v3verifystatusrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |