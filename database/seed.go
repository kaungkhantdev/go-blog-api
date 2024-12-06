package database

import (
	iconModel "go-blog-api/internal/icon/models"
	reactionTypeModel "go-blog-api/internal/reaction_type/models"
	"log"

	"gorm.io/gorm"
)

func Seed() {
	DB, err := Connect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	seedIcons(DB)
	seedReactionTypes(DB)
	log.Println("Seed operation succeed.")
}

func seedIcons(db *gorm.DB) {
	icons := []iconModel.Icon{
		{Icon: "love", Url: ""},
		{Icon: "unlike", Url: ""},
		{Icon: "haha", Url: ""},
		{Icon: "like", Url: ""},
	}

	seedTable("Icons", &iconModel.Icon{}, icons, db)
}

func seedReactionTypes(db *gorm.DB) {
	reactionTypes := []reactionTypeModel.ReactionType{
		{Type: "love", IconId: 1},
		{Type: "unlike", IconId: 2},
		{Type: "haha", IconId: 3},
		{Type: "like", IconId: 4},
	}

	seedTable("ReactionTypes", &reactionTypeModel.ReactionType{}, reactionTypes, db)
}

func seedTable(tableName string, model interface{}, data interface{}, db *gorm.DB) {
	var count int64
	if err := db.Model(model).Count(&count).Error; err != nil {
		log.Printf("Error counting %s: %v", tableName, err)
		return
	}

	if count > 0 {
		log.Printf("%s already seeded, skipping.", tableName)
		return
	}

	if err := db.Create(data).Error; err != nil {
		log.Printf("Error seeding %s: %v", tableName, err)
	} else {
		log.Printf("%s seeded successfully.", tableName)
	}
}
