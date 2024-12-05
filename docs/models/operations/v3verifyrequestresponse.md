# V3VerifyRequestResponse


## Fields

| Field                                                                                                                                                                                   | Type                                                                                                                                                                                    | Required                                                                                                                                                                                | Description                                                                                                                                                                             | Example                                                                                                                                                                                 |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `HTTPMeta`                                                                                                                                                                              | [components.HTTPMetadata](../../models/components/httpmetadata.md)                                                                                                                      | :heavy_check_mark:                                                                                                                                                                      | N/A                                                                                                                                                                                     |                                                                                                                                                                                         |
| `V3VerifyResponse`                                                                                                                                                                      | [*components.V3VerifyResponse](../../models/components/v3verifyresponse.md)                                                                                                             | :heavy_minus_sign:                                                                                                                                                                      | Successful request.                                                                                                                                                                     | {<br/>"success": "success",<br/>"authToken": "eyJhbGciOi...",<br/>"possessionResult": "possessionResult",<br/>"verifyResult": "verifyResult",<br/>"correlationId": "713189b8-5555-4b08-83ba-75d08780aebd"<br/>} |