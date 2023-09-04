package model

type User struct {
	ID         string `bson:"_id,omitempty"`
	Username   string `bson:"username"`
	Email      string `bson:"email"`
	IsVerified bool   `bson:"isVerified"`
	PoolId     string `bson:"poolId,omitempty"`
}
