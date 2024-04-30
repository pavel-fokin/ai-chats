package sqlite

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"pavel-fokin/ai/apps/ai-bot/internal/app"
)

type Sqlite struct {
	db *gorm.DB
}

func New(url string) (app.ChatDB, func() error) {
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(Actor{}, Chat{}, Message{}, ChatToActor{})

	// aiActor := Actor{Type: "ai"}
	// userActor := Actor{Type: "user"}
	// db.Create(&aiActor)
	// db.Create(&userActor)

	// chat := Chat{Actors: []Actor{aiActor, userActor}}
	// db.Create(&chat)

	db_, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	return &Sqlite{
		db: db,
	}, db_.Close
}
