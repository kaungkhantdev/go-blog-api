package database

import (
	"fmt"
	"log"

	"go-blog-api/internal/modules/blog/article"
	"go-blog-api/internal/modules/blog/bookmark"
	"go-blog-api/internal/modules/blog/comment"
	"go-blog-api/internal/modules/blog/icon"
	"go-blog-api/internal/modules/blog/reaction"
	"go-blog-api/internal/modules/blog/reaction_type"
	"go-blog-api/internal/modules/blog/tag"
	"go-blog-api/internal/modules/core/otp"
	"go-blog-api/internal/modules/core/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migration() {
	// Load database credentials from environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// Establish connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Successfully connected to the database.")

	// Run the migrations
	err = runMigrations(db)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Database migration completed successfully.")
}

func runMigrations(DB *gorm.DB) error {
	// Run AutoMigrate for each model
	if err := DB.AutoMigrate(
		&user.UserEntity{},
		&otp.OtpEntity{},
		&icon.IconEntity{},
		&tag.TagEntity{},
		&article.ArticleEntity{},
		&bookmark.BookmarkEntity{},
		&comment.CommentEntity{},
		&reaction.ReactionEntity{},
		&reaction_type.ReactionTypeEntity{},
	); err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}

	return nil
}
