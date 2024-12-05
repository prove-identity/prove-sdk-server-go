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

	ctx := context.Background()
	res, err := s.V3.V3StartRequest(ctx, &components.V3StartRequest{
		Dob:            provesdkservergo.String("1981-01"),
		EmailAddress:   provesdkservergo.String("mpinsonm@dyndns.org"),
		FinalTargetURL: provesdkservergo.String("https://www.example.com/landing-page"),
		FlowType:       "mobile",
		IPAddress:      provesdkservergo.String("10.0.0.1"),
		PhoneNumber:    provesdkservergo.String("2001001695"),
		SmsMessage:     provesdkservergo.String("#### is your temporary code to continue your application. Caution: for your security, don't share this code with anyone."),
		Ssn:            provesdkservergo.String("0596"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.V3StartResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->