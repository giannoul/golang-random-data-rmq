# golang-random-data-rmq
Combine 3 services using RabbitMQ

This is a showcase of using 2 data sources dith distinct services:
* https://random-data-api.com/api/v2/users
* https://random-data-api.com/api/v2/credit_cards
in order to get data and combine them in using a 3rd service

```
|------------|
|    user    |----------|
|------------|          |
                        |      |----------|     |-----------|
                        |----->| RabbitMQ |---->| processor |
                        |      |----------|     |-----------|
|-------------|         |
| credit card |---------|
|-------------|
```

The end result is something like this:
```
processor     | {
processor     |         "User": {
processor     |                 "id": 9783,
processor     |                 "uid": "417ed33e-fa3e-4fb9-a6df-e1a2fdbc34f2",
processor     |                 "password": "0Uk6fy4sv2",
processor     |                 "first_name": "Roberto",
processor     |                 "last_name": "Bednar",
processor     |                 "username": "roberto.bednar",
processor     |                 "email": "roberto.bednar@email.com",
processor     |                 "avatar": "https://robohash.org/eaquenamquos.png?size=300x300\u0026set=set1",
processor     |                 "gender": "Genderqueer",
processor     |                 "phone_number": "+256 1-155-009-5393 x8248",
processor     |                 "social_insurance_number": "906298948",
processor     |                 "date_of_birth": "1958-08-08",
processor     |                 "employment": {
processor     |                         "title": "Hospitality Manager",
processor     |                         "key_skill": "Work under pressure"
processor     |                 },
processor     |                 "address": {
processor     |                         "city": "East Donettachester",
processor     |                         "street_name": "Emard Plaza",
processor     |                         "street_address": "395 Rosemary Landing",
processor     |                         "zip_code": "17833",
processor     |                         "state": "New Hampshire",
processor     |                         "country": "United States",
processor     |                         "coordinates": {
processor     |                                 "lat": 45.001007,
processor     |                                 "lng": 69.8333
processor     |                         }
processor     |                 },
processor     |                 "credit_card": {
processor     |                         "cc_number": "4187528062434"
processor     |                 },
processor     |                 "subscription": {
processor     |                         "plan": "Gold",
processor     |                         "status": "Idle",
processor     |                         "payment_method": "Bitcoins",
processor     |                         "term": "Annual"
processor     |                 }
processor     |         },
processor     |         "CreditCard": {
processor     |                 "uid": "3fedc306-5495-4a44-a3a9-0f600e056325",
processor     |                 "credit_card_number": "1211-1221-1234-2201",
processor     |                 "credit_card_expiry_date": "2025-06-10",
processor     |                 "credit_card_type": "jcb"
processor     |         }
processor     | }
```
## How to run

Just issue:
```
docker-compose up
```
and than you may visit http://localhost:15672/#/queues with credentials:


| username  | password |
| ------------- | ------------- |
| guest  | guest  |




## Some "weird" parts

In order to not set up 4 distinct repositories I used the following code structure:
```
.
├── docker-compose.yaml
├── README.md
└── src
    ├── credit-cards
    │   ├── go.mod
    │   ├── go.sum
    │   ├── internal
    │   │   └── ccard
    │   │       ├── logic.go
    │   │       └── structs.go
    │   └── main.go
    ├── pkg
    │   ├── rabbitmq
    │   │   └── rabbitmq.go
    │   └── rda
    │       └── random_data_api.go
    ├── processor
    │   ├── go.mod
    │   ├── go.sum
    │   ├── internal
    │   │   └── processor
    │   │       ├── processor.go
    │   │       └── structs.go
    │   └── main.go
    └── users
        ├── go.mod
        ├── go.sum
        ├── internal
        │   ├── user
        │   │   ├── logic.go
        │   │   └── structs.go
        │   └── users
        └── main.go
```

Using `docker-compose` volumes and the `module igiannoulas/golang-microservices/src` in `go.mod` we can mount them in a way that we can use and reference them. For example, the `processor` container will have the following upon startup:
```
root@f5d3ab810c44:/src# pwd
/src
root@f5d3ab810c44:/src# tree
.
|-- go.mod
|-- go.sum
|-- internal
|   |-- ccard
|   |   |-- logic.go
|   |   `-- structs.go
|   |-- credit-card
|   |-- processor
|   |   |-- processor.go
|   |   `-- structs.go
|   `-- user
|       |-- logic.go
|       `-- structs.go
|-- main.go
`-- pkg
    |-- rabbitmq
    |   `-- rabbitmq.go
    `-- rda
        `-- random_data_api.go
``` 
