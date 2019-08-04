package dao

import (
	"PowerShop/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


func Init() *gorm.DB{
	db,err:=gorm.Open("mysql","root:roxi@(127.0.0.1:3306)/powershop?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("?")
	}
	return db
}

func Create(obj interface{}){
	db:=Init()
	db.Create(obj)
	fmt.Println(&obj,"=-=")
	defer db.Close()
}

func Update(shop string,good string,number int)  {
	db:=Init()
	defer db.Close()
	store:=&model.Store{ShopId:shop,GoodId:good}
	db.First(&store)
	fmt.Println(store,"-=-")
	db.Model(&store).Update("number",store.Number-number)
}