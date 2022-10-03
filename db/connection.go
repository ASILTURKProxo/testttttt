package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"e-vet/globals"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CstmTime time.Time

type Credentials struct {
	host    string
	port    string
	user    string
	pass    string
	sslMode string
	dBName  string
	schemas []string
}

type Database struct {
	Credentials
	DB *gorm.DB
}

var (
	Dbs = map[string]*Database{
		"main": {Credentials: Credentials{"localhost", "5432", "postgres", "1", "disable", "ee-vet", []string{"public"}}},
	}
	DBConn *Database
)

func (db *Database) Connect() {
	if DBConn != nil {
		dbCon, _ := db.DB.DB()
		dbCon.Close()
	}

	dsn := "host=" + db.Credentials.host + " port=" + db.Credentials.port + " user=" + db.Credentials.user + " password=" + db.Credentials.pass + " dbname=" + db.Credentials.dBName + " sslmode=" + db.Credentials.sslMode + " TimeZone=Europe/Istanbul"
	dba, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db.DB = dba
	DBConn = db
}

func (db *Database) Migrate() {
	var err error
	for {
		for _, schema := range db.Credentials.schemas {
			DBConn.DB.Exec("CREATE SCHEMA IF NOT EXISTS " + schema + " AUTHORIZATION " + DBConn.Credentials.user + ";")
			fmt.Println("Migrating schema: " + schema)
			err = DBConn.DB.AutoMigrate(migrateModelList...)
			fmt.Println(err)
			for _, v := range migrateRelationList {
				if err := DBConn.DB.SetupJoinTable(v.model, v.field, v.joinTable); err != nil {
					fmt.Println(err)
				}
			}
		}
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			break
		}
	}
}

func (a Database) String() string {
	return fmt.Sprintf("%s:%s@%s:%s/%s", a.Credentials.user, a.Credentials.pass, a.Credentials.host, a.Credentials.port, a.Credentials.dBName)
}

func (db *Database) Close() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}

func (db *Database) CreateSchema(schema string, user ...string) {
	var userName string
	if len(user) > 1 {
		panic("CreateSchema only accept one user")
	} else if len(user) == 0 {
		userName = db.Credentials.user
	} else {
		userName = user[0]
	}
	if err := db.DB.Exec("CREATE SCHEMA IF NOT EXISTS " + schema + " AUTHORIZATION " + userName + ";").Error; err != nil {
		panic(err)
	}
}

func (db *Database) DropSchema(schema string) {
	if err := db.DB.Exec("DROP SCHEMA IF EXISTS " + schema + " CASCADE;").Error; err != nil {
		panic(err)
	}
}

func (db *Database) CreateTable(table interface{}) {
	if hasFlag := db.DB.Migrator().HasTable(table); !hasFlag {
		if err := db.DB.Migrator().CreateTable(table); err != nil {
			panic(err)
		}
	}
}

func (db *Database) DropTable(table interface{}) {
	if err := db.DB.Migrator().DropTable(table); err != nil {
		panic(err)
	}
}

func (db *Database) CreateJoinTable(model any, field string, joinTable any) {
	modelType := reflect.TypeOf(model).Elem().Kind()
	joinTableType := reflect.TypeOf(joinTable).Elem().Kind()
	if modelType != reflect.Struct && joinTableType != reflect.Struct {
		if err := db.DB.SetupJoinTable(model, field, joinTable); err != nil {
			panic(err)
		}
	}
}

func (db *Database) Add(structTemp interface{}) *gorm.DB {
	if tx := db.DB.Create(structTemp); tx.Error != nil {
		log.Printf("%s", tx.Error)
	}
	return db.DB
}

func (db *Database) Delete(structTemp interface{}, id uint) *gorm.DB {
	if tx := db.DB.Delete(structTemp, id); tx.Error != nil {
		log.Printf("%s", tx.Error)
	}
	return db.DB
}

func (db *Database) SeedSchema(models ...globals.Seeder) error {
	var (
		subjects []globals.Seeder
		fileByte []byte
		err      error
		// tempMember map[string]interface{}
	)
	if len(models) == 0 {
		subjects = seederModelList
	} else {
		subjects = models
	}
	for _, model := range subjects {
		fileName := model.Seed()
		fmt.Printf("%+v\n", model)
		if fileByte, err = os.ReadFile("db/seeds/" + fileName); err != nil {
			fmt.Println("asd", err)
			// return err
		}
		if err = json.Unmarshal(fileByte, &model); err != nil {
			fmt.Println("dsa", err)
			// return err
		}
		// newSlice := reflect.MakeMap(sliceType).Interface().(map[string]interface{})
		modelType := reflect.TypeOf(model).Elem()
		modelPtr2 := reflect.New(modelType)
		fmt.Printf("%s\n", modelPtr2) //3
	}
	return nil
}
