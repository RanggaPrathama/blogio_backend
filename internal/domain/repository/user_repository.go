package repository

import (
	"blogio/config"
	"blogio/internal/domain/entity"
	"blogio/internal/domain/repository/interfaces"
	"context"
	"fmt"
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


func (db *UserDatabase) UpdateUser(c context.Context, id string, user entity.User) (entity.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// users , _ := db.FindByID(c, id)	
	hexId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("Invalid ObjectID:", err)
		return user, err
	}

	_, err = db.collection.UpdateOne(ctx, bson.M{"_id" : hexId}, bson.M{"$set": bson.M{
		"username": user.USERNAME,
		"email": user.EMAIL,
		"password": user.PASSWORD,
	}})

	return user, err
}

func (db *UserDatabase) DeleteUser(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// hex_id, _ := primitive.ObjectIDFromHex(id)
	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid ObjectID:", err)
		return err
	}

	// users , _ := db.FindByID(c, id)
	_ , err = db.collection.DeleteOne(ctx, bson.M{"_id": hexID})
	return  err 
}