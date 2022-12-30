package helper

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/kekasicoid/kekasigohelper"
	"github.com/kekasicoid/restapi-socketio/enum"
	"github.com/kekasicoid/restapi-socketio/table"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstanceMySQLConn *gorm.DB
var dbOnceMySQLConn sync.Once
var rdb *redis.Client

func DBRedisConn() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.Get("REDIS_ADDRESS").(string) + ":" + viper.Get("REDIS_PORT").(string),
		Password: viper.Get("REDIS_PASSWORD").(string), // no password set
		DB:       0,                                    // use default DB
	})
	var ctx = context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		kekasigohelper.LoggerFatal("DBRedisConn : Cannot Connect")
	}
	return rdb
}

func DBMysqlConn() *gorm.DB {
	if dbInstanceMySQLConn == nil {
		dbOnceMySQLConn.Do(openDBMysqlConn)
	}
	return dbInstanceMySQLConn
}

func openDBMysqlConn() {
	dbHost := viper.Get("DB_HOST").(string)
	dbPort := viper.Get("DB_PORT").(string)
	dbUser := viper.Get("DB_USER").(string)
	dbPass := viper.Get("DB_PASSWORD").(string)
	dbName := viper.Get("DB_NAME").(string)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName) //Build connection string

	gormDB, err := gorm.Open(
		mysql.Open(connectionString),
		&gorm.Config{SkipDefaultTransaction: true},
	)

	if err != nil {
		kekasigohelper.LoggerFatal("dbconn : cannot open database")
	}
	dbInstanceMySQLConn = gormDB
	db, err := dbInstanceMySQLConn.DB()
	if err != nil {
		kekasigohelper.LoggerFatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
}

