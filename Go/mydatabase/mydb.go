package mydatabase

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var myDB *gorm.DB

type public_keys struct {
	username   string `gorm:"primaryKey;size:100"`
	public_key []byte `gorm:"type:longblob"`
}
type agree struct {
	username string    `gorm:"primaryKey;column:username;foreignKey:username;references:public_keys:username"`
	mytime   time.Time `gorm:"primaryKey;column:mytime"`
	operator string
}

type Password struct {
	application string `gorm:"primaryKey;column:application"`
	username    string `gorm:"primaryKey;column:username;foreignKey:username;references:public_keys:username"`
	Saved_key   []byte `gorm:"type:longblob"`
	single_key  string `gorm:"size:100"`
}

type important_message struct {
	keyword   string `gorm:"primaryKey;column:keyword;size:200"`
	username  string `gorm:"primaryKey;column:username;foreignKey:username;references:public_keys:username"`
	saved_key []byte `gorm:"type:longblob"`
}

func Get_dict_value(dict map[string]interface{}, key string) interface{} {
	return dict[key]
}

func Download_db() *gorm.DB {
	return myDB
}

func Db_load(db_path string) (*gorm.DB, int) {
	user := "user"
	password := "user"
	host := "localhost"
	port := "3306"
	dbname := "MessageSave"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, dbname, charset)
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := os.Stat(db_path)
	if os.IsNotExist(err) {
		myDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			//panic(err)
			return nil, 0
		}
	} else {
		sqlDB, err := sql.Open("mysql", dsn)
		if err != nil {
			//panic(err)
			return nil, 0
		}
		myDB, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
	}
	flag := DB_init_table()
	if flag != 1 {
		return nil, 0
	}
	return myDB, 1
}

func DB_init_table() int {
	err := myDB.AutoMigrate(&public_keys{})
	if err != nil {
		return 0
	}
	err = myDB.AutoMigrate(&agree{})
	if err != nil {
		return 0
	}
	err = myDB.AutoMigrate(&Password{})
	if err != nil {
		return 0
	}
	err = myDB.AutoMigrate(&important_message{})
	if err != nil {
		return 0
	}
	return 1
}

// 以下成功返回1，失败返回其他
func Get_all_public_keys() ([]public_keys, int) {
	var temps []public_keys
	myDB.Find(&temps)
	return temps, 1
}

func Get_single_public_key(username string) (public_keys, int) {
	var temps []public_keys
	var err int
	var temp public_keys
	myDB.Where("username = ? ", username).Find(&temps)
	if len(temps) > 1 {
		err = 0
		temp = public_keys{}
	} else {
		temp = temps[0]
		err = 1
	}
	return temp, err
}

func Add_public_key(username string, public_key []byte) int {
	temp := public_keys{username: username, public_key: public_key}
	result := myDB.Create(&temp)
	//myDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Set("gorm:insert_option", "ON DUPLICATE KEY UPDATE").Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_public_key(username string) int {
	result := myDB.Delete(&public_keys{}, username)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Change_public_key(username string, public_key []byte) int {
	result := myDB.Model(&public_keys{}).Where("username=?", username).Update("public_key", public_key)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Get_all_agree() ([]agree, int) {
	var temps []agree
	result := myDB.Find(&temps)
	if result.Error != nil {
		return temps, 0
	} else {
		return temps, 1
	}
}

func Get_single_user_agree(username string) ([]agree, int) {
	var temps []agree
	myDB.Where("username = ? ", username).Find(&temps)
	return temps, 0
}

func Get_single_user_time_agree(username string, mytime time.Time) (agree, int) {
	var temps []agree
	var err int
	var temp agree
	myDB.Where("username = ? AND mytime=?", username, mytime).Find(&temps)
	if len(temps) > 1 {
		err = 0
		temp = agree{}
	} else {
		temp = temps[0]
		err = 1
	}
	return temp, err
}

func Add_agree(username string, mytime time.Time, operator string) int {
	temp := agree{username: username, mytime: mytime, operator: operator}
	result := myDB.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_agree(username string, mytime time.Time) int {
	result := myDB.Where("username=? AND mytime=?", username, mytime).Delete(&agree{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func Clean_create_agree(mytime time.Time) int {
	//清除agree表中小于mytime的记录
	result := myDB.Where(" mytime<?", mytime).Delete(&agree{})
	if result.Error != nil {
		return 0
	}
	return 1
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
		if len(temps) > 1 {
			return Password{}, 0
		} else {
			return temps[0], 1
		}
	}
}

func Add_password(application string, username string, saved_key []byte, single_key string) int {
	temp := Password{username: username, application: application, Saved_key: saved_key, single_key: single_key}
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

func Get_keyword_important() ([]important_message, int) {
	var temps []important_message
	myDB.Find(&temps)
	return temps, 1
}

func Get_keyword2user_important(keyword string) ([]important_message, int) {
	var temps []important_message
	myDB.Where("keyword=?", keyword).Find(&temps)
	return temps, 1
}

func Get_user2keyword_important(username string) ([]important_message, int) {
	var temps []important_message
	myDB.Where("username=?", username).Find(&temps)
	return temps, 1
}

func Get_keyword_name2key_important(keyword string, username string) (important_message, int) {
	var temps []important_message
	result := myDB.Where("keyword=? AND username=?", keyword, username).Find(&temps)
	if result.Error != nil {
		return important_message{}, 0
	} else {
		if len(temps) > 1 {
			return important_message{}, 0
		} else {
			return temps[0], 1
		}
	}
}

func Add_important(keyword string, username string, saved_key []byte) int {
	temp := important_message{username: username, keyword: keyword, saved_key: saved_key}
	result := myDB.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_single_important(keyword string, username string) int {
	result := myDB.Where("username=? AND keyword=?", username, keyword).Delete(&important_message{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delte_username_important(username string) int {
	result := myDB.Where("username=?", username).Delete(&important_message{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func Delete_keyword_important(keyword string) int {
	result := myDB.Where("keyword=?", keyword).Delete(&important_message{})
	if result.Error != nil {
		return 0
	}
	return 1
}
