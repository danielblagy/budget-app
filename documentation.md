# budget-app REST API Documentation

* [Introduction](#introduction)
* [Access](#access)
* [Entries](#v1entries)
* [Test Data for Manual Testing](#test-data)

## Introduction

This documentation provides a description of all available API handlers.

To use most of the API handlers you need to be logged in using /access/login, see [/access](#access) section for more information.

1. To create a user, use [POST /users](#post-users).
2. Then log in with created user's credentials via [POST /access/login](#post-accesslogin).
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

## v1/entries

Entries of user's expenses and incomes. Used by users to view the history of financial transactions. 

### GET v1/entries/:type

🔑 Requires user to be logged in.

Get all user's entries of type. 

Possible types
```
*/expense
*/income
```
Success response
```
[
  {
    "id": 3,
    "user_id": "testusername111",
    "category_id": 25,
    "amount": 1257.6,
    "date": "2006-12-08T00:00:00Z",
    "description": "bread",
    "type": "expense"
  },
  {
    "id": 4,
    "user_id": "testusername111",
    "category_id": 3,
    "amount": 850.2,
    "date": "2023-06-09T00:00:00Z",
    "description": "milk and eggs",
    "type": "expense"
  }
]
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | entry type is not valid
| Unauthorized      | 401  | token has expired or user not authorized

### GET v1/entries/by_id/:id

🔑 Requires user to be logged in.

Get user's entry by entry id.

Input example
```
GET v1/entries/by_id/3
```
Success response
```
{
  "id": 3,
  "user_id": "testusername111",
  "category_id": 25,
  "amount": 1257.6,
  "date": "2006-12-08T00:00:00Z",
  "description": "bread",
  "type": "expense"
}
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | entry id is not valid
| Unauthorized      | 401  | token has expired or user not authorized
| Not Found | 404 | entry not found

### POST v1/entries

🔑 Requires user to be logged in.

Create entry

Body
```
{
    "category_id": 25,
    "amount": 1257.60,
    "date": "2023-07-02",
    "description" : "bread",
    "type": "expense"
}
```
Success response
```
{
  "id": 8,
  "user_id": "testusername111",
  "category_id": 25,
  "amount": 1257.6,
  "date": "2023-07-02T00:00:00Z",
  "description": "bread",
  "type": "expense"
}
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | params are not valid or missing
| Unauthorized      | 401  | token has expired or user not authorized

### PUT v1/entries/:entry_id

🔑 Requires user to be logged in.

Update entry

Input example
```
PUT v1/entries/3
```
Body
```
{
    "category_id": 25,
    "amount": 1800.60,
    "date": "2023-07-01",
    "description" : "bread and eggs",
    "type": "expense"
}
```
Success response(updated data)
```
{
  "id": 3,
  "user_id": "testusername111",
  "category_id": 25,
  "amount": 1800.6,
  "date": "2023-07-01T00:00:00Z",
  "description": "bread and eggs",
  "type": "expense"
}
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | params are not valid or missing
| Unauthorized      | 401  | token has expired or user not authorized
| Not Found | 404 | entry not found

###  DELETE v1/entries/:id

🔑 Requires user to be logged in.

Delete entry by id

Input example
```
DELETE v1/entries/3
```
Success response (deleted entry)
```
{
  "id": 3,
  "user_id": "testusername111",
  "category_id": 25,
  "amount": 1800.6,
  "date": "2023-07-01T00:00:00Z",
  "description": "bread and eggs",
  "type": "expense"
}
```
Errors
| Error             | Code          | Description   |
| -------------     | ------------- | -             |
| Bad Request         | 400  | entry id is not valid 
| Unauthorized      | 401  | token has expired or user not authorized
| Not Found | 404 | entry not found

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