package model

import "strings"

const COLLECTION_USER = "users"

type Address struct {
    Street   string `json:"street" bson:"street"`
    Block    string `json:"block" bson:"block"`
    District string `json:"district" bson:"district"`
    City     string `json:"city" bson:"city"`
    Country  string `json:"country" bson:"country"`
}

func (this *Address) ToString() string {
    param := []string{
        this.Street, this.Block, this.District, this.City, this.Country,
    }
    return strings.Join(param, ", ")
}

type User struct {
    Base `bson:",inline"`
    Email     string `json:"email" bson:"email"`
    Password  string `json:"-" bson:"password"`
    Avatar    string `json:"avatar" bson:"avatar"`
    FirstName string `json:"firstName" bson:"firstName"`
    LastName  string `json:"lastName" bson:"lastName"`
    Phone     string `json:"phone" bson:"phone"`
    FacebookId     string `json:"facebookId" bson:"facebookId"`
    Active    bool `json:"active" bson:"active"`
    Guest     bool `json:"guest" bson:"guest"`
    Roles     []string `json:"roles" bson:"roles"`
    Nothing   string `json:"roles" bson:"-"`
    Address   *Address `json:"address" bson:"address"`
}

func NewUser() *User {
    user := User{}
    user.Base.construct()
    return &user
}