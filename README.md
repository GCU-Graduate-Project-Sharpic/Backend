# Backend

This is a backend implementation of Sharpic.

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

Just check the cookie.

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

Just check the cookie.

##### Response
```json
{
  "list": [],
}
```

### /album/:albumId - GET

#### Request

Just check the cookie and the url param.

##### Response
```json
{
  "username": "username",
  "title": "title",
  "imageIds": [],
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

Just check the cookie.

#### Response

`Content-Type: image/png`

### /image/new/:albumId/:up - POST

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
