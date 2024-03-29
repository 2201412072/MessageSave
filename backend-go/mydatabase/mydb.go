package mydatabase

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var myDB *gorm.DB
var err error

type Public_keys struct {
	Username   string `gorm:"primaryKey;size:100"`
	Public_key []byte `gorm:"type:longblob"`
}
type Agree struct {
	Username string    `gorm:"primaryKey;column:username;foreignKey:username;references:public_keys:username"`
	Mytime   time.Time `gorm:"primaryKey;column:mytime"`
	Operator string
}

type Password struct {
	Application string `gorm:"primaryKey;column:application"`
	Username    string `gorm:"primaryKey;column:username;foreignKey:username;references:public_keys:username"`
	Saved_key   []byte `gorm:"type:longblob"`
	Single_key  string `gorm:"size:100"`
}

type Important_message struct {
	Keyword   string `gorm:"primaryKey;column:keyword;size:200"`
	Username  string `gorm:"primaryKey;column:username;foreignKey:username;references:public_keys:username"`
	Saved_key []byte `gorm:"type:longblob"`
}

func Get_dict_value(dict map[string]interface{}, key string) interface{} {
	return dict[key]
}

func Download_db() *gorm.DB {
	return myDB
}

func Db_load() (*gorm.DB, int) {
	user := "message_save"
	password := "message_save"
	host := "localhost"
	port := "3306"
	dbname := "MessageSaveDB"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, dbname, charset)

	// _ = dsn
	// // 建立数据库连接
	// sqlDB, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	// if err != nil {
	// 	return nil, 0
	// }
	// // defer MyDB.Close()
	// // 测试连接是否成功
	// err = sqlDB.Ping()
	// if err != nil {
	// 	return nil, 0
	// }
	// myDB, err = gorm.Open(mysql.New(mysql.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})
	// fmt.Println("数据库连接成功", sqlDB)

	// 不能加冒号，一加冒号就会在函数内部生成一个局部的myDB返回，导致下面显示的myDB为nil
	// https://blog.csdn.net/m0_46251547/article/details/123294581
	myDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("myDB open failed!", err)
		return nil, 0
	}

	// 数据库初始化表（默认数据库已建立）
	flag := DB_init_table()
	if flag != 1 {
		return nil, 0
	}
	// fmt.Println("MyDB", myDB)  // &{0xc0001290e0 invalid transaction 0 0xc000450540 1}
	return myDB, 1
}

func DB_init_table() int {
	err := myDB.AutoMigrate(&Public_keys{})
	if err != nil {
		return 0
	}
	err = myDB.AutoMigrate(&Agree{})
	if err != nil {
		return 0
	}
	err = myDB.AutoMigrate(&Password{})
	if err != nil {
		return 0
	}
	err = myDB.AutoMigrate(&Important_message{})
	if err != nil {
		return 0
	}
	return 1
}

// 以下成功返回1，失败返回其他
func Get_all_public_keys() ([]Public_keys, int) {
	var temps []Public_keys
	myDB.Find(&temps)
	return temps, 1
}

func Get_single_public_key(username string) (Public_keys, int) {
	var temps []Public_keys
	var err int
	var temp Public_keys
	myDB.Where("username = ? ", username).Find(&temps)
	if len(temps) > 1 || len(temps) == 0 {
		err = 0
		temp = Public_keys{}
	} else {
		temp = temps[0]
		err = 1
	}
	return temp, err
}

func Add_public_key(username string, public_key []byte) int {
	temp := Public_keys{Username: username, Public_key: public_key}
	result := myDB.Create(&temp)
	//myDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Set("gorm:insert_option", "ON DUPLICATE KEY UPDATE").Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_public_key(username string) int {
	result := myDB.Delete(&Public_keys{}, username)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Change_public_key(username string, public_key []byte) int {
	result := myDB.Model(&Public_keys{}).Where("username=?", username).Update("public_key", public_key)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Get_all_agree() ([]Agree, int) {
	var temps []Agree
	result := myDB.Find(&temps)
	if result.Error != nil {
		return temps, 0
	} else {
		return temps, 1
	}
}

func Get_single_user_agree(username string) ([]Agree, int) {
	var temps []Agree
	myDB.Where("username = ? ", username).Find(&temps)
	return temps, 0
}

func Get_single_user_time_agree(username string, mytime time.Time) (Agree, int) {
	var temps []Agree
	var err int
	var temp Agree
	myDB.Where("username = ? AND mytime=?", username, mytime).Find(&temps)
	if len(temps) > 1 {
		err = 0
		temp = Agree{}
	} else {
		temp = temps[0]
		err = 1
	}
	return temp, err
}

func Add_agree(username string, mytime time.Time, operator string) int {
	temp := Agree{Username: username, Mytime: mytime, Operator: operator}
	result := myDB.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_agree(username string, mytime time.Time) int {
	result := myDB.Where("username=? AND mytime=?", username, mytime).Delete(&Agree{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func Clean_create_agree(mytime time.Time) int {
	//清除agree表中小于mytime的记录
	result := myDB.Where(" mytime<?", mytime).Delete(&Agree{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func Get_password() ([]Password, int) {
	var temps []Password
	result := myDB.Find(&temps)
	if result.Error != nil {
		return temps, 0
	} else {
		return temps, 1
	}
}

func Get_application2user_password(application string) ([]Password, int) {
	var temps []Password
	result := myDB.Where("application=?", application).Find(&temps)
	if result.Error != nil {
		return temps, 0
	} else {
		return temps, 1
	}
}

func Get_application_name2key_password(application string, username string) (Password, int) {
	var temps []Password
	result := myDB.Where("application=? AND username=?", application, username).Find(&temps)
	if result.Error != nil {
		return Password{}, 0
	} else {
		if len(temps) != 1 {
			return Password{}, 0
		} else {
			return temps[0], 1
		}
	}
}

func Add_password(application string, username string, saved_key []byte, single_key string) int {
	temp := Password{Username: username, Application: application, Saved_key: saved_key, Single_key: single_key}
	result := myDB.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_single_password(application string, username string) int {
	result := myDB.Where("username=? AND application=?", username, application).Delete(&Password{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_username_password(username string) int {
	result := myDB.Where("username=? ", username).Delete(&Password{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func Get_important() ([]Important_message, int) {
	var temps []Important_message
	result := myDB.Find(&temps)
	if result.Error != nil {
		return temps, 0
	} else {
		return temps, 1
	}
}

func Get_keyword_important() ([]Important_message, int) {
	var temps []Important_message
	myDB.Find(&temps)
	return temps, 1
}

func Get_keyword2user_important(keyword string) ([]Important_message, int) {
	var temps []Important_message
	myDB.Where("keyword=?", keyword).Find(&temps)
	return temps, 1
}

func Get_user2keyword_important(username string) ([]Important_message, int) {
	var temps []Important_message
	myDB.Where("username=?", username).Find(&temps)
	return temps, 1
}

func Get_keyword_name2key_important(keyword string, username string) (Important_message, int) {
	var temps []Important_message
	result := myDB.Where("keyword=? AND username=?", keyword, username).Find(&temps)
	if result.Error != nil {
		return Important_message{}, 0
	} else {
		if len(temps) > 1 {
			return Important_message{}, 0
		} else {
			return temps[0], 1
		}
	}
}

func Add_important(keyword string, username string, saved_key []byte) int {
	temp := Important_message{Username: username, Keyword: keyword, Saved_key: saved_key}
	result := myDB.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_single_important(keyword string, username string) int {
	result := myDB.Where("username=? AND keyword=?", username, keyword).Delete(&Important_message{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delte_username_important(username string) int {
	result := myDB.Where("username=?", username).Delete(&Important_message{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_keyword_important(keyword string) int {
	result := myDB.Where("keyword=?", keyword).Delete(&Important_message{})
	if result.Error != nil {
		return 0
	}
	return 1
}
