A service to support a retailer with hosting his store

Service provided to the retailer:
1) Add new products
2) Update quantity and price of products
3) Get a list of the products available
4) Delete any product

Service provided to the customer:
1) Create a new account
2) Place an order
3) Update account and order details
4) Delete account and order details

A cooldown time of 30sec is provided to the retailer after a customer places any order.
All the subsequent orders placed within 30sec would be executed after 30sec.

If the order quantity is greater than the quanitity available, it will result in a failed order.

Instructions to use the services:
1) Execute 'go run main.go' which will launch the service
2) Head to URL /product for product services
3) Head to URL /customer for customer services
4) Head to URL /order for order services

