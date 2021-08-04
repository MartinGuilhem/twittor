package bd

import (
	"context"
	"time"

	"github.com/martinguilhem/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoRegistro is the last step with DB to insert user data */
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
