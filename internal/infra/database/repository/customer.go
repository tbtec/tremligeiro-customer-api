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
	database *mongo.Collection
}

func NewCustomerRepository(database *mongo.Collection) ICustomerRepository {
	return &CustomerRepository{
		database: database,
	}
}

func (repository *CustomerRepository) Create(ctx context.Context, customer *model.Customer) error {

	result, err := repository.database.InsertOne(context.Background(), &customer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("result: ", result)

	return nil
}

func (repository *CustomerRepository) FindOne(ctx context.Context, id string) (*model.Customer, error) {
	customer := &model.Customer{}

	err := repository.database.FindOne(ctx, bson.M{"id": id}).Decode(&customer)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("record not found")
			//return nil, nil
		}
		log.Printf("Erro ao buscar cliente: %v", err)
		return nil, err
	}

	return customer, nil
}

func (repository *CustomerRepository) FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Customer, error) {

	customer := &model.Customer{}

	err := repository.database.FindOne(ctx, bson.M{"documentnumber": documentNumber}).Decode(&customer)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("record not found")
			//return nil, nil
		}
		log.Printf("Erro ao buscar cliente: %v", err)
		return nil, err
	}

	return customer, nil

	//return nil, nil
}
