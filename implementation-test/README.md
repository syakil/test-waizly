---
title: waizly
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# waizly

Base URLs:

# Authentication

- HTTP Authentication, scheme: bearer

# Auth

## POST Register

POST /api/register

> Body Parameters

```json
{
  "name": "Syakil",
  "email": "syakil@example.com",
  "password": "12345678",
  "password_confirmation": "12345678"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» name|body|string| yes |none|
|» email|body|string| yes |none|
|» password|body|string| yes |none|
|» password_confirmation|body|string| yes |none|

> Response Examples

> Register

```json
{
  "message": "User registered successfully"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Register|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» message|string|true|none||none|

## POST Logout

POST /api/logout

> Body Parameters

```json
{}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> Logout

```json
{
  "message": "Logged out successfully"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Logout|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» message|string|true|none||none|

## POST Login

POST /api/login

> Body Parameters

```json
{
  "email": "john@example.com",
  "password": "password"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» email|body|string| yes |none|
|» password|body|string| yes |none|

> Response Examples

> Login

```json
{
  "access_token": "4|5bBJbLcs4YhFr1El5oZEVJpz0EvbGlKI6gj9lwK3dfd37356",
  "token_type": "Bearer"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Login|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» access_token|string|true|none||none|
|» token_type|string|true|none||none|

# User

## GET get user

GET /api/user

> Response Examples

> get user

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "email_verified_at": null,
  "created_at": "2024-05-16T17:20:26.000000Z",
  "updated_at": "2024-05-16T17:20:26.000000Z"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|get user|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» id|integer|true|none||none|
|» name|string|true|none||none|
|» email|string|true|none||none|
|» email_verified_at|null|true|none||none|
|» created_at|string|true|none||none|
|» updated_at|string|true|none||none|

# Product

## GET Get Products

GET /api/products

> Response Examples

> Get Products

```json
[
  {
    "id": 2,
    "name": "New Product",
    "description": "Description for new product",
    "price": "150.00",
    "stock": 20,
    "created_at": "2024-05-16T18:38:22.000000Z",
    "updated_at": "2024-05-16T18:38:22.000000Z"
  }
]
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Get Products|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» id|integer|false|none||none|
|» name|string|false|none||none|
|» description|string|false|none||none|
|» price|string|false|none||none|
|» stock|integer|false|none||none|
|» created_at|string|false|none||none|
|» updated_at|string|false|none||none|

## POST Add Product

POST /api/products

> Body Parameters

```json
{
  "name": "New Product 2",
  "description": "Description for new product",
  "price": 150,
  "stock": 20
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» name|body|string| yes |none|
|» description|body|string| yes |none|
|» price|body|integer| yes |none|
|» stock|body|integer| yes |none|

> Response Examples

> Add Product

```json
{
  "name": "New Product 2",
  "description": "Description for new product",
  "price": 150,
  "stock": 20,
  "updated_at": "2024-05-16T18:41:20.000000Z",
  "created_at": "2024-05-16T18:41:20.000000Z",
  "id": 3
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Add Product|Inline|

### Responses Data Schema

HTTP Status Code **201**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» name|string|true|none||none|
|» description|string|true|none||none|
|» price|integer|true|none||none|
|» stock|integer|true|none||none|
|» updated_at|string|true|none||none|
|» created_at|string|true|none||none|
|» id|integer|true|none||none|

## PUT Update Product

PUT /api/products/1

> Body Parameters

```json
{
  "name": "Updated Product",
  "description": "Updated description for product",
  "price": 200,
  "stock": 30
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» name|body|string| yes |none|
|» description|body|string| yes |none|
|» price|body|integer| yes |none|
|» stock|body|integer| yes |none|

> Response Examples

> Update

```json
{
  "id": 1,
  "name": "Updated Product",
  "description": "Updated description for product",
  "price": 200,
  "stock": 30,
  "created_at": "2024-05-16T18:37:54.000000Z",
  "updated_at": "2024-05-16T18:39:51.000000Z"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Update|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» id|integer|true|none||none|
|» name|string|true|none||none|
|» description|string|true|none||none|
|» price|integer|true|none||none|
|» stock|integer|true|none||none|
|» created_at|string|true|none||none|
|» updated_at|string|true|none||none|

## DELETE Delete Product

DELETE /api/products/3

> Body Parameters

```json
{}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> Delete

```json
{
  "success": "Succed to delete product"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Delete|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» success|string|true|none||none|

# Data Schema

