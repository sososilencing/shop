package buy

import (
	"PowerShop/model"
	"PowerShop/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)
func getHttp() (*gin.Engine){
	app:=gin.Default()
	return app
}


func Do() {
	app:=getHttp()

	app.GET("/buy", func(context *gin.Context) {
		user:=context.Query("user")
		good:=context.Query("good")
		shop:=context.Query("shop")
		number,err:=strconv.Atoi(context.Query("number"))
		if err != nil {
			number=0
			context.JSON(200,gin.H{
				"state":"Fail",
			})
		}else {

			body := &model.Order{UserId: user, GoodId: good, ShopId: shop, Number: number, Time: time.Unix(time.Now().Unix(), 0).String()}
			
			service.Produce(*body)

			context.JSON(http.StatusOK,gin.H{
				"message":gin.H{
					"user-id":user,
					"good-id":good,
					"shop-id":shop,
					"number":number,
				},
				"status":"OK",
			})

		}
	})
	app.GET("/login",inquire)
	app.Run(":8080")
}

func inquire(context *gin.Context) {
	context.Query("")
}

