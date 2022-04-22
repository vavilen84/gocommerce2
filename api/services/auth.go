package services

import (
	"api/helpers"
	"api/models"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gbrlsnchs/jwt/v3"
	_ "github.com/go-sql-driver/mysql"
)

type JWTAuthService struct {
	User models.User
}

type JWTPayload struct {
	jwt.Payload
	JWTInfoId int64 `json:"jwt_info_id"`
}

func (a *JWTAuthService) generateJWTSecret() string {
	return ""
}

func (a *JWTAuthService) insertJWTInfo(o orm.Ormer) (jwtInfo models.JWTInfo, err error) {
	now := helpers.GetNowUTCTimestamp()
	jwtInfo = models.JWTInfo{
		User:      &a.User,
		Secret:    a.generateJWTSecret(),
		CreatedAt: now,
	}
	err = models.InsertJWTInfo(o, &jwtInfo)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (a *JWTAuthService) generateJWT(jwtInfo models.JWTInfo) (token []byte, err error) {
	algorithm := jwt.NewHS256([]byte(jwtInfo.Secret))
	payload := JWTPayload{
		JWTInfoId: jwtInfo.Id,
	}
	token, err = jwt.Sign(payload, algorithm)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (a *JWTAuthService) CreateJWT(o orm.Ormer) (token []byte, err error) {
	jwtInfo, err := a.insertJWTInfo(o)
	if err != nil {
		logs.Error(err)
		return
	}
	token, err = a.generateJWT(jwtInfo)
	if err != nil {
		logs.Error(err)
	}
	return
}
