package queries

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"github.com/icaksh/cripis/app/models"
	"github.com/icaksh/cripis/app/utils"
	"strconv"
	"time"
)

type JWTQueries struct {
	*redis.Client
}

func (q *JWTQueries) CreateAuth(userId uuid.UUID, td *models.JwtTokenDetails) error {
	fmt.Println(td)
	at := time.Until(time.Unix(td.AtExpires, 0))
	rt := time.Until(time.Unix(td.RtExpires, 0))
	fmt.Println(rt)
	errAccess := q.Set(td.AccessUuid, userId, at).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := q.Set(td.RefreshUuid, userId, rt).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (q *JWTQueries) FetchAuth(authD *utils.JWTAccess) (uint64, error) {
	userid, err := q.Get(authD.AccessUuid.String()).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func (q *JWTQueries) DeleteAuth(uuid uuid.UUID) (int64, error) {
	fmt.Println("awokaokawo1")
	deleted, err := q.Del(uuid.String()).Result()
	fmt.Println("awokaokawo2")
	if err != nil {
		fmt.Println("awokaokawo3")
		return 0, err
	}
	return deleted, nil
}
