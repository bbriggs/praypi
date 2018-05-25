package praypi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satori/go.uuid"
	"os"
)

type Prayer struct {
	ID        int       `gorm:"primary_key" json:"-"`
	PrayerID  uuid.UUID `gorm:"unique;not_null" json:"id"`
	Type      string    `gorm:"not_null" json:"type"`
	Language  string    `gorm:"not_null" json:"lang"`
	Fulfilled bool      `gorm:"not_null" json:"fulfilled"`
	Content   string    `gorm:"not_null" json:"content"`
}

func dbConnect(user string, pass string, dbname string, host string, port string) *gorm.DB {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", user, dbname, pass, host, port)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Unable to connect to to database")
		fmt.Println(err)
		os.Exit(1)
	}

	db.Table("prayers").CreateTable(&Prayer{})
	return db
}

func (s *Server) insertPrayer(r Request) (uuid.UUID, error) {
	tx := s.db.Begin()
	p := &Prayer{
		PrayerID:  uuid.Must(uuid.NewV4()),
		Type:      r.Type,
		Language:  r.Lang,
		Fulfilled: false,
		Content:   r.Content,
	}
	if err := tx.Create(p).Error; err != nil {
		tx.Rollback()
		return p.PrayerID, err
	}

	return p.PrayerID, tx.Commit().Error
}

func (s *Server) queryPrayer(id string) (gin.H, error) {
	p := Prayer{}

	result := s.db.Where("prayer_id = ?", id).First(&p)
	if result.Error != nil {
		return gin.H{"error": "Database error"}, result.Error
	}
	resp := gin.H{
		"id":        p.PrayerID,
		"type":      p.Type,
		"lang":      p.Language,
		"fulfilled": p.Fulfilled,
		"content":   p.Content,
	}
	return resp, nil
}

func (s *Server) getAllPrayers() (gin.H, error) {
	var prayers []Prayer
	rows := s.db.Select("prayer_id, content, language, type, fulfilled").Find(&prayers)
	if rows.Error != nil {
		return gin.H{"error": "Database error"}, rows.Error
	}
	return gin.H{"prayers": prayers}, rows.Error
}
