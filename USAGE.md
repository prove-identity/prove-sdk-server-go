<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	provesdkservergo "github.com/prove-identity/prove-sdk-server-go"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"log"
)

func main() {
	s := provesdkservergo.New(
		provesdkservergo.WithSecurity(components.Security{
			ClientID:     provesdkservergo.String("<YOUR_CLIENT_ID_HERE>"),
			ClientSecret: provesdkservergo.String("<YOUR_CLIENT_SECRET_HERE>"),
		}),
	)
	var request *components.V3StartRequest = &components.V3StartRequest{
		Dob:            provesdkservergo.String("2024-05-02T00:00:00Z"),
		EmailAddress:   provesdkservergo.String("jdoe@example.com"),
		FinalTargetURL: provesdkservergo.String("https://www.example.com/landing-page"),
		FlowType:       "mobile",
		IPAddress:      provesdkservergo.String("10.0.0.1"),
		PhoneNumber:    provesdkservergo.String("2001001686"),
		Ssn:            provesdkservergo.String("1234"),
	}
	ctx := context.Background()
	res, err := s.V3.V3StartRequest(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.V3StartResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->