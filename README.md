
# MyAPI Documentation

## Introduction
Welcome to the documentation for MyAPI. This API provides endpoints for user authentication, managing products, and managing shopping carts.

## Authentication

### Login
- **Endpoint:** `/api/v1/auth/new-login`
- **Method:** `POST`
- **Description:** Authenticate user with username and password.
- **Request Body:**
  ```json
  {
      "username": "username",
      "password": "your_password"
  }
Response:
Success:
{
    {
    "code": 200,
    "message_code": 0,
    "message": "Success",
    "data": {
        "user": {
            "customer_id": "",
            "user_name": "",
            "user_email": "",
            "created_at": "",
            "session_id": ""
        },
        "access_token": {
            "type": "bearer",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ4MDgzNzQsImlhdCI6MTcxNDcyMTk3NCwicmVmcmVzaF90b2tlbiI6IjBiZTM3YWY4ZjM1N2M5NDkzNTM5Zjc0MTBhMjkwMDkwNWZjNDFhMzUyMDcyZDljMTZkODMyMDA4NGU0MWU2NzYiLCJzZXNzaW9uIjoiIiwidXNlciI6eyJjdXN0b21lcl9pZCI6ImZmNjNkMDhkLWZhNTMtNDVhYi1iOGZiLTY4OGM0OTk3M2M4MyIsInVzZXJfbmFtZSI6ImNsYXJlX2FkbWluIiwidXNlcl9lbWFpbCI6IiIsImNyZWF0ZWRfYXQiOiIyMDI0LTA1LTAyIDE4OjU4OjQyLjQyMjAxNSIsInNlc3Npb25faWQiOiIifX0.IVUppVCFGBZxChMX2EFwOJ_a2Aw_TG85eABPZWxoq2A",
            "refresh_token": "0be37af8f357c9493539f7410a2900905fc41a352072d9c16d8320084e41e676"
        }
    }
}
}
Error:
Copy code
{
    "code": 400,
    "message_code": 13,
    "message": "Incorrect username or password",
    "data": null
}
Products
Product List by Category
Endpoint: /api/v1/product?category
Method: GET
Description: Get a list of products by category.
Parameters:
category (required): Category of products to retrieve.
Response:
{
    {
    "code": 200,
    "message_code": 0,
    "message": "Successfully",
    "data": [
        {
            "product_id": "30e3b319-3de1-41a5-a0ce-e9e69e9d281a",
            "name": "Sikat Gigi",
            "stock": 12,
            "price": 2781,
            "category": "Kesehatan, Peralatan"
        },
        {
            "product_id": "18798852-fdc9-410d-ad82-a9c5bd151af4",
            "name": "Tissu",
            "stock": 52,
            "price": 1992,
            "category": "Kesehatan"
        }
    ]
}
}
Carts
Cart List
Endpoint: api/v1/cart?customer_id=
Method: GET
Description: Get a list of items in the user's shopping cart.
Headers:
Authorization: Bearer your_jwt_token
Response:
json
Copy code
{
    "code": 200,
    "message_code": 0,
    "message": "Successfully",
    "data": [
        {
            "cart_id": "",
            "name": "",
            "quantity": ""
        }
    ]
}
Add to Cart
Endpoint: api/v1/cart/add
Method: POST
Description: Add a product to the user's shopping cart.
Headers:
Authorization: Bearer your_jwt_token
Request Body:
{
    "customer_id":"",
    "product_id":"",
    "quantity":4
}