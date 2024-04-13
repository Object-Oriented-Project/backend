# backend


**API Documentation**

## Endpoints

**BaseUrl** `localhost:3001`

Here's a quick reference table summarizing the main API sections:

| Section        | Description                                              | Group Path |
|----------------|------------------------------------------------------------|-------------|
| Health Check    | Returns "OK" to indicate the server is running.         | `/_hc`       |
| Menu Management | CRUD operations (Create, Read, Update, Delete) for menus. | `/api/menu`   |
| User Management | CRUD operations for user accounts.                   | `/api/user`  |
| Authentication  | Login                                                  | `/api/auth`  |
| Kitchen        | Real-time communication for kitchen staff (WebSocket).     | `/api/kitchen` |

**Detailed Documentation by Section:**

### 1. Health Check

* **Route:** `/_hc`
* **Method:** `GET`
* **Description:** Performs a basic health check to verify the server is running.
* **Response:**
    * Status Code: 200 OK
    * Body: String containing "OK"

### 2. Menu Management (`/api/menu`)

* **Authentication:** Required for all operations (details in `/api/auth` section)
* **List Menus (GET /api/menu)**
* **Create Menu (POST /api/menu)**
* **Get Menu Details (GET /api/menu/:id)**
* **Update Menu (PUT /api/menu/:id)**
* **Delete Menu (DELETE /api/menu/:id)**
* **Example JSON**
```JSON
{
    "ItemName": "Cake",
    "ItemPriceSmall":-1,
    "ItemPriceMedium": -1,
    "ItemPriceLarge": 129,
    "ItemType": "Bakery",
    "IsRecommended": false,
    "PictureName": "bakery-4.png",
    "ItemDescription": "Indulge in our selection of delectable cakes."
}
```

### 3. User Management (`/api/user`)

* **Authentication:** Required for all operations (details in `/api/auth` section)
  * **List User (GET /api/user)**
  * **Create User (POST /api/user)**
* **Example JSON**
```JSON
  {
  "username":"admin",
  "password":"admin1234"
  }
  ```


### 4. Authentication (`/api/auth`)
  * **Login (POST /api/auth)**
  * **Example JSON**
  ```JSON
  {
  "username":"admin",
  "password":"admin1234"
  }
  ```

### 5. Kitchen (WebSocket) (`/api/kitchen`)
  * **Send Order (POST /api/kitchen) [RestAPI]**
  * **View Order (GET /api/kitchen/ws/kitchen) [WebSocket]** 
  * **Example JSON** 
  ```JSON
  {
  "from": "kitchen",
  "to": "kitchen",
  "customer_id": 12345,
  "customer_name": "John Doe",
  "order_id": 98765, 
  "order_status": "ONGOING",
  "order_items": [
    {
      "menu_id": 1,
      "quantity": 1,
      "size": "LARGE",
      "sweetness": "NORMAL",
      "spicy_level": "NORMAL",
      "note": "Extra cheese, please"
    },
    {
      "menu_id": 2,
      "quantity": 2,
      "size": "SMALL",
      "sweetness": "LESS",
      "spicy_level": "LESS",
      "note": ""
    }
  ]
}
```
**Additional Notes:**
* have any problem contact AK