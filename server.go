// Package server implements the server side of VFT-'s client-server interaction
package praypi

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

type Server struct {
}

type Request struct {
	Id        uuid.UUID `json:"id"`
	Type      string    `json:"type"`
	Lang      string    `json:"lang"`
	Fulfilled bool      `json:"fulfilled"`
	Content   string    `json:"content"`
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
		if isValidPrayer(json) {
			c.JSON(200, gin.H{
				"message": "Your prayer has been heard!",
				"id":      uuid.Must(uuid.NewV4()),
			})
		} else if json.Lang == "angelic" {
			c.JSON(301, gin.H{
				"message": "Moved permanently",
			})
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

func (s *Server) getRequests(c *gin.Context) {
	c.JSON(501, gin.H{
		"error": "Not implemented yet. Come back later.",
	})
}

func (s *Server) getRequestsId(c *gin.Context) {
	c.JSON(501, gin.H{
		"error": "Not implemented yet. Come back later.",
	})
}

func isValidPrayer(r Request) bool {
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
		return true
	}
	return false
}

func isValidLanguage(t string) bool {
	if t == "human" {
		return true
	} else {
		return false
	}

}
