# Backend

This is a backend implementation of Sharpic.

# User API

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

# Image API

### /image - POST

#### Request

Form-data format

```json
"images": image_file,
"images": image_file,
"images": image_file,
...
```

#### Response

```json
{
    "status": "files uploaded!"
}
```

### /image/list - GET

#### Request

Just check the cookie.

#### Response

```json
"list": [1, 2, 3]
```

### /image/:id - GET

#### Requset

Just check the cookie.

#### Response

`Content-Type: image/png`