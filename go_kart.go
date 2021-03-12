package main

import "fmt"

var productList []string

type Order struct {
	product_name string
	qty          int
	user         Users
}
type Users struct {
	User_name    string
	User_mobile  int64
	User_address string
	order        []Order
}
type Go_kart interface {
	authUser()
	showProduct()
	placeOrder()
	modifyOder()
	deleteOrder()
}

var user []Users
var orderlist []Order
var length int
var prodarr [6]string = [6]string{"iron", "fridge", "cooler", "iphone", "tv", "washing machine"}

func deleteOrder(curUser Users) int {
	var prodname string
	fmt.Println("enter the delete product name")
	fmt.Scanf("%s\n", &prodname)
	for j := 0; j < length; j++ {
		if curUser.order[j].product_name == prodname {
			curUser.order[j] = curUser.order[length-1]

			curUser.order = curUser.order[:length-1]
			length--
			break

		}

	}

	return length
}
func printProduct(curUser Users, length int) {
	if length > 0 {
		for i := 0; i < length; i++ {
			fmt.Printf("ordername%s\n", curUser.order[i].product_name)
		}
	} else {
		fmt.Println("\n************************ Go_kart is empty, it's time for you to order *********************** ")
	}
}

func modifyOder(curUser Users) {
	var qty int
	var modname string
	fmt.Printf("\nEnter the Product name to modify")
	fmt.Scanf("%s\n", &modname)
	var flag bool = false
	for i := 0; i < 6; i++ {
		if modname == prodarr[i] {
			fmt.Printf("\nEnter the quantity to modify\n")
			fmt.Scanf("%d\n", &qty)
			flag = true
		}
	}
	if flag {
		for j := 0; j < len(orderlist); j++ {
			if curUser.User_name == orderlist[j].user.User_name && orderlist[j].product_name == modname {
				orderlist[j].qty = qty
				fmt.Printf("username:%s\nproductname:%s\nqty:%d\n", orderlist[j].user.User_name,
					orderlist[j].product_name, orderlist[j].qty)
				break
			}
		}

	}
}

func placeOrder(curUser Users) Order {
	var prodName string
	var qty int
	fmt.Println("\nEnter the Product Name")
	fmt.Scanf("%s\n", &prodName)
	fmt.Println("\nEnter the Quantity")
	fmt.Scanf("%d\n", &qty)
	for i := 0; i < 6; i++ {
		if prodarr[i] == prodName {
			var order1 Order = Order{product_name: prodName, qty: qty, user: curUser}

			return order1
		}

	}
	return Order{}
}
func showProduct() {
	fmt.Printf("\nName of products in Go_kart\n")
	for _, v := range prodarr {
		fmt.Printf("%s\n", v)
	}
}

func authUser(mobile int64) Users {
	for i := 0; i < len(user); i++ {
		if user[i].User_mobile == mobile && mobile > 5999999999 && mobile < 9999999999 {
			return user[i]
		}
	}
	var newUser Users = createUser(mobile)
	user = append(user, newUser)
	return newUser
}
func createUser(mobile int64) Users {
	var name, address string
	fmt.Println("\n*****************Enter the name**********************")
	fmt.Scanf("%s\n", &name)
	fmt.Println("\n*****************Enter the address*******************")
	fmt.Scanf("%s\n", &address)
	newUser := Users{User_name: name, User_mobile: mobile, User_address: address}
	return newUser
}

func printOrder(curUser Users) {
	fmt.Println("************Print the order of cur user****************")
	for i := 0; i < length; i++ {
		fmt.Printf("username:%s\nproductname:%s\nQty:%d\n", curUser.User_name, curUser.order[i].product_name, curUser.order[i].qty)
	}

}

var curOrderList []Order
var orderplace Order

func main() {

	var mobile int64
	fmt.Println("\n**********************Enter the mobile number*************************")
	flag := true

	for flag {
		fmt.Scanf("%d\n", &mobile)
		if !(mobile > 5999999999 && mobile < 9999999999) {
			fmt.Println("\nEnter the valid mobile 10 digit number")

		} else {
			flag = false
		}

	}

	firstOrder := []Order{{product_name: "tv", qty: 56}, {product_name: "cooler", qty: 52}}
	firstUser := Users{User_name: "harsh", User_mobile: 9693880525, User_address: "E/12", order: firstOrder}
	user = append(user, firstUser)
	var curUser Users = authUser(mobile)
	var op int
	for true {
		fmt.Printf("\n")
		fmt.Println("Enter 1 to Get Products")
		fmt.Println("Enter 2 to Place Orders")
		fmt.Println("Enter 3 to Modify Products")
		fmt.Println("Enter 4 to Cancel Products")
		fmt.Println("Enter 5 to Print Products")
		fmt.Println("Enter 6 to Print Products ordered from me")
		fmt.Println("Enter 7 to Login to another account")
		fmt.Scanln(&op)
		if op == 1 {
			showProduct()
		} else if op == 2 {
			orderplace = placeOrder(curUser)
			curOrderList = append(curOrderList, orderplace)
			curUser.order = curOrderList
			orderlist = append(orderlist, curUser.order...)
			length = len(curUser.order)
		} else if op == 3 {
			modifyOder(curUser)
		} else if op == 4 {
			deleteOrder(curUser)
		} else if op == 5 {
			printProduct(curUser, length)
		} else if op == 6 {
			printOrder(curUser)
		} else {
			break
		}
	}

}
