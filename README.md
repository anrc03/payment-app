PAYMENT APPLICATION API
---
This API enables customer to do payment to merchant. No limit for payment amount. Customer needs to be registered and logged in to do payment.
<br>
<br>
Link to API Postman Documentation:
https://documenter.getpostman.com/view/31825310/2s9YytggMj
<br>
<br>
How to use the API:
* Download and extract the project
* Create a database, database name must match with the one on .env file. If you want to create a different name, please don't forget to change it in the .env file also
* Run the main file using go run main.go
* Register a user using either /signup endpoint or /register endpoint. /register endpoint requires you to enter more data as it's an endpoint for customer registration
* Login using the same credential
* You can now do any get all operation and payment operation. Your token is automatically saved as cookies. The token has an expiry date of 30 days.
* If you're logged out you CANNOT do payment or get operation
---

## Technologies
This project is created with:
* Go 1.21.6
* Gin Web Framework
* GORM
* JWT
* PostgreSQL
