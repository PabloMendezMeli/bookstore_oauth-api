package db

import (
	"github.com/PabloMendezMeli/bookstore_oauth-api/src/clients/cassandra"
	"github.com/PabloMendezMeli/bookstore_oauth-api/src/domain/access_token"
	"github.com/PabloMendezMeli/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct{}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//TODO: logic to get access_token from cassandra db
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
