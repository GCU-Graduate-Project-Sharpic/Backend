# Backend

This is a backend implementation of Sharpic.

To use full sharpic service, see [sharpic](https://github.com/GCU-Sharpic/sharpic)

## Run test Backend

```zsh
docker compose -f ./docker-compose-test.yml up --build -d
```

# API

This is a format of the request and response of the Sharpic server.

## Authentication API

To make a request to sharpic-server, you must first log in.<br>
Below are the APIs related to login.


### /signup - POST

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

### /login - POST

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

### /logout - POST

#### Response
```json
{
  "status": "logout success",
}
```

## User API

Provides functions related to user information.

### /user - GET

#### Request 

Check `userId` in cookie.

##### Response
```json
{
  "username": "username",
  "email": "email@email.com",
}
```

## Album API

### /album/list - GET

#### Request

Check `userId` in cookie.

##### Response
```json
{
  "list": [
    {
      "id": "int",
      "username": "username",
      "title": "title",
      "imageIds": "[]int",
    }
  ],
}
```

### /album/:albumId - GET

#### Request

Check `userId` in cookie.

##### Response
```json
{
  "username": "username",
  "title": "title",
  "imageIds": "[]int",
}
```

### /album/new - POST

#### Request
```json
{
  "username": "username",
  "title": "title",
}
```

#### Response
```json
{
  "status": "new album success",
}
```

## Image API

### /image/:imageId - GET

#### Requset

Check `userId` in cookie.

#### Response

`Content-Type: image/png`

### /image/processed/:imageId - GET

#### Requset

Check `userId` in cookie.

#### Response

`Content-Type: image/png`

### /image/info/:imageId - GET

#### Requset

Check `userId` in cookie.

#### Response

```json
{
  "fileName":   "string",
  "size":       "int64",
  "added_date": "time.Time",
  "up":         "int",
  "status":     "bool",
}
```

### /image/new/:albumId - POST

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
    "status": "images uploaded!"
}
```

### /image/up/:imageId/:newUp - PATCH

#### Request

Just check the cookie.

#### Response

```json
{
    "status": "up changed"
}
```
