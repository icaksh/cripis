package database

import (
	"github.com/icaksh/cripis/app/queries"
)

type Queries struct {
	*queries.LogQueries
	*queries.UserQueries
	*queries.CivilQueries
	*queries.AnnouncementQueries
	*queries.TrademarkQueries
}

type RedisQueries struct {
	*queries.JWTQueries
}

func Connect() (*Queries, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		LogQueries:          &queries.LogQueries{DB: db},
		UserQueries:         &queries.UserQueries{DB: db, LogQueries: queries.LogQueries{DB: db}},
		CivilQueries:        &queries.CivilQueries{DB: db},
		AnnouncementQueries: &queries.AnnouncementQueries{DB: db},
		TrademarkQueries:    &queries.TrademarkQueries{DB: db, LogQueries: queries.LogQueries{DB: db}},
	}, nil
}

func RedisConnect() (*RedisQueries, error) {
	redis, err := RedisConnection()
	if err != nil {
		return nil, err
	}

	return &RedisQueries{
		JWTQueries: &queries.JWTQueries{Client: redis},
	}, nil

}
