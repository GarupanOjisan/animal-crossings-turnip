# animal-crossings-turnip

## Buyer API

### GET /buyer

get index of buyer models

#### Request params
|param|type|required|description|
|-|-|-|-|
|page|int|-|a page index (0 >)|

#### Response

for example
```json
[
    {
        "ID": 71,
        "Price": 333,
        "Password": "test1",
        "LimitNumVisitor": -1,
        "CreatedAt": "2020-04-19T15:41:51+09:00",
        "ExpiredAt": null
    },
    {
        "ID": 75,
        "Price": 333,
        "Password": "test1",
        "LimitNumVisitor": -1,
        "CreatedAt": "2020-04-19T15:41:51+09:00",
        "ExpiredAt": "2020-04-20T14:41:51+09:00"
    }
]
```

### POST /buyer

Create a new buyer model.

#### Request params
|param|type|required|description|
|-|-|-|-|
|price|int|o|the sell price of turnips on your island.|
|password|string|o|the password to enter your island.|
|limit|int|-|the limit of friends who can come your island.|
|expired_at|datetime|-|datetime string when your island close.(RFC3339 formatted)|

#### Response

It returns a created Buyer model in json format.
For example,

```json
{
    "ID": 76,
    "Price": 90,
    "Password": "test1",
    "LimitNumVisitor": 0,
    "CreatedAt": "2020-04-19T16:09:00.386163+09:00",
    "ExpiredAt": "2020-04-19T19:00:00+09:00"
}
```


## Seller API

### GET /seller

get index of seller models

#### Request params
|param|type|required|description|
|-|-|-|-|
|page|int|-|a page index (0 >)|

#### Response

for example
```json
[
    {
        "ID": 71,
        "Price": 333,
        "Password": "test1",
        "LimitNumVisitor": -1,
        "CreatedAt": "2020-04-19T15:41:51+09:00",
        "ExpiredAt": null
    },
    {
        "ID": 75,
        "Price": 333,
        "Password": "test1",
        "LimitNumVisitor": -1,
        "CreatedAt": "2020-04-19T15:41:51+09:00",
        "ExpiredAt": "2020-04-20T14:41:51+09:00"
    }
]
```

### POST /seller

Create a new seller model.

#### Request params
|param|type|required|description|
|-|-|-|-|
|price|int|o|the sell price of turnips on your island.|
|password|string|o|the password to enter your island.|
|limit|int|-|the limit of friends who can come your island.|
|expired_at|datetime|-|datetime string when your island close.(RFC3339 formatted)|

#### Response

It returns a created Buyer model in json format.
For example,

```json
{
    "ID": 76,
    "Price": 90,
    "Password": "test1",
    "LimitNumVisitor": 0,
    "CreatedAt": "2020-04-19T16:09:00.386163+09:00",
    "ExpiredAt": "2020-04-19T19:00:00+09:00"
}
```