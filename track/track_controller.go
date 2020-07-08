package track

import (
	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
)

// Controller handles all track routes
type Controller interface {
	Search() gin.HandlerFunc
}

type controller struct {
	service Service
}

// NewController initializes a new Track Controller
func NewController(service Service) Controller {
	return controller{service: service}
}

func (controller controller) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.Query("query")
		if query == "" {
			c.JSON(200, map[string]string{})
			return
		}

		client, _ := c.Get("spotify-client")

		tracks, err := controller.service.SearchTrack(query, client.(spotify.Client))
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, map[string][]Track{"tracks": tracks})
	}
}
