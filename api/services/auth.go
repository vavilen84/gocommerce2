package services

import (
	"api/models"
	"encoding/base64"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gbrlsnchs/jwt/v3"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
	"time"
)

type JWTAuthService struct {
	User models.User
}

type JWTPayload struct {
	jwt.Payload
	JWTInfoId int64 `json:"jwt_info_id"`
}

func (a *JWTAuthService) insertJWTInfo(o orm.Ormer) (jwtInfo models.JWTInfo, err error) {
	jwtInfo = models.JWTInfo{
		User: &a.User,
	}
	err = models.InsertJWTInfo(o, &jwtInfo)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (a *JWTAuthService) generateJWT(jwtInfo models.JWTInfo) (token []byte, err error) {
	algorithm := jwt.NewHS256([]byte(jwtInfo.Secret))
	expirationTime := time.Unix(int64(jwtInfo.ExpiresAt), 0)
	payload := JWTPayload{
		Payload: jwt.Payload{
			ExpirationTime: jwt.NumericDate(expirationTime),
			IssuedAt:       jwt.NumericDate(time.Now()),
		},
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

func (a *JWTAuthService) parseJWTPayload(token []byte) (jwtPayload JWTPayload, err error) {
	re, err := regexp.Compile(`(.*)\.(?P<payload>.*)\.(.*)`)
	if err != nil {
		logs.Error(err)
		return
	}
	matches := re.FindStringSubmatch(string(token))
	i := re.SubexpIndex("payload")
	payloadData, err := base64.StdEncoding.DecodeString(matches[i])
	if err != nil {
		logs.Error(err)
	}
	err = json.Unmarshal(payloadData, &jwtPayload)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (a *JWTAuthService) VerifyJWT(o orm.Ormer, token []byte) (isValid bool, err error) {
	payload, err := a.parseJWTPayload(token)
	if err != nil {
		logs.Error(err)
		return
	}
	jwtInfo, err := models.FindJWTInfoById(o, payload.JWTInfoId)
	if err != nil {
		logs.Error(err)
		return
	}
	algorithm := jwt.NewHS256([]byte(jwtInfo.Secret))
	_, err = jwt.Verify(token, algorithm, &payload)
	if err != nil {
		logs.Error(err)
		return
	}
	isValid = true
	return
}
