# Go-lang

UrbanKart is planning to launch an online store. In preparation for that they plan to develop
an application for online shopping, managing their delivery logistics
#1. Shopping
Develop an application that can be used for online shopping
1. Allow users to register with mobile numbers, and place(address)
// return true if Success or false if fails - like mobile number already exists / invalid
mobile number
AddUsers(user User) bool
2. Allow users to fetch list of products
3. //type - product type, such as mobile, tv etc
GetProducts(type string)
4. Allow users to order products
PlaceOrder(item Product, qty int)
5. Allow users to modify order
ChangeOrder((item Product, qty int)
6. Allow users to cancel order
ChangeOrder((item Product, qty int)
7. Allow users to fetch their order
//return array of prod , array of qty
GetOrders() ([] prod, [] qty)

#2 Delivery Logistics
1. Allow admin to register delivery reps , with mobile number and place
// return true if Success or false if fails - like mobile number already exists / invalid mobile
number
AddDeliveryRep(rep Rep) bool
2. Allow admin to assign a customer order to a delivery rep matching the place
Consider the rep should not have any other assignment
If no matching rep found return false
AssignOrder(order Order) bool
3. Allow admin to view delivery assignments for all reps
 ViewAssignments()
