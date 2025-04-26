# Error403


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `Code`                                                                   | **int64*                                                                 | :heavy_minus_sign:                                                       | An internal error code that identifies the specific authorization issue. | 8003                                                                     |
| `Message`                                                                | *string*                                                                 | :heavy_check_mark:                                                       | An error message describing why access is forbidden.                     | Access forbidden: insufficient permissions                               |