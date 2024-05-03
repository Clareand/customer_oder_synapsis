MyAPI Documentation
===================

Introduction
------------
Welcome to the documentation for MyAPI. This API provides endpoints for user authentication, managing products, and managing shopping carts.

Authentication
--------------
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
  ```
- **Response:**
  - Success:
    ```json
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
                "token": "your_access_token",
                "refresh_token": "your_refresh_token"
            }
        }
    }
    ```
  - Error:
    ```json
    {
        "code": 400,
        "message_code": 13,
        "message": "Incorrect username or password",
        "data": null
    }
    ```

Products
--------
### Product List by Category
- **Endpoint:** `/api/v1/product?category`
- **Method:** `GET`
- **Description:** Get a list of products by category.
- **Parameters:**
  - `category` (required): Category of products to retrieve.
- **Response:**
  ```json
  {
      "code": 200,
      "message_code": 0,
      "message": "Successfully",
      "data": [
          {
              "product_id": "",
              "name": "",
              "stock": ,
              "price": ,
              "category": ""
          },
          {
              "product_id": "",
              "name": "",
              "stock": ,
              "price": ,
              "category": ""
          }
      ]
  }
  ```

Carts
-----
### Cart List
- **Endpoint:** `/api/v1/cart?customer_id=`
- **Method:** `GET`
- **Description:** Get a list of items in the user's shopping cart.
- **Headers:**
  - `Authorization: Bearer your_jwt_token`
- **Response:**
  ```json
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
  ```

### Add to Cart
- **Endpoint:** `/api/v1/cart/add`
- **Method:** `POST`
- **Description:** Add a product to the user's shopping cart.
- **Headers:**
  - `Authorization: Bearer your_jwt_token`
- **Request Body:**
  ```json
  {
      "customer_id": "",
      "product_id": "",
      "quantity": 4
  }
  ```
