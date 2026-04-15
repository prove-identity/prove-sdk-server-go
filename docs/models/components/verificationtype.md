# VerificationType

The verification method based on the use case and authorization level. Current allowed values: "verifiedUser", "accountOpening", "humanAssurance", "prefill", "prefillForBiz", "identityResolution".

## Example Usage

```go
import (
	"github.com/prove-identity/prove-sdk-server-go/models/components"
)

value := components.VerificationTypeHumanAssurance
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `VerificationTypeHumanAssurance`     | humanAssurance                       |
| `VerificationTypeVerifiedUser`       | verifiedUser                         |
| `VerificationTypeAccountOpening`     | accountOpening                       |
| `VerificationTypePrefill`            | prefill                              |
| `VerificationTypePrefillForBiz`      | prefillForBiz                        |
| `VerificationTypeIdentityResolution` | identityResolution                   |