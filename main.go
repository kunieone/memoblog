package main

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

var j = jwt.New(jwt.Config{
	// Extract by "token" url parameter.
	// Extractor: jwt.FromParameter("token"),
	Extractor: jwt.FromAuthHeader,

	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("My Secret"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func main() {
	app := iris.Default()

	app.Get("/", getToken)
	app.Get("/secured", j.Serve, func(ctx iris.Context) {
		ctx.JSON("hello")
	})
	app.Run(iris.Addr(":3000"))
}

func getToken(ctx iris.Context) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":  "bar",
		"开始时间": "2022/5/19",
	})
	tokenString, _ := token.SignedString([]byte("My Secret"))
	ctx.JSON(tokenString)
}
