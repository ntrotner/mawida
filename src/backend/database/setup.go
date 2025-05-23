package database

import (
	"context"
	"fmt"
	"template_backend/core/config"
	database_location "template_backend/database/paths/location"
	database_payments "template_backend/database/paths/payments"
	database_product "template_backend/database/paths/product"
	database_rent_contract "template_backend/database/paths/rent_contract"
	database_user "template_backend/database/paths/user"
	"time"

	_ "github.com/go-kivik/couchdb"
	"github.com/go-kivik/kivik"
	"github.com/rs/zerolog/log"
)

type databaseConnector struct {
	Client *kivik.Client
}

var Connection databaseConnector
var databases = make(map[string]*kivik.DB)

// createConnection creates a connection to the database
func createConnection(ctx context.Context) *kivik.Client {
	for {
		client, err := kivik.New("couch", config.GlobalConfig.Database.Host)

		if err != nil {
			log.Fatal().Msg("Database Connection Failed: " + err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		_, err = client.Session(ctx)
		if err != nil {
			fmt.Printf("Authentication error: %v\n", err)
			fmt.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		log.Info().Msg("Database Connection Successful")
		return client
	}
}

// createDatabase creates a database
func createDatabase(ctx context.Context, dbName string) {
	exists, err := Connection.Client.DBExists(ctx, dbName)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	if !exists {
		if err := Connection.Client.CreateDB(ctx, dbName); err != nil {
			log.Error().Msg(err.Error())
		}
	}
	databases[dbName] = Connection.Client.DB(ctx, dbName)
}

// Connect connects to the database
func Connect(ctx context.Context) {
	Connection = databaseConnector{
		Client: createConnection(ctx),
	}

	for _, value := range []string{
		database_user.USER_DB,
		database_location.LOCATION_DB,
		database_product.PRODUCT_DB,
		database_rent_contract.RENT_CONTRACT_DB,
		database_payments.PAYMENTS_DB,
	} {
		createDatabase(ctx, value)
	}

	go database_user.SetupUser(ctx, databases)
	go database_location.SetupLocation(ctx, databases)
	go database_payments.SetupPayments(ctx, databases)
	go database_product.SetupProduct(ctx, databases)
	go database_rent_contract.SetupRentContract(ctx, databases)
}
