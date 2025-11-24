# Auth
(*Auth*)

## Overview

### Available Operations

* [AuthContinueRequest](#authcontinuerequest) - AuthContinue /v1/server/auth/continue
* [AuthFinishRequest](#authfinishrequest) - AuthFinish /v1/server/auth/finish
* [AuthStartRequest](#authstartrequest) - AuthStart /v1/server/auth/start

## AuthContinueRequest

If the initial Prove Auth authenticators fail, the Force Bind authenticator can be used which permits using another
authentication method to verify a mobile number (like Prove Instant Link™). Once the mobile number is verified, send
this AuthContinue request.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/auth/continue

### Example Usage

<!-- UsageSnippet language="go" operationID="AuthContinueRequest" method="post" path="/v1/server/auth/continue" -->
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

    res, err := s.Auth.AuthContinueRequest(ctx, &components.AuthContinueRequest{
        AuthID: "713189b8-5555-4b08-83ba-75d08780aebd",
        RequestID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
        Subjects: components.AuthContinueRequestSubjects{
            Mobile: &components.AuthContinueRequestSubjectMobile{
                Claim: &components.AuthContinueRequestSubjectMobileClaim{
                    MobileNumber: "12065550100",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthContinueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.AuthContinueRequest](../../models/components/authcontinuerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.AuthContinueRequestResponse](../../models/operations/authcontinuerequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AuthFinishRequest

To complete the auth flow and get the authentication result, send this AuthFinish request.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/auth/finish

### Example Usage

<!-- UsageSnippet language="go" operationID="AuthFinishRequest" method="post" path="/v1/server/auth/finish" -->
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

    res, err := s.Auth.AuthFinishRequest(ctx, &components.AuthFinishRequest{
        AuthID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthFinishResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.AuthFinishRequest](../../models/components/authfinishrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AuthFinishRequestResponse](../../models/operations/authfinishrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## AuthStartRequest

To start an auth flow, send this AuthStart request to control how Prove Auth will authenticate the end users and
their devices. The expected authenticators should be included in the body parameters. There are many supported
scenarios to use Prove Auth so each of the request types are explained in the "Server Integration Guide".

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/auth/start

### Example Usage

<!-- UsageSnippet language="go" operationID="AuthStartRequest" method="post" path="/v1/server/auth/start" -->
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

    res, err := s.Auth.AuthStartRequest(ctx, &components.AuthStartRequest{
        CallbackURL: provesdkservergo.Pointer("https://example.com/webhook"),
        Delivery: &components.AuthStartRequestDelivery{
            Push: &components.AuthStartRequestDeliveryPush{
                Notification: components.AuthStartRequestDeliveryPushNotification{
                    Body: "Do you want to allow login from this device?",
                    ConfirmBtn: provesdkservergo.Pointer("Confirm"),
                    DenyBtn: provesdkservergo.Pointer("Deny"),
                    OriginatingDevice: provesdkservergo.Pointer("iPhone12"),
                    OriginatingIP: provesdkservergo.Pointer("198.51.100.10"),
                    Title: "Confirm Login",
                },
            },
            Scan: &components.AuthStartRequestDeliveryScan{
                Notification: components.AuthStartRequestDeliveryScanNotification{
                    Body: "Please confirm you are trying to sign in...",
                    ConfirmBtn: provesdkservergo.Pointer("Confirm"),
                    DenyBtn: provesdkservergo.Pointer("Deny"),
                    OriginatingDevice: provesdkservergo.Pointer("Google Chrome on Windows"),
                    OriginatingIP: provesdkservergo.Pointer("198.51.100.10"),
                    Title: "Confirm Sign In",
                },
            },
        },
        RequestID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
        Subjects: components.AuthStartRequestSubjects{
            Card: &components.AuthStartRequestSubjectCard{
                Authenticators: components.AuthStartRequestSubjectCardAuthenticators{
                    AirKey: &components.AuthStartRequestSubjectCardAuthenticatorAirKey{
                        Claim: &components.AuthStartRequestSubjectCardAuthenticatorAirKeyClaim{
                            Puids: []string{
                                "puids",
                                "puids",
                            },
                        },
                        TestMode: provesdkservergo.Pointer("testMode"),
                    },
                },
            },
            Device: &components.AuthStartRequestSubjectDevice{
                Authenticators: &components.AuthStartRequestSubjectDeviceAuthenticators{
                    Passive: &components.AuthStartRequestSubjectDeviceAuthenticatorPassive{
                        AllowPasscodeFallback: provesdkservergo.Pointer(true),
                        UserVerificationLevel: provesdkservergo.Pointer("userVerificationLevel"),
                    },
                },
                Claim: &components.AuthStartRequestSubjectDeviceClaim{
                    DeviceID: provesdkservergo.Pointer("deviceId"),
                },
            },
            Mobile: &components.AuthStartRequestSubjectMobile{
                Authenticators: &components.AuthStartRequestSubjectMobileAuthenticators{
                    Instant: &components.AuthStartRequestSubjectMobileAuthenticatorInstant{
                        Consent: &components.AuthStartRequestInstantAuthConsent{
                            CollectedTimestamp: provesdkservergo.Pointer("2022-02-17T00:00:00Z"),
                            Description: provesdkservergo.Pointer("Customer Application Name"),
                            Status: "optedIn",
                            TransactionID: provesdkservergo.Pointer("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
                        },
                        TestMode: provesdkservergo.Pointer("testMode"),
                    },
                    Instantlink: &components.AuthStartRequestSubjectMobileAuthenticatorInstantLink{
                        AllowMobileNumberRePrompt: provesdkservergo.Pointer(true),
                        FinalTargetURL: "https://example.com/finishpage",
                        MessageText: "Please click the link to authenticate your mobile number: ####",
                        SourceIP: provesdkservergo.Pointer("123..."),
                        SubscriptionCustomerID: provesdkservergo.Pointer("123"),
                        TestMode: provesdkservergo.Pointer("testMode"),
                    },
                    Otp: &components.AuthStartRequestSubjectMobileAuthenticatorOTP{
                        AllowMobileNumberRePrompt: provesdkservergo.Pointer(true),
                        AllowOtpRetry: provesdkservergo.Pointer(true),
                        MessageText: "Your pin is: ####",
                        TestMode: provesdkservergo.Pointer("testMode"),
                    },
                    Passive: &components.AuthStartRequestSubjectMobileAuthenticatorPassive{
                        CacheResult: provesdkservergo.Pointer(true),
                        LocalDomain: provesdkservergo.Pointer(true),
                        MaxAge: provesdkservergo.Pointer[int64](7776000),
                        TestMode: provesdkservergo.Pointer("testMode"),
                    },
                    Universal: &components.AuthStartRequestSubjectMobileAuthenticatorUniversal{
                        FinalTargetURL: "https://example.com/finishpage",
                        TestMode: provesdkservergo.Pointer("testMode"),
                    },
                },
                Claim: &components.AuthStartRequestSubjectMobileClaim{
                    MobileNumber: provesdkservergo.Pointer("12065550100"),
                },
            },
            User: &components.AuthStartRequestSubjectUser{
                Authenticators: &components.AuthStartRequestSubjectUserAuthenticators{
                    Docv: &components.Docv{},
                    Passive: &components.AuthStartRequestSubjectUserAuthenticatorPassive{
                        UserVerificationLevel: provesdkservergo.Pointer("userVerificationLevel"),
                    },
                    Ppb: &components.Ppb{},
                    Present: &components.Present{},
                },
                Claim: &components.AuthStartRequestSubjectUserClaim{
                    Provided: provesdkservergo.Pointer(true),
                    UserID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AuthStartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.AuthStartRequest](../../models/components/authstartrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../models/operations/option.md)                   | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.AuthStartRequestResponse](../../models/operations/authstartrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error400 | 400                | application/json   |
| sdkerrors.Error    | 500                | application/json   |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |