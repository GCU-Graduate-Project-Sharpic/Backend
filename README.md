# Backend

This is a backend implementation of Sharpic.

The sharpic server manages user account information and provides login logout API. And by providing album and image api, it serves as a relay so that users can upload images to the database and super resolution processing can operate.

To use full sharpic service, see [sharpic](https://github.com/GCU-Sharpic/sharpic)

## Run test Backend

```zsh
docker compose -f ./docker-compose-test.yml up --build -d
```

# Structure

![Screenshot from 2023-05-31 16-36-46](https://github.com/GCU-Sharpic/sharpic-server/assets/20539422/b3585f46-5507-43a7-b2d9-99cd5b2bb8b6)

The overall process of sharpic server is as above. Below are the modules of sharpic server.

`main.go` : API paths are defined and appropriate handlers are called.

`/handler` : Parse the json data and do whatever you need. If you need to call the DB, call the database module.

`/database` : Connects the database and provides necessary functionality.

`/type`: Data types used in sharpic server are defined.

# APIs

This is a format of the request and response of the Sharpic server.

- [/singup - POST](#signup---post)
- [/login - POST](#login---post)
- [/logout - POST](#logout---post)
- [/user - GET](#user---get)
- [/album/list - GET](#albumlist---get)
- [/album/:albumId - GET](#albumalbumid---get)
- [/album/new - POST](#albumnew---post)
- [/image/:imageId - GET](#imageimageid---get)
- [/image/processed/:imageId - GET](#imageprocessedimageid---get)
- [/image/info/:imageId - GET](#imageinfoimageid---get)
- [/image/new/:albumId - POST](#imagenewalbumid---post)
- [/image/up/:imageId/:newUp - PATCH](#imageupimageidnewup---patch)

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