func DbMirgrator(dbConn *gorm.DB) {
	if !dbConn.Migrator().HasTable(table.Orang{}) {
		dbConn.Debug().AutoMigrate(
			table.Orang{},
			table.Asset{},
		)
		if err := dbConn.First(&table.Orang{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			addOrang := []table.Orang{
				{
					Id:           1,
					Nama:         "Bani",
					JenisKelamin: enum.Pria,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           2,
					Nama:         "Budi",
					JenisKelamin: enum.Pria,
					OrangTua:     1,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           3,
					Nama:         "Nida",
					JenisKelamin: enum.Wanita,
					OrangTua:     1,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           4,
					Nama:         "Andi",
					JenisKelamin: enum.Pria,
					OrangTua:     1,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           5,
					Nama:         "Sigit",
					JenisKelamin: enum.Pria,
					OrangTua:     1,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           6,
					Nama:         "Hari",
					JenisKelamin: enum.Pria,
					OrangTua:     2,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           7,
					Nama:         "Siti",
					JenisKelamin: enum.Wanita,
					OrangTua:     2,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           8,
					Nama:         "Bila",
					JenisKelamin: enum.Wanita,
					OrangTua:     3,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           9,
					Nama:         "Lesti",
					JenisKelamin: enum.Wanita,
					OrangTua:     3,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           10,
					Nama:         "Diki",
					JenisKelamin: enum.Pria,
					OrangTua:     4,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           11,
					Nama:         "Doni",
					JenisKelamin: enum.Pria,
					OrangTua:     5,
					ModelDate:    table.ModelDate{},
				}, {
					Id:           12,
					Nama:         "Toni",
					JenisKelamin: enum.Pria,
					OrangTua:     5,
					ModelDate:    table.ModelDate{},
				},
			}
			dbConn.Create(&addOrang)
			addAsset := []table.Asset{
				{
					Id:          1,
					OrangID:     2,
					IdProduct:   3,
					Tittle:      "Samsung Universe 9",
					Description: "Samsung's new variant which goes beyond Galaxy to the Universe",
					Price:       1249,
					Brand:       "Samsung",
					Category:    "smartphones",
					Thumbnail:   "https://i.dummyjson.com/data/products/3/thumbnail.jpg",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          2,
					OrangID:     2,
					IdProduct:   7,
					Tittle:      "Samsung Galaxy Book",
					Description: "Samsung Galaxy Book S (2020) Laptop With Intel Lakefield Chip, 8GB of RAM Launched",
					Price:       1499,
					Brand:       "Samsung",
					Category:    "laptops",
					Thumbnail:   "https://i.dummyjson.com/data/products/7/thumbnail.jpg",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          3,
					OrangID:     6,
					IdProduct:   1,
					Tittle:      "iPhone 9",
					Description: "An apple mobile which is nothing like apple",
					Price:       549,
					Brand:       "Apple",
					Category:    "smartphones",
					Thumbnail:   "https://i.dummyjson.com/data/products/1/thumbnail.jpg",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          4,
					OrangID:     7,
					IdProduct:   2,
					Tittle:      "iPhone X",
					Description: "SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...",
					Price:       899,
					Brand:       "Apple",
					Category:    "smartphones",
					Thumbnail:   "https://i.dummyjson.com/data/products/2/thumbnail.jpg",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          5,
					OrangID:     3,
					IdProduct:   5,
					Tittle:      "Huawei P30",
					Description: "Huawei’s re-badged P30 Pro New Edition was officially unveiled yesterday in Germany and now the device has made its way to the UK.",
					Price:       499,
					Brand:       "499",
					Category:    "Huawei",
					Thumbnail:   "smartphones",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          6,
					OrangID:     8,
					IdProduct:   3,
					Tittle:      "Samsung Universe 9",
					Description: "Samsung's new variant which goes beyond Galaxy to the Universe",
					Price:       1249,
					Brand:       "Samsung",
					Category:    "smartphones",
					Thumbnail:   "https://i.dummyjson.com/data/products/3/thumbnail.jpg",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          7,
					OrangID:     9,
					IdProduct:   5,
					Tittle:      "Huawei P30",
					Description: "Huawei’s re-badged P30 Pro New Edition was officially unveiled yesterday in Germany and now the device has made its way to the UK.",
					Price:       499,
					Brand:       "499",
					Category:    "Huawei",
					Thumbnail:   "smartphones",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          8,
					OrangID:     9,
					IdProduct:   2,
					Tittle:      "iPhone X",
					Description: "SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...",
					Price:       899,
					Brand:       "Apple",
					Category:    "smartphones",
					Thumbnail:   "https://i.dummyjson.com/data/products/2/thumbnail.jpg",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          9,
					OrangID:     4,
					IdProduct:   3,
					Tittle:      "Samsung Universe 9",
					Description: "Samsung's new variant which goes beyond Galaxy to the Universe",
					Price:       1249,
					Brand:       "Samsung",
					Category:    "smartphones",
					Thumbnail:   "https://i.dummyjson.com/data/products/3/thumbnail.jpg",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          10,
					OrangID:     10,
					IdProduct:   7,
					Tittle:      "Samsung Galaxy Book",
					Description: "Samsung Galaxy Book S (2020) Laptop With Intel Lakefield Chip, 8GB of RAM Launched",
					Price:       1499,
					Brand:       "Samsung",
					Category:    "laptops",
					Thumbnail:   "https://i.dummyjson.com/data/products/7/thumbnail.jpg",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          11,
					OrangID:     5,
					IdProduct:   5,
					Tittle:      "Huawei P30",
					Description: "Huawei’s re-badged P30 Pro New Edition was officially unveiled yesterday in Germany and now the device has made its way to the UK.",
					Price:       499,
					Brand:       "499",
					Category:    "Huawei",
					Thumbnail:   "smartphones",
					ModelDate:   table.ModelDate{},
				}, {
					Id:          12,
					OrangID:     11,
					IdProduct:   2,
					Tittle:      "iPhone X",
					Description: "SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...",
					Price:       899,
					Brand:       "Apple",
					Category:    "smartphones",
					Thumbnail:   "https://i.dummyjson.com/data/products/2/thumbnail.jpg",
					ModelDate:   table.ModelDate{},
				},
			}
			dbConn.Create(&addAsset)
		}
	}

}
