# Backend

This is a backend implementation of Sharpic.

# API

This is a format of the request and response of the Sharpic server.

### /user - GET

#### Response
```json
{
  "username": "username",
  "email": "email@email.com",
}
```

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
  "id": "id",
}
```

### /user/login - POST

#### Request
```json
{
  "username": "username",
  "password": "password",
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