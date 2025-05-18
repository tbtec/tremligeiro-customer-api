package mongodb

import (
	"context"
	"fmt"
	"log/slog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConf struct {
	Url            string
	DbName         string
	CollectionName string
	User           string
	Pass           string
	Port           int
}

func New(conf MongoConf) (*mongo.Collection, error) {

	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", conf.User, conf.Pass, conf.Url, conf.Port, conf.DbName)
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/", conf.User, conf.Pass, conf.Url, conf.Port)

	slog.InfoContext(context.Background(), "Conectando ao MongoDB...")

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
		return nil, err
	}

	db := client.Database(conf.DbName)

	slog.InfoContext(context.Background(), "✅ Conexão com o MongoDB realizada com sucesso")

	return db.Collection(conf.CollectionName), nil
}

func Migrate(conf MongoConf) error {
	//dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", conf.User, conf.Pass, conf.Url, conf.Port, conf.DbName)
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%d/", conf.User, conf.Pass, conf.Url, conf.Port)

	slog.InfoContext(context.Background(), "Initializing migrations...")

	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
		return err
	}

	database := client.Database(conf.DbName)

	database.Collection(conf.CollectionName)

	slog.InfoContext(context.Background(), "Finished migrations")

	return nil
}
