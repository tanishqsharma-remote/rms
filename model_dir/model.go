package model_dir

import (
	"database/sql"
	"time"
)

var JwtKey = []byte("MyKey")

/*type Person struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"passWord"`
	CreatedBy int    `json:"createdBy"`
}
type PersonDetails struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"passWord"`
	Role      string `json:"role"`
	Address   Point  `json:"address"`
	CreatedBy int    `json:"createdBy"`
}*/

type PersonList struct {
	Id        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"passWord"`
	Role      string         `json:"role"`
	Address   sql.NullString `json:"address"`
	CreatedBy int            `json:"createdBy"`
}
type SubAdmin struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type UserBySub struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"passWord"`
	Address   Point  `json:"address"`
	CreatedBy int    `json:"createdBy"`
}
type PersonByAdmin struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"passWord"`
	Role      string `json:"role"`
	Address   Point  `json:"address"`
	CreatedBy int    `json:"createdBy"`
}
type NewPerson struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"passWord"`
	CreatedBy int    `json:"createdBy"`
}
type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"passWord"`
	Address  Point  `json:"address"`
}
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type Distance struct {
	RId int     `json:"rId"`
	X   float64 `json:"x"`
	Y   float64 `json:"y"`
}
type Location struct {
	SNum    int   `json:"sNum"`
	PId     int   `json:"pId"`
	Address Point `json:"address"`
}

type Roles struct {
	SNum int    `json:"sNum"`
	PId  int    `json:"PId"`
	Role string `json:"role"`
}

/*type Restaurant struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Address   Point  `json:"address"`
	CreatedBy int    `json:"createdBy"`
}*/

type NewRestaurant struct {
	Name      string `json:"name"`
	Address   Point  `json:"address"`
	CreatedBy int    `json:"createdBy"`
}
type RestaurantList struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type Pointer struct {
	X string `json:"x"`
	Y string `json:"y"`
}
type RestaurantLister struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Address Pointer `json:"address"`
}
type RestaurantId struct {
	RId int `json:"rId"`
}

/*type Dish struct {
	SNum      int    `json:"sNum"`
	RId       int    `json:"rId"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	CreatedBy int    `json:"createdBy"`
}*/

type NewDish struct {
	RId       int    `json:"rId"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	CreatedBy int    `json:"createdBy"`
}
type DishByRestaurant struct {
	RestaurantName string `json:"restaurantName"`
	DishName       string `json:"dishName"`
	Price          int    `json:"price"`
}
type DishList struct {
	SNum     int    `json:"sNum"`
	DishName string `json:"dishName"`
	Price    int    `json:"price"`
}
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"passWord"`
}
type Authentication struct {
	Id   string `json:"id"`
	Role string `json:"role"`
}
type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type Session struct {
	Email  string    `json:"email"`
	Expiry time.Time `json:"expiry"`
}
