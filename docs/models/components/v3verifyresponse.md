# V3VerifyResponse


## Fields

| Field                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                | Example                                                                                                                                                                                                                    |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `AuthToken`                                                                                                                                                                                                                | **string*                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                         | AuthToken is a bearer token for use by the Prove Client SDK.                                                                                                                                                               | eyJhbGciOi...                                                                                                                                                                                                              |
| `CorrelationID`                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                         | Correlation ID is the unique ID that Prove generates for the flow. To continue the flow, the field will also be used for each of the subsequent API calls in the same flow - it cannot be reused outside of a single flow. | 713189b8-5555-4b08-83ba-75d08780aebd                                                                                                                                                                                       |
| `PossessionResult`                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                        |                                                                                                                                                                                                                            |
| `Success`                                                                                                                                                                                                                  | *string*                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                        |                                                                                                                                                                                                                            |
| `VerifyResult`                                                                                                                                                                                                             | *string*                                                                                                                                                                                                                   | :heavy_check_mark:                                                                                                                                                                                                         | N/A                                                                                                                                                                                                                        |                                                                                                                                                                                                                            |