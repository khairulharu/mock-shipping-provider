# Mock Shipping Provider

This repository mock a shipping gateway provider.

## Endpoints

### Estimate

Look up price and estimated hour arrival for an item (with a list of providers)
based on coordinates, dimension, and weight.

```http request
POST /estimate
Content-Type: application/json
Accept: application/json

{
  "sender": {
    "latitude": 1.234,
    "longitude: 1.234
  },
  "recipient": {
    "latitude": 1.234,
    "longitude": 1.234
  },
  "dimension": {
    "height": 1.234,
    "width": 1.234,
    "depth": 1.234
  },
  "weight": 1.234
}
```

Weight, height, and depth are in centimeters. Weight is in kilograms.

Response body:

```json
{
    "estimation": [
        {
            "provider": "JNE",
            "estimated_price": 1234,
            "estimated_hour_arrival": 1234
        }
    ]
}
```

### Order

Creates a new order. It recalculates the price and the hour of arrival. 
It returns a set of reference number and air waybill number that client 
should keep safe.

Once created, it will issue a webhook calls whenever there is a status change
for the specified shipping order to the specified target URL that should be 
configured before.

```http request
POST /order
Content-Type: application/json
Accept: application/json

{
  "provider": "JNE",
  "sender": {
    "name": "John Doe",
    "phone_number": "628123456",
    "address": "string",
    "city": "string",
    "state": "string",
    "country": "string",
    "postal_code": "string",
    "coordinate": {
      "latitude": 1.234,
      "longitude: 1.234
    }
  },
  "recipient": {
    "name": "Anne Doe",
    "phone_number": "628123456",
    "address": "string",
    "city": "string",
    "state": "string",
    "country": "string",
    "postal_code": "string",
    "coordinate": {
      "latitude": 1.234,
      "longitude: 1.234
    }
  },
  "dimension": {
    "height": 1.234,
    "width": 1.234,
    "depth": 1.234
  },
  "weight": 1.234,
  "item_description": "string",
  "item_category": "string",
  "fragile": true
}
```

Response body:

```json
{
    "status_code": 1,
    "status_description": "ORDER_PLACED",
    "reference_number": "string",
    "air_waybill": "string",
    "price": 1234,
    "estimated_hour_arrival": 1234
}
```

### Status History

Returns an array of status history log of the specified reference number and air waybill.

```http request
GET /status-history?reference_number=string&air_waybill=string
Accept: application/json
```

Response body:

```json
{
    "reference_number": "string",
    "air_waybill": "string",
    "history": [
        {
            "status_code": 1,
            "status_description": "ORDER_PLACED",
            "timestamp": "2023-01-01T00:00:00Z",
            "note": "string"
        }
    ]
}
```

## License

```
   Copyright 2023 TokoBapak

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
```

See [LICENSE](./LICENSE)
