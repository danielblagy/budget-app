# budget-app REST API Documentation

* [Introduction](#introduction)
* [Access](#v1access)
* [Users](#v1users)
* [Categories](#v1categories)
* [Test Data for Manual Testing](#test-data)

## Introduction

This documentation provides a description of all available API handlers.

To use most of the API handlers you need to be logged in using /access/login, see [/access](#access) section for more information.

1. To create a user, use [POST /users](#post-v1users).
2. Then log in with created user's credentials via [POST /access/login](#post-v1accesslogin).
3. You are in!

Test user with test data is automatically created on migrate-up, check out [test data section](#test-data) for more information.

## v1/access

The access section. Used to authenticate the api user.

### POST v1/access/login

Body
```
{
    "username": "testusername2",
    "password": "password2"
}
```

Success response
```
{
    "successfully logged in"
}
```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | username or password are not valid
| Not Found      | 404  | user with username not found
| Forbidden | 403 | password is incorrect

### POST v1/access/logout

🔑 Requires user to be logged in.

1. Deletes http-only cookies storing jwt access and refresh tokens.
2. Adds jwt tokens to the blacklist, rendering them invalid.

### POST v1/access/refresh

🔑 Requires user to be logged in.

1. Blacklists current access and refresh tokens.
2. Issues new access and refresh token pair.

## v1/users

TODO

### POST v1/users

Create a new user.

Body
```
{
    "username": "testusername2",
    "email": "testemail2@mail.com",
    "full_name": "Test Name 2",
    "password": "password2"
}
```

Success response
```
{
    "username": "testusername3",
    "email": "testemail3@mail.com",
    "full_name": "Test Name 3"
}
```

Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | invalid json, username may only contain letters, numbers, underscores, and dashes
| Conflict      | 409  | username or email is taken
| Forbidden | 403 | password is incorrect

## v1/categories

Budget categories. Used by users to analyze expenses and incomes. 

### GET v1/categories/:type

🔑 Requires user to be logged in.

Get all user's categories of type. 

Possible types
```
*/expense
*/income
```
Success response
```
[
  {
    "id": 1,
    "user_id": "testusername111",
    "name": "day job",
    "type": "income"
  },
  {
    "id": 2,
    "user_id": "testusername111",
    "name": "vigilante",
    "type": "income"
  }
]
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | category type is not valid
| Unauthorized      | 401  | token has expired or user not authorized
| Not Found | 404 | categories not found

### GET v1/categories/by_id/:id

🔑 Requires user to be logged in.

Get user's category by category id.

Input example
```
GET v1/categories/by_id/1
```
Success response
```
{
  "id": 1,
  "user_id": "testusername111",
  "name": "day job",
  "type": "income"
}
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | category id is not valid
| Unauthorized      | 401  | token has expired or user not authorized
| Not Found | 404 | category not found

### POST v1/categories

🔑 Requires user to be logged in.

Create category

Body
```
{
  "name": "some category name",
  "type": "income"
}
```
Success response
```
{
  "id": 6,
  "user_id": "testusername111",
  "name": "some category name",
  "type": "income"
}
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | name or type is not valid or missing
| Unauthorized      | 401  | token has expired or user not authorized

### PUT v1/categories

🔑 Requires user to be logged in.

Update category

Body
```
{
  "id": 3,
  "name": "food",
}
```
Success response(updated data)
```
{
  "id": 3,
  "user_id": "testusername111",
  "name": "food",
  "type": "expense"
}
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | id or name is not valid or missing
| Unauthorized      | 401  | token has expired or user not authorized
| Not Found | 404 | category not found

###  DELETE v1/categories/:id

🔑 Requires user to be logged in.

Delete category by id

Input example
```
DELETE v1/categories/2
```
Success response (deleted category)
```
{
  "id": 2,
  "user_id": "testusername111",
  "name": "vigilante",
  "type": "income"
}
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | category id is not valid 
| Unauthorized      | 401  | token has expired or user not authorized
| Not Found | 404 | category not found



## Test Data

On migrate-up, the user and some test user data is created.

To start testing with test user, log in via [POST v1/access/login](#post-v1accesslogin).
Body
```
{
    "username": "testusername111",
    "password": "password111"
}
```