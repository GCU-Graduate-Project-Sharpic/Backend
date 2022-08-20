# Backend

This is a backend implementation of Sharpic.

# API

This is a format of the request and response of the Sharpic server.

### /user - GET

Just check the cookie.

### /user/signup - POST

#### Request
```json
{
  "username": "username",
  "password": "password",
  "email": "email@email.com",
}
```

#### Response
```json
{
  "status": "signup success",
}
```

### /user/login - POST

#### Request
```json
{
  "username": "username",
  "password": "password",
  "email": " ",
}
```

#### Response
```json
{
  "status": "login success",
}
```

### /user/logout - POST

#### Response
```json
{
  "status": "logout success",
}
```