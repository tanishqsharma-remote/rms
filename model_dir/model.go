package model_dir

import (
	"database/sql"
	"time"
)

var JwtKey = []byte("MyKey")

type PersonList struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"role"`
	Address   sql.NullString `json:"address"`
	CreatedBy int            `json:"createdBy"`
}
type SubAdmin struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
type UserBySub struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   Point  `json:"address"`
	CreatedBy int    `json:"createdBy"`
}
type PersonByAdmin struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Address   Point  `json:"address"`
	CreatedBy int    `json:"createdBy"`
}
type NewPerson struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedBy int    `json:"createdBy"`
}
type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  Point  `json:"address"`
}
type Point struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
type Distance struct {
	RID  int     `json:"rId"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}
type Location struct {
	SNum    int   `json:"sNum"`
	PID     int   `json:"pId"`
	Address Point `json:"address"`
}

type Roles struct {
	SNum int    `json:"sNum"`
	PID  int    `json:"PId"`
	Role string `json:"role"`
}

type NewRestaurant struct {
	Name      string `json:"name"`
	Address   Point  `json:"address"`
	CreatedBy int    `json:"createdBy"`
}
type RestaurantList struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

//todo naming should be better no idea what is x and y

type Pointer struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}
type RestaurantLister struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Address Pointer `json:"address"`
}
type RestaurantId struct {
	RID int `json:"rId"`
}

type NewDish struct {
	RID       int    `json:"rId"`
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
type NewLocation struct {
	ID      int   `json:"id"`
	Address Point `json:"address"`
}
type NewRole struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"passWord"`
}
type Authentication struct {
	ID   string `json:"id"`
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
