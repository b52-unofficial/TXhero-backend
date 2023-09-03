# TXHero-backend
## Prerequisite
- Golang version 1.20
- Docker

## State Manager
- Update status in the database (dashboard, ingress controller information).
- Update bidding results.
- Update the status in the smart contract (reward).
- Change routing in the ACL manager (builder).

## Dashboard API
[API test link](https://www.postman.com/lunar-module-saganist-38772117/workspace/ethcon/collection/24650932-aae92a2c-8a9d-4c7e-8b8a-bde05fd643ed?action=share&creator=24650932)

|Path|Params|Description|
|------|---|---|
|/tx/accumulated_info||Summary information of all transactions|
|/tx/metadata|address string|User metadata|
|/tx/user|address string|User transaction information|
|/bid/current_round||Current bid round|
|/bid/round_info|round string|Round info|
|/bid/rounds|round string|Builder information from the previous round|
|/tx/chart_info|address string, period string|Chart info|

## Getting Started
- Dashboard API
  ```
  docker build . -t [name]:[tag]
  docker run -d -p 3002:3002 [image]:[tag]
  ```
- State Manager
  ```
  docker build . -t name:tag
  docker run -d [imgae]:[tag]
  ```


