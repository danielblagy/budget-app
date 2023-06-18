# budget-app REST API Documentation

* [Introduction](#introduction)
* [Access](#access)
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
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VybmFtZTIiLCJleHAiOjE2ODYxNzIwNzgsIm5iZiI6MTY4NjE3MTE3OCwiaWF0IjoxNjg2MTcxMTc4fQ.Z0SBpqTK_mkHTwP9Re0xCiBA2iLyfgYpi_5PYp4vwAc",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VybmFtZTIiLCJleHAiOjE2ODc5ODU1NzgsIm5iZiI6MTY4NjE3MTE3OCwiaWF0IjoxNjg2MTcxMTc4fQ.YIz2yiZa94FOGH0qjFOu9FofbANw_mGmHAU1qDxyxqk"
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