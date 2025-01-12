package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"firstName" validate:"required,min=2,max=30"`
	Last_Name       *string            `json:"lastName" validate:"required,min=2,max=30"`
	Email           *string            `json:"email" validate:"email,required"`
	Password        *string            `json:"password" validate:"required,min=6"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refreshToken"`
	User_ID         string             `json:"userId"`
	UserCart        []ProductUser      `json:"userCart" bson:"userCart"`
	Address_Details []Address          `json:"addressDetails" bson:"addressDetails"`
	Order_Status    []Order            `json:"orderStatus" bson:"orderStatus"`
	Created_At      time.Time          `json:"createdAt"`
	Updated_At      time.Time          `json:"updatedAt"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"productName"`
	Price        *int               `json:"price"`
	Image        *string            `json:"image"`
	Rating       *uint              `json:"rating"`
	Created_At   *time.Time         `json:"createdAt"`
	Updated_At   *time.Time         `json:"updatedAt"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"productName" bson:"productName"`
	Price        *uint64            `json:"price" bson:"price"`
	Ratings      *uint8             `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house" bson:"house"`
	Street     *string            `json:"street" bson:"street"`
	City       *string            `json:"city" bson:"city"`
	PinCode    *string            `json:"pinCode" bson:"pinCode"`
}

type Order struct {
	Order_ID   primitive.ObjectID `bson:"_id"`
	Order_Cart []ProductUser      `json:"orderList" bson:"orderList"`
	Ordered_At time.Time          `json:"orderedAt" bson:"orderedAt"`
	Price      int                `json:"price" bson:"price"`
	Discount   *int               `json:"discount" bson:"discount"`
	Payment    Payment            `json:"payment" bson:"payment"`
}

type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
