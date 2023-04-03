package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RaghibA/pokemon-tcg-go/internal/database"
	"github.com/RaghibA/pokemon-tcg-go/internal/models"
	"github.com/gin-gonic/gin"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

func QueryCardHandler(c *gin.Context) {
	client := tcg.NewClient(os.Getenv("API_SECRET"))
	queryString := fmt.Sprintf("name:%s", c.Query("q"))
	cards, err := client.GetCards(
		request.Query(queryString),
		request.PageSize(20),
	)
	if err != nil {
		log.Println("Failed to get cards")
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": "Failed to retrieve cards",
			"code":    4090001,
		})

		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"cards": cards,
	})
}

// Add user card
func AddCardHandler(c *gin.Context) {
	client := tcg.NewClient(os.Getenv("API_SECRET"))

	var cardBody struct {
		Id string `json:"id"`
	}

	if c.Bind(&cardBody) != nil { // map body to struct
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"code":    4000001,
		})

		return
	}

	// Get specific card
	card, err := client.GetCardByID(cardBody.Id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Unable to find card with specified ID",
			"code":    4040001,
		})

		return
	}

	// Get user from cookie
	user := models.User{}
	u, _ := c.Get("user")
	database.PokeDb.Db.Find(&user, u)

	// check if card exists in users deck
	hasCard := models.Card{}
	database.PokeDb.Db.Where("card_id = ?", cardBody.Id).Where("user_id = ?", user.ID).First(&hasCard)
	if hasCard.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The selected card exists in the users deck",
			"code":    4090001,
		})

		return
	}

	// Create card record, attach user ID & commit to db
	newCard := models.Card{
		Name:   card.Name,
		Types:  card.Types,
		Value:  *card.CardMarket.Prices.AverageSellPrice,
		Img:    card.Images.Small,
		CardId: card.ID,
		UserId: user.ID,
	}
	database.PokeDb.Db.Create(&newCard)
}

// Get user cards
func GetCardsHandler(c *gin.Context) {
	cards := []models.Card{}

	// Get user from jwt cookie
	user := models.User{}
	u, _ := c.Get("user")
	database.PokeDb.Db.Find(&user, u)
	userId := user.ID

	// Find all cards with the userid
	database.PokeDb.Db.Find(&cards, "user_id = ?", userId)
	if len(cards) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No cards found for user",
			"code":    4040002,
		})

		return
	}

	// return array with all cards
	c.JSON(http.StatusAccepted, gin.H{
		"cards": cards,
	})
}

// Delete user card
// Get user cards
func DeleteCardHandler(c *gin.Context) {

	// Get user from jwt cookie
	user := models.User{}
	u, _ := c.Get("user")
	database.PokeDb.Db.Find(&user, u)
	userId := user.ID

	// Get card Id from query params
	cardId := c.Query("id")
	card := models.Card{}
	database.PokeDb.Db.Where("user_id = ?", userId).Delete(&card, cardId)
	c.JSON((http.StatusOK), gin.H{})
}
