
# Backend Engineer - Asset Transaction API
This is a backend API for an asset transaction service, built using Golang, GORM, Gin, MySQL, and Hexagonal Architecture. The API provides functionality for managing assets, wallet, transactions, and users. It is designed to be scalable, maintainable, and performant.


## Technologies Used
- Golang - A programming language designed for building scalable, high-performance applications
- GORM - A Golang Object-Relational Mapping library for working with databases
- Gin - A web framework for building APIs in Golang
- MySQL - A popular open-source relational database management system
- Hexagonal Architecture - A software design pattern that emphasizes separation of concerns, testability, and maintainability.


## Features
- User management: register, login, refresh-token, create, update, and delete
- Asset management: get, create, update, and delete assets
- Wallet management: get, create, update, and delete wallet
- Transaction management: transfer assets
- Authentication and authorization: ensure that only authorized users can access specific features

## Installation
Clone project

```bash
  git clone https://github.com/ezizull/tennet-gethired.git
  cd tennet-gethired
```

Install Dependency

```bash
  go mod tidy
```

Run Service 
```bash
  go run main.go
```

Open in the web browser

```bash
  localhost:4000/{{version}}/swagger/index.html#/
```

**Save link as** and **Import to postman**

[![Download Collection](https://img.shields.io/badge/Download%20Collection-EF5B25?style=for-the-badge&logo=postman&logoColor=white)](https://github.com/ezizull/tennet-gethired/blob/master/docs/tennet-gethired.postman_collection.json)  

[![Download Environment](https://img.shields.io/badge/Download%20Environment-EF5B25?style=for-the-badge&logo=postman&logoColor=white)](https://github.com/ezizull/tennet-gethired/blob/master/docs/tennet-gethired%40localhost.postman_environment.json)

## API Reference

### Authorization
#### Login
```bash
  POST /{{version}}/auth/login
```
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email`   | `string` | **Required**. Your email   |


#### Register
```bash
  POST /{{version}}/user
```
| Parameter  | Type     | Description                       |
| :--------- | :------- | :-------------------------------- |
| `username` | `string` | **Required**. Your name account   |
| `email`    | `string` | **Required**. Your email account  |
| `firstName`| `string` | **Required**. Your first name account   |
| `lastName` | `string` | **Required**. Your last name account   |
| `password` | `string` | **Required**. Your password account   |


#### Refresh Token
```bash
  POST /{{version}}/auth/access-token
```
| Parameter      | Type     | Description                       |
| :---------     | :------- | :-------------------------------- |
| `refreshToken` | `string` | **Required**. Jwt refresh token   |


### User
#### Get User
```bash
  GET /{{version}}/user/:id
```
Takes id and returns the user by the id.

#### Update User
```bash
  PATCH /{{version}}/user/:id
```
| Parameter  | Type     | Description                       |
| :--------- | :------- | :-------------------------------- |
| `username` | `string` | **Optional**. Your name account   |
| `email`    | `string` | **Optional**. Your email account  |
| `firstName`| `string` | **Optional**. Your first name account   |
| `lastName` | `string` | **Optional**. Your last name account   |
| `password` | `string` | **Optional**. Your password account   |

Takes id and update body. returns the user updated.

#### Delete User
```bash
  DELETE /{{version}}/user/:id
```
Takes id and delete the user by the id.


### Assets
Assets can access only after **login** first.
| Header          | Value           |
| :---------      | :-------        | 
| `Authorization` | JwtAccessToken  | 

#### Get Assets
```bash
  GET /{{version}}/assets
```
Get all asset by limit 20 and return as pagination.

#### Get Asset
```bash
  GET /{{version}}/assets/:id
```
Get asset returns the asset by the id.

#### Create Asset
```bash
  POST /{{version}}/assets
```
| Parameter   | Type     | Description                       |
| :---------  | :------- | :-------------------------------- |
| `wallet_id` | `int`    | **Required**. Your wallet_id asset|
| `name`      | `string` | **Required**. Your name asset     |
| `symbol`    | `string` | **Required**. Your symbol asset   |
| `network`   | `string` | **Required**. Your network asset  |
| `address`   | `string` | **Required**. Your address asset  |
| `balance`   | `decimal`| **Required**. Your balance asset  |

#### Update Asset
```bash
  PATCH /{{version}}/assets/:id
```
| Parameter   | Type     | Description                       |
| :---------  | :------- | :-------------------------------- |
| `wallet_id` | `int`    | **Optional**. Your wallet_id asset|
| `name`      | `string` | **Optional**. Your name asset     |
| `symbol`    | `string` | **Optional**. Your symbol asset   |
| `network`   | `string` | **Optional**. Your network asset  |
| `address`   | `string` | **Optional**. Your address asset  |
| `balance`   | `decimal`| **Optional**. Your balance asset  |

Takes id and update body. returns the asset updated.

#### Delete Asset
```bash
  DELETE /{{version}}/assets/:id
```
Takes id and delete the asset by the id.

### Wallet
Wallet can access only after **login** first.
| Header          | Value           |
| :---------      | :-------        | 
| `Authorization` | JwtAccessToken  | 

#### Get Wallets
```bash
  GET /{{version}}/wallet
```
Get all wallet by limit 20 and return as pagination.

#### Get Wallet
```bash
  GET /{{version}}/wallet/:id
```
Get wallet returns the wallet by the id.

#### Create Wallet
```bash
  POST /{{version}}/wallet
```
| Parameter   | Type     | Description                       |
| :---------  | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Your name wallet    |

#### Update Wallet
```bash
  PATCH /{{version}}/wallet/:id
```
| Parameter   | Type     | Description                       |
| :---------  | :------- | :-------------------------------- |
| `name`      | `string` | **Optional**. Your name wallet     |

Takes id and update body. returns the wallet updated.

#### Delete Wallet
```bash
  DELETE /{{version}}/wallet/:id
```
Takes id and delete the wallet by the id.

### Transaction
Transaction can access only after **login** first.
| Header          | Value           |
| :---------      | :-------        | 
| `Authorization` | JwtAccessToken  | 

#### Transaction Asset
```bash
  POST /{{version}}/transaction
```
| Parameter        | Type     | Description                                           |
| :--------------  | :------- | :---------------------------------------------------  |
| `src_wallet_id`  | `int`    | **Required**. Your src wallet id transaction          |
| `src_asset_id`   | `string` | **Required**. Your src asset id transaction           |
| `dest_wallet_id` | `string` | **Required**. Your destination wallet id transaction  |
| `amount`         | `string` | **Required**. Your amount transaction                 |
| `gas_fee`        | `string` | **Required**. Your gas fee transaction                |
| `total`          | `string` | **Required**. Your total transaction                  |

Update asset wallet_id and record to asset transaction. Return created asset transaction

### Documentation
#### Swagger 
##### Open in the web browser
```bash
  localhost:4000/{{version}}/swagger/index.html#/
```
Show all api documentation with firendly interface.

#### Postman
[![Download Collection](https://img.shields.io/badge/Download%20Collection-EF5B25?style=for-the-badge&logo=postman&logoColor=white)](https://github.com/ezizull/tennet-gethired/blob/master/docs/tennet-gethired.postman_collection.json)  

[![Download Environment](https://img.shields.io/badge/Download%20Environment-EF5B25?style=for-the-badge&logo=postman&logoColor=white)](https://github.com/ezizull/tennet-gethired/blob/master/docs/tennet-gethired%40localhost.postman_environment.json)

**Save link as** and **Import to postman**. Show all api documentation in postman.
