/*
gorm mysql demo
*/
package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"`      // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               // 忽略本字段
	Description  string  `gorm:"type:varchar(256)"`
	CityId       uint
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Order struct {
	gorm.Model
	OrderNo string
}

type City struct {
	gorm.Model
	CityName string
}

func main() {
	//获取连接
	db, err := gorm.Open("mysql", "gxl:123456@/health?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("gorm.Open err:", err)
		panic("连接数据库失败")
	}
	defer db.Close()

	// 自动迁移模式
	//db.AutoMigrate(&Product{})

	// 创建
	//db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	//var product Product
	// 查询id为1的product
	//db.First(&product, 1)
	// 查询code为l1212的product
	//db.First(&product, "code = ?", "L1212")
	//log.Printf("Code:%s, Price:%d, CreatedAt:%s, UpdatedAt:%s", product.Code, product.Price, product.CreatedAt, product.UpdatedAt)

	// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	//db.Delete(&product)

	//自动迁移自动迁移模式将保持更新到最新。
	//警告：自动迁移仅仅会创建表，缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据。
	//db.AutoMigrate(&User{})

	//db.AutoMigrate(&City{})

	//db.AutoMigrate(&User{}, &Product{}, &Order{})

	// 创建表时添加表后缀
	//db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	// 检查模型`User`表是否存在
	//hasUseTable := db.HasTable(&User{})

	// 为模型`User`创建表
	//db.CreateTable(&User{})

	// 检查表`users`是否存在
	//db.HasTable("users")

	// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
	//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})

	// 删除模型`User`的表
	//db.DropTable(&User{})

	// 删除表`users`
	//db.DropTable("users")

	// 删除模型`User`的表和表`products`
	//db.DropTableIfExists(&User{}, "products")

	// 修改模型`User`的description列的数据类型为`text`
	//db.Model(&User{}).ModifyColumn("description", "text")

	// 删除模型`User`的description列
	//db.Model(&User{}).DropColumn("description")

	// 添加主键
	// 1st param : 外键字段
	// 2nd param : 外键表(字段)
	// 3rd param : ONDELETE
	// 4th param : ONUPDATE
	//db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")

	// 为`name`列添加索引`idx_user_name`
	//db.Model(&User{}).AddIndex("idx_user_name", "name")

	// 为`name`, `age`列添加索引`idx_user_name_age`
	//db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")

	// 添加唯一索引
	//db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")

	// 为多列添加唯一索引
	//db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")

	// 删除索引
	//db.Model(&User{}).RemoveIndex("idx_user_name")

}
