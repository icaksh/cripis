package models

import "github.com/google/uuid"

type JwtTokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type JwtAuthModel struct {
	AccessId  uuid.UUID
	RefreshId uuid.UUID
	UserId    uuid.UUID
	Email     string
	FirstName string
	LastName  string
	Role      int16
	Duration  int64
}

type JwtRtModel struct {
	AccessId  uuid.UUID
	RefreshId uuid.UUID
	UserId    uuid.UUID
	Email     string
	Role      int16
	Duration  int64
}
