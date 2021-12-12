package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"TwitterClone/api/models"

	"github.com/go-pg/pg/v10/orm"
	"github.com/go-pg/pg/v10"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresDB struct {
	DB      *gorm.DB
	MysqlDB *gorm.DB
	NewDB   *pg.DB
}

func (p *PostgresDB) InitNewDB(){
	p.NewDB = pg.Connect(&pg.Options{
        User: "darienmiller88",
		Password: "nintendowiiu000",
		Database: "TwitterClone",
    })
	 
	if err := p.NewDB.Ping(context.Background()); err != nil{
		fmt.Println("connection dead")
	}else{
		fmt.Println("Connected to DB!")
	}

	for _, model := range []interface{}{&models.ReplyBase{}, &models.Reply2{}}{
		err := p.NewDB.Model(model).CreateTable(&orm.CreateTableOptions{
			FKConstraints: true,
		})

		fmt.Println("err:", err)
	}
	

}

//Method to connect to the Postgres database.
func (p *PostgresDB) InitializeDatabase() {
	var err error
	db_url := os.Getenv("DATABASE_URL")

	p.DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}
	// p.mysqlConnection()

	fmt.Println("Connected to DB!")
}

func (p *PostgresDB) mysqlConnection() {
	password := "nintendowiiu000"
	user := "root"
	dbname := "twitterclone"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	p.MysqlDB = db
}

func (p *PostgresDB) migrateModels(){
// p.MysqlDB.Migrator().DropTable(&models.User{}, &models.Tweet{}, &models.Reply{}, &models.Like{}, &models.Retweet{},
	// 	&models.QuoteTweet{}, &models.Message{})
	p.DB.AutoMigrate(&models.User{}, &models.Tweet{}, &models.Reply{}, &models.Like{}, &models.Retweet{},
		&models.QuoteTweet{}, &models.Message{})

	//GORM forces you to manually create EACH FUCKING foreign key...
	p.DB.Debug().Migrator().CreateConstraint(&models.Tweet{}, "User")
	p.DB.Debug().Migrator().CreateConstraint(&models.Like{}, "User")
	p.DB.Debug().Migrator().CreateConstraint(&models.Like{}, "Tweet")
	p.DB.Debug().Migrator().CreateConstraint(&models.Retweet{}, "User")
	p.DB.Debug().Migrator().CreateConstraint(&models.Retweet{}, "Tweet")
	p.DB.Debug().Migrator().CreateConstraint(&models.QuoteTweet{}, "User")
	p.DB.Debug().Migrator().CreateConstraint(&models.QuoteTweet{}, "Tweet")
	p.DB.Debug().Migrator().CreateConstraint(&models.Message{}, "Sender")
	p.DB.Debug().Migrator().CreateConstraint(&models.Message{}, "Receiver")
	p.DB.Debug().Migrator().CreateConstraint(&models.Reply{}, "Replier")
	p.DB.Debug().Migrator().CreateConstraint(&models.Reply{}, "Replied")
}