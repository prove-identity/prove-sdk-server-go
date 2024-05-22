# V3ValidateResponse


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ChallengeMissing`                                                                     | *bool*                                                                                 | :heavy_check_mark:                                                                     | Challenge missing returns true if a DOB or SSN needs to be passed in on the next step. | true                                                                                   |
| `Next`                                                                                 | map[string]*string*                                                                    | :heavy_check_mark:                                                                     | Next contains the next set of allowed calls in the same flow.                          | {<br/>"v3-challenge": "/v3/challenge"<br/>}                                            |
| `Success`                                                                              | *bool*                                                                                 | :heavy_check_mark:                                                                     | Success returns true if the phone number was validated.                                | true                                                                                   |