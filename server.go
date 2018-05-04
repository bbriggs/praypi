package praypi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Server struct {
	db  *gorm.DB
	log *logrus.Entry
}

type Request struct {
	Id        uuid.UUID `json:"id"`
	Type      string    `json:"type"`
	Lang      string    `json:"lang"`
	Fulfilled bool      `json:"fulfilled"`
	Content   string    `json:"content"`
}

func NewServer(dbUser string, dbPass string, dbName string, dbHost string, dbPort string) *Server {
	db := dbConnect(dbUser, dbPass, dbName, dbHost, dbPort)
	var s = &Server{
		log: logrus.WithField("context", "praypi"),
		db:  db,
	}
	return s
}

func (s *Server) Run() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to PrayPI, the API for prayer requests.\n")
	})
	r.POST("/requests", s.postRequests)
	r.GET("/requests", s.getRequests)
	r.GET("/requests/:id", s.getRequestsId)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func (s *Server) postRequests(c *gin.Context) {
	var json Request
	if err := c.ShouldBindJSON(&json); err == nil {
		if s.isValidPrayer(json) {
			code, resp := s.parsePrayer(json)
			c.JSON(code, resp)
		} else {
			c.JSON(400, gin.H{
				"error": "Malformed prayer.",
			})
		}
	} else {
		c.JSON(400, gin.H{
			"error": "Malformed prayer.",
		})
	}
}

func (s *Server) parsePrayer(r Request) (int, gin.H) {
	id, err := s.insertPrayer(r)
	if err != nil {
		return 500, gin.H{"error": "Unable to handle prayers right now. Pray again later."}
	}

	resp := gin.H{
		"id":      id,
		"message": "Your prayer has been heard!",
	}
	return 200, resp
}

func (s *Server) getRequests(c *gin.Context) {
	prayers, err := s.getAllPrayers()
	if err != nil {
		c.JSON(200, gin.H{
			"error": err})
	} else {
		c.JSON(200, prayers)
	}
}

func (s *Server) getRequestsId(c *gin.Context) {
	id := c.Param("id")
	resp, err := s.queryPrayer(id)
	if err != nil {
		c.JSON(200, gin.H{"error": "Prayer not found."})
	} else {
		c.JSON(200, resp)
	}
}

func (s *Server) isValidPrayer(r Request) bool {
	fmt.Println(r)
	if isValidPrayerType(r.Type) && isValidLanguage(r.Lang) && len(r.Content) > 0 {
		return true
	} else {
		return false
	}
}

func isValidPrayerType(t string) bool {
	switch t {
	case
		"adoration",
		"confession",
		"thanksgiving",
		"supplication",
		"imprecation",
		"unspoken":
		fmt.Println("Valid type")
		return true
	}
	return false
}

func isValidLanguage(t string) bool {
	fmt.Println(t)
	if t == "human" {
		fmt.Println("Valid language")
		return true
	} else {
		return false
	}

}
