# V3
(*V3*)

## Overview

### Available Operations

* [V3TokenRequest](#v3tokenrequest) - Request OAuth token.
* [V3ChallengeRequest](#v3challengerequest) - Submit challenge.
* [V3CompleteRequest](#v3completerequest) - Complete flow.
* [V3MFARequest](#v3mfarequest) - Initiate possession check.
* [V3MFABindRequest](#v3mfabindrequest) - Check status of MFA session.
* [V3MFAStatusRequest](#v3mfastatusrequest) - Check status of MFA session.
* [V3StartRequest](#v3startrequest) - Start flow.
* [V3ValidateRequest](#v3validaterequest) - Validate phone number.
* [V3VerifyRequest](#v3verifyrequest) - Initiate verified users session.
* [V3VerifyStatusRequest](#v3verifystatusrequest) - Perform checks for verified users session.

## V3TokenRequest

Send this request to request the OAuth token.

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

Send this request to submit challenge information. Either a DOB or last 4 of SSN needs to be submitted if neither was submitted to the /start endpoint (challenge fields submitted to this endpoint will overwrite the /start endpoint fields submitted). It will return a correlation ID, user information, and the next step to call in the flow. This capability is only available in Pre-Fill®, it's not available in Prove Identity®. You'll notice that when using Prove Identity®, if /validate is successful, it will then return `v3-complete` as one of the keys in the `Next` field map instead of `v3-challenge`.

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

Send this request to verify the user and complete the flow. It will return a correlation ID, user information, and the next step to call in the flow. There is a validation check that requires at least first + last name or SSN passed in, else an HTTP 400 is returned. Additionally, specific to the Pre-Fill® or Prove Identity® with KYC use case, you need to pass in first name, last name, DOB and SSN (or address) to ensure you receive back the KYC elements and correct CIP values.

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
                components.V3CompleteAddressEntryRequest{
                    Address: provesdkservergo.String("4861 Jay Junction"),
                    City: provesdkservergo.String("Boston"),
                    ExtendedAddress: provesdkservergo.String("Apt 78"),
                    PostalCode: provesdkservergo.String("02208"),
                    Region: provesdkservergo.String("MS"),
                },
            },
            Dob: provesdkservergo.String("1981-01"),
            EmailAddresses: []string{
                "jdoe@example.com",
                "dsmith@example.com",
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

## V3MFARequest

Send this request to initiate a possession check. It will return a correlation ID
and authToken for the client SDK.

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

    res, err := s.V3.V3MFARequest(ctx, &components.V3MFARequest{
        ClientCustomerID: provesdkservergo.String("e0f78bc2-f748-4eda-9d29-d756844507fc"),
        ClientRequestID: provesdkservergo.String("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        EmailAddress: provesdkservergo.String("user@example.com"),
        FinalTargetURL: provesdkservergo.String("https://www.example.com/landing-page"),
        IPAddress: provesdkservergo.String("192.168.1.1"),
        PhoneNumber: provesdkservergo.String("2001004011"),
        PossessionType: "mobile",
        SmsMessage: provesdkservergo.String("#### is your verification code"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3MFAResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [components.V3MFARequest](../../models/components/v3mfarequest.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.V3MFARequestResponse](../../models/operations/v3mfarequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3MFABindRequest

Send this request to bind Prove Key to a phone nuymber of an MFA session and get the possession result.

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

    res, err := s.V3.V3MFABindRequest(ctx, &components.V3MFABindRequest{
        ClientRequestID: provesdkservergo.String("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        CorrelationID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
        PhoneNumber: provesdkservergo.String("2001004011"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3MFABindResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.V3MFABindRequest](../../models/components/v3mfabindrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.V3MFABindRequestResponse](../../models/operations/v3mfabindrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3MFAStatusRequest

Send this request to check the status of an MFA session and get the possession result.

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

    res, err := s.V3.V3MFAStatusRequest(ctx, &components.V3MFAStatusRequest{
        ClientRequestID: provesdkservergo.String("71010d88-d0e7-4a24-9297-d1be6fefde81"),
        CorrelationID: provesdkservergo.String("713189b8-5555-4b08-83ba-75d08780aebd"),
        PhoneNumber: provesdkservergo.String("2001004011"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.V3MFAStatusResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.V3MFAStatusRequest](../../models/components/v3mfastatusrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.V3MFAStatusRequestResponse](../../models/operations/v3mfastatusrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error401 | 401                | application/json   |
| sdkerrors.Error403 | 403                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## V3StartRequest

Send this request to start a Prove flow. It will return a correlation ID and an authToken for the client SDK.

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

## V3ValidateRequest

Send this request to check the phone number entered/discovered earlier in the flow is validated. It will return a correlation ID and the next step.

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

Send this request to initiate a Verified Users session. It will return a correlation ID, authToken for the client SDK, and the results of the possession and verify checks (usually pending from this API).

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

Send this request to perform the necessary checks for a Verified Users session. It will return the results of the possession and verify checks, as well as the overall success.

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