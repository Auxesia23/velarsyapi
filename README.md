# Velarsy API Documentation

## VelarsyAPI

## Auth

### Register User

**Method:** `POST`  
**URL:** `{{host}}/auth/register`  
**Content-Type:** `application/json`

**Request Body:**
```json
{
    "username": "str",
    "password": "str"
}
```

### Edit User

**Method:** `PUT`  
**URL:** `{{host}}/auth/users/{user_id}`  
**Content-Type:** `application/json`

**Request Body:**
```json
{
    "username": "str",
    "password": "str"
}
```

### Login

**Method:** `POST`  
**URL:** `{{host}}/auth/login`  
**Content-Type:** `application/json`

**Request Body:**
```json
{
    "username": "str",
    "password": "str"
}
```

### Delete User

**Method:** `DELETE`  
**URL:** `{{host}}/auth/users/{user_id}`  
**Content-Type:** `No Body`


### Gett All Users

**Method:** `GET`  
**URL:** `{{host}}/auth/users`  
**Content-Type:** `No Body`




## Projects

### Create

**Method:** `POST`  
**URL:** `{{host}}/works/{work_id}/projects`  
**Content-Type:** `multipart/form`

**Request Body:**
```form
{
    "name"              : "text"
    "about_brand"       : "text"
    "design_execution"  : "text"
    "image"             : "file"
}
```

### Get Single

**Method:** `GET`  
**URL:** `{{host}}/projects/{slug}`  
**Content-Type:** `No Body`


### Update

**Method:** `PUT`  
**URL:** `{{host}}/projects/{project_id}`  
**Content-Type:** `multipart/form`

**Request Body:**
```form
{
    "name"              : "text"
    "about_brand"       : "text"
    "design_execution"  : "text"
    "image"             : "file"
}
```

### Delete

**Method:** `DELETE`  
**URL:** `{{host}}/projects/{project_id}}`  
**Content-Type:** `No Body`




## Images

### Create

**Method:** `POST`  
**URL:** `{{host}}/projects/{project_id}/images`  
**Content-Type:** `multipart/form`

**Request Body:**
```form
{
    "image" : "[]file"
}
```

### Delete

**Method:** `DELETE`  
**URL:** `{{host}}/images/{image_id}`  
**Content-Type:** `No Body`




## Work

### Create

**Method:** `POST`  
**URL:** `{{host}}/works/`  
**Content-Type:** `multipart/form`

**Request Body:**
```form
{
    "title" : "str"
    "image" : "file"
}
```

### Get One

**Method:** `GET`  
**URL:** `{{host}}/works/{slug}`  
**Content-Type:** `No Body`


### Get All

**Method:** `GET`  
**URL:** `{{host}}/works`  
**Content-Type:** `No Body`

### Update

**Method:** `PUT`  
**URL:** `{{host}}/works/{work_id}`  
**Content-Type:** `multipart/form`

**Request Body:**
```form
{
    "title" : "str"
    "image" : "file"
}
```


### Delete

**Method:** `DELETE`  
**URL:** `{{host}}/works/{work_id}`  
**Content-Type:** `No Body`



## Services


### Create

**Method:** `POST`  
**URL:** `{{host}}/services`  
**Content-Type:** `application/json`

**Request Body:**
```json
{
    "title"         : "str",
    "description"   : "str",
    "icon"          : "str"
}
```


### Get

**Method:** `GET`  
**URL:** `{{host}}/services`  
**Content-Type:** `No Body`


### Update

**Method:** `PUT`  
**URL:** `{{host}}/services/{service_id}`  
**Content-Type:** `application/json`

**Request Body:**
```json
{
    "title": "str",
    "description": "str",
    "icon": "str"
}
```


### Delete

**Method:** `DELETE`  
**URL:** `{{host}}/services/{service_id}`  
**Content-Type:** `No Body`







