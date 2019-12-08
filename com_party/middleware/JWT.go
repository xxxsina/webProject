package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"time"
	"webProject/com_party/helper"
	"webProject/com_party/models"
)

// 一些常量
var (
	TokenExpired     error  = errors.New(helper.CodeText(helper.Code10005))
	TokenNotValidYet error  = errors.New(helper.CodeText(helper.Code10006))
	TokenMalformed   error  = errors.New(helper.CodeText(helper.Code10007))
	TokenInvalid     error  = errors.New(helper.CodeText(helper.Code10008))
	SignKey          string = "d3ab7ede15cfdcfff930553da3889d01" //初始化key，可以通过SetSignKey设置
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID    int    `json:"id"`
	Mobile string `json:"mobile"`
	jwt.StandardClaims
}

// 登录时调用 生成token
func JWTXcreate(u *models.User, c *gin.Context) (string, error) {

	j := NewJWT()

	claims := CustomClaims{
		ID:    u.Id,
		Mobile: u.Mobile,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    GetSignKey(),                    //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		return "", err
	}

	return token, err
}

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": helper.Code10004,
				"msg":    helper.CodeText(helper.Code10004),
			})
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": helper.Code10005,
					"msg":    helper.CodeText(helper.Code10005),
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": helper.Code0,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
