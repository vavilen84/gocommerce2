package services

import (
	"api/helpers"
	"api/models"
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_JWTAuthFlow(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// create user
	u := usersFixtures[user1key]
	err := models.InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	assert.Empty(t, u.ValidationErrors)

	// create jwt
	jwtAuthService := JWTAuthService{
		User: u,
	}
	token, err := jwtAuthService.CreateJWT(o)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// parse jwt payload
	payload, err := jwtAuthService.parseJWTPayload(token)
	assert.Nil(t, err)
	assert.NotEmpty(t, payload.JWTInfoId)

	// find jwt info
	jwtInfo, err := models.FindJWTInfoById(o, payload.JWTInfoId)
	assert.Nil(t, err)
	assert.NotEmpty(t, jwtInfo.Id)
	assert.Equal(t, u.Id, jwtInfo.Id)
	assert.Equal(t, int(payload.Payload.ExpirationTime.UTC().Unix()), jwtInfo.ExpiresAt)

	// verify jwt
	isValid, err := jwtAuthService.VerifyJWT(o, token)
	assert.Nil(t, err)
	assert.True(t, isValid)
}

func Test_OverridedJWTTokenExpirationTime(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// create user
	u := usersFixtures[user1key]
	err := models.InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	assert.Empty(t, u.ValidationErrors)

	// create jwt
	now := helpers.GetNowUTCTimestamp()
	jwtAuthService := JWTAuthService{
		User:           u,
		ExpirationTime: now + 60*60,
	}
	token, err := jwtAuthService.CreateJWT(o)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// parse jwt payload
	payload, err := jwtAuthService.parseJWTPayload(token)
	assert.Nil(t, err)
	assert.NotEmpty(t, payload.JWTInfoId)

	// find jwt info
	jwtInfo, err := models.FindJWTInfoById(o, payload.JWTInfoId)
	assert.Nil(t, err)
	assert.NotEmpty(t, jwtInfo.Id)
	assert.Equal(t, u.Id, jwtInfo.Id)
	assert.Equal(t, int(payload.Payload.ExpirationTime.UTC().Unix()), jwtInfo.ExpiresAt)

	// verify jwt
	isValid, err := jwtAuthService.VerifyJWT(o, token)
	assert.Nil(t, err)
	assert.True(t, isValid)
}

func Test_ExpiredJWTToken(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// create user
	u := usersFixtures[user1key]
	err := models.InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	assert.Empty(t, u.ValidationErrors)

	// create jwt
	past := int(time.Now().UTC().Unix()) - 60*60
	jwtAuthService := JWTAuthService{
		User:           u,
		ExpirationTime: past,
	}
	token, err := jwtAuthService.CreateJWT(o)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	// parse jwt payload
	payload, err := jwtAuthService.parseJWTPayload(token)
	assert.Nil(t, err)
	assert.NotEmpty(t, payload.JWTInfoId)

	// find jwt info
	jwtInfo, err := models.FindJWTInfoById(o, payload.JWTInfoId)
	assert.Nil(t, err)
	assert.NotEmpty(t, jwtInfo.Id)
	assert.Equal(t, u.Id, jwtInfo.Id)
	assert.Equal(t, int(payload.Payload.ExpirationTime.UTC().Unix()), jwtInfo.ExpiresAt)

	// verify jwt
	isValid, err := jwtAuthService.VerifyJWT(o, token)
	assert.Nil(t, err)
	assert.False(t, isValid)
}
