package tokenModule

import (
	// "fmt"
	"errors"
	"strings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"golang/pkg/helpers"
)

type TokenService struct{
}

func NewTokenService() *TokenService{
	return &TokenService{}
}

type payload struct{
	MemberId string `json:"memberId"`
	Name string `json:"name"`
	TimeStamp int `json:"timeStamp"`
}

type header struct{
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

func (t *TokenService) Create(memberId string , name string)(string,error){
	header := header{
		Typ:"JWT",
		Alg: "HS256",
	}
	headerJson,err := json.Marshal(header)
	if err != nil {
		return "",err
    }

	payload := payload{
		MemberId:memberId,
		Name : name,
		TimeStamp : helpers.GetTimeStamp(),
	}
	
	payloadJson,err := json.Marshal(payload)
	if err != nil {
        return "",err
    }

	base64UrlHeaderJson := t.createUrlBase64(headerJson)
	base64UrlPayloadJson := t.createUrlBase64(payloadJson)

	secretKey,err := t.getSecretKey()
	if err!=nil{
		return "",err
	}
	signature := t.createSign(base64UrlHeaderJson, base64UrlPayloadJson ,secretKey)
	base64Signature := t.createUrlBase64(signature)
	return base64UrlHeaderJson + "." + base64UrlPayloadJson + "." + base64Signature , nil
}

func (t *TokenService) createBase64(item []byte)string{
	encodedData := base64.RawStdEncoding.EncodeToString(item)
	return encodedData
}

func (t *TokenService) createUrlBase64(item []byte)string{
	encodedData := base64.RawURLEncoding.EncodeToString(item)
	return encodedData
}

func (t *TokenService) decodeUrlBase64(item string)(string,error){
	decodeBytes,err := base64.RawURLEncoding.DecodeString(item)
	if err!=nil{
		return "",err
	}
	return string(decodeBytes),nil
}

func (t *TokenService) createSign(header string ,paylaod string , key string)[]byte {
	byteKey := []byte(key)
	h := hmac.New(sha256.New , byteKey)
	h.Write([]byte(header+"."+paylaod))
	return h.Sum(nil)
}

func (t *TokenService) getSecretKey()(string,error){
	key := helpers.GetEnvStr("jwt.key")
	if len(key) == 0{
		return "",errors.New("env_setting_error")
	}
	return key,nil
}

func (t *TokenService) getExpireTime()(int ,error){
	expireTime ,err := helpers.GetEnvInt("jwt.expireTime")
	if err!=nil{
		return 0,errors.New("jwtExpireTime missing")
	}
	return expireTime,nil
}

func (t *TokenService) IsValidJwt(jwtToken string) (bool,error) {
	splitedToken := strings.Split(jwtToken,".")
	if len(splitedToken)!=3{
		return false,nil
	}
	base64HeaderJson := splitedToken[0]
	base64PayloadJson := splitedToken[1]
	base64Signature := splitedToken[2]

	key,err := t.getSecretKey()
	if err!=nil{
		return false,err
	}

	sign :=t.createSign(base64HeaderJson , base64PayloadJson , key)
	correctSign := t.createUrlBase64(sign)
	if base64Signature != correctSign{
		return false,nil
	}
	return true,nil
}

func (t *TokenService) IsJwtInTime(token string)(bool,error){
	
	expireTime , err := t.getExpireTime()
	if err!=nil{
		return false,err
	}

	splitedToken := strings.Split(token,".")
	base64PayloadJson := splitedToken[1]
	payloadJson,err := t.decodeUrlBase64(base64PayloadJson)
	if err!=nil{
		return false,err
	}

	var payload payload
	err = json.Unmarshal([]byte(payloadJson), &payload)
	if err!=nil{
		return false,err
	}

	if helpers.GetTimeStamp() - payload.TimeStamp > expireTime*60{
		return false,nil
	}
	return true,nil
}