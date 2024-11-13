package repository

import (
	"blogio/config"
	"blogio/internal/domain/entity"
	"blogio/internal/domain/repository/interfaces"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserDatabase struct {
	collection *mongo.Collection
}

func NewUserRepository() interfaces.UserRepository{
	return &UserDatabase{
		collection: config.GetCollection(config.Connection, "user"),
	}
}


func (db *UserDatabase) FindAll(c context.Context) ([]entity.User, error) {

	ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer cancel()
	
	var users []entity.User
	
	cursor, err := db.collection.Find(ctx, bson.D{})
	
	cursor.All(ctx, &users)

	// fmt.Print(users)
	
	// defer cursor.Close(ctx)

	return users, err
}


func (db *UserDatabase) FindByID(c context.Context, id string) (entity.User, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)
	defer cancel()

	var users entity.User

	hex_id, _ := primitive.ObjectIDFromHex(id)

	err := db.collection.FindOne(ctx, bson.M{"_id" : hex_id}).Decode(&users)
	
	return users, err
}	

func (db *UserDatabase) FindByEmail(c context.Context, email string) (entity.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)
	defer cancel()

	var user entity.User

     err := db.collection.FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&user)

	return user, err

}

func (db *UserDatabase) CreateUser(c context.Context, user entity.User) (entity.User, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10 *time.Second)
	defer cancel()

	// var user entity.User

	newUser := entity.User{
		ID: primitive.NewObjectID(),
		USERNAME: user.USERNAME,
		EMAIL: user.EMAIL,
		PASSWORD: user.PASSWORD,
	}

	_ , err := db.collection.InsertOne(ctx, newUser)
	
	return newUser, err
}


