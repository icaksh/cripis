package database

import "github.com/icaksh/cripis/app/queries"


type Queries struct {
	*queries.UserQueries
}

func Connect() (*Queries, error){
		db, err := PostgreSQLConnection()
		if err != nil {
			return nil, err
		}
	
		return &Queries{
			UserQueries: &queries.UserQueries{DB: db},
		}, nil
	}