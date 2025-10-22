# V3VerifyBatchRequest


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ClientRequestID`                                                               | **string*                                                                       | :heavy_minus_sign:                                                              | N/A                                                                             |
| `Items`                                                                         | [][components.VerifyItem](../../models/components/verifyitem.md)                | :heavy_check_mark:                                                              | Batch of verify requests to process. The array length cannot exceed 1000 items. |