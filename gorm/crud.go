package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int    `gorm:"AUTO_INCREMENT"` // 自增
	Role     string
}

//您可以在gorm tag中定义默认值，然后插入SQL将忽略具有默认值的这些字段，
//并且其值为空，并且在将记录插入数据库后，gorm将从数据库加载这些字段的值。
type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//如果要在BeforeCreate回调中设置主字段的值，可以使用scope.SetColumn，例如：
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.Nil)
	return nil
}

func main() {
	//获取连接
	db, err := gorm.Open("mysql", "gxl:123456@/health?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("gorm.Open err:", err)
		panic("连接数据库失败")
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})

	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

	//主键为空返回true
	//saveResult := db.NewRecord(user)
	//log.Fatal(saveResult)

	//db.Create(&user)

	//创建user后返回false
	//saveResult := db.NewRecord(user)
	//log.Fatal(saveResult)

	//var animal = Animal{Age: 99, Name: ""}
	//db.Create(&animal)

	//var product Product
	// 为Instert语句添加扩展SQL选项
	//db.Set("gorm:insert_option", "ON CONFLICT").Create(&product)

	//var user User
	//查询
	// 获取第一条记录，按主键排序
	//db.First(&user)
	//// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取最后一条记录，按主键排序
	//db.Last(&user)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 获取所有记录
	//db.Find(&users)
	//// SELECT * FROM users;

	// 使用主键获取记录
	//db.First(&user, 1)
	//// SELECT * FROM users WHERE id = 1;

	//log.Fatal(user)

	var users User
	//Where查询条件 (简单SQL)
	// 获取第一个匹配记录
	//db.Where("name = ?", "jinzhu").First(&user)
	//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

	// 获取所有匹配记录
	//db.Where("name = ?", "jinzhu").Find(&users)
	//// SELECT * FROM users WHERE name = 'jinzhu';

	//db.Where("name <> ?", "jinzhu").Find(&users)

	// IN
	//db.Where("name in (?)", []string{"jinzhu", "jinzhu 2"}).Find(&users)

	// LIKE
	//db.Where("name LIKE ?", "%jin%").Find(&users)

	// AND
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)

	//lastWeek := time.Date(2020, 2, 24, 0, 0, 0, 0, &time.Location{})
	// Time
	//db.Where("updated_at > ?", lastWeek).Find(&users)

	//today := time.Now()
	//db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)

	//Where查询条件 (Struct & Map)
	//注意：当使用struct查询时，GORM将只查询那些具有值的字段

	// Struct
	//db.Where(&User{Name: "jinzhu", Age: 18}).First(&user)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

	// Map
	//db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键的Slice
	db.Where([]int64{20, 21, 22}).Find(&users)
	//// SELECT * FROM users WHERE id IN (20, 21, 22);

	log.Fatal(users)
}
