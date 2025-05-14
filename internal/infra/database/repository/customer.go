package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICustomerRepository interface {
	Create(ctx context.Context, customer *model.Customer) error
	FindOne(ctx context.Context, id string) (*model.Customer, error)
	FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Customer, error)
}

type CustomerRepository struct {
	database *mongo.Database
}

func NewCustomerRepository(database *mongo.Database) ICustomerRepository {
	return &CustomerRepository{
		database: database,
	}
}

func (repository *CustomerRepository) Create(ctx context.Context, customer *model.Customer) error {

	client := repository.database.Client()
	repository.database = client.Database(repository.database.Name())
	collection := repository.database.Collection("customer")

	result, err := collection.InsertOne(context.Background(), &customer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("result: ", result)

	return nil
}

func (repository *CustomerRepository) FindOne(ctx context.Context, id string) (*model.Customer, error) {
	customer := &model.Customer{}

	/*result := repository.database.DB.WithContext(ctx).Where("customer_id = ?", id).First(&customer)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, nil
		}
		return nil, result.Error
	}*/

	return customer, nil
}

func (repository *CustomerRepository) FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Customer, error) {

	customer := &model.Customer{}

	//client := repository.database.Client()

	//repository.database = client.Database("tremligeiro")

	collection := repository.database.Collection("customer")

	err := collection.FindOne(ctx, bson.M{"documentnumber": documentNumber}).Decode(&customer)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("Erro ao buscar cliente: %v", err)
		return nil, err
	}

	return customer, nil

	//return nil, nil
}
