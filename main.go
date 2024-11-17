package main

import (
	"net/http"
	"time"
    "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


//JWT TOKEN 
	var jwtKey = []byte("my_secret_key")

        type Credentials struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }

        type Claims struct {
            Username string `json:"username"`
            jwt.StandardClaims
        }

		var users = []Credentials{
			{Username: "user", Password: "password"},
			{Username: "user1", Password: "password1"},
			{Username: "user2", Password: "password2"},
			{Username: "user3", Password: "password3"},
		}

		// Генерация access и refresh токенов
		func generateToken(username string) (string, string, error) {
			expirationTime := time.Now().Add(1 * time.Minute) // Время жизни access токена
			refreshExpirationTime := time.Now().Add(24 * time.Hour) // Время жизни refresh токена
		
			// Создание access токена
			claims := &Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			},
			}
		
			accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			accessTokenString, err := accessToken.SignedString(jwtKey)
			if err != nil {
			return "", "", err
			}
		
			// Создание refresh токена
			refreshClaims := &Claims{
			Username: username,
			StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
			},
			}
		
			refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
			refreshTokenString, err := refreshToken.SignedString(jwtKey)
			if err != nil {
			return "", "", err
			}
		
			return accessTokenString, refreshTokenString, nil
		}


        /*func generateToken(username string) (string, error) {
            expirationTime := time.Now().Add(5 * time.Minute)
            claims := &Claims{
                Username: username,
                StandardClaims: jwt.StandardClaims{
                    ExpiresAt: expirationTime.Unix(),
                },
            }
            token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
            return token.SignedString(jwtKey)
		}*/

		func login(c *gin.Context) {
            var creds Credentials
            if err := c.BindJSON(&creds); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
                return
            }

            var validUser *Credentials
				for _, user := range users {
					if user.Username == creds.Username && user.Password == creds.Password {
						validUser = &user
						break
					}
				}

			if validUser == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				return
			}

            accesstoken, refreshtoken, err := generateToken(creds.Username)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
                return
            }

            c.JSON(http.StatusOK, gin.H{"accesstoken": accesstoken, "refreshtoken": refreshtoken})
        }

        func authMiddleware() gin.HandlerFunc {
            return func(c *gin.Context) {
                tokenString := c.GetHeader("Authorization")

                claims := &Claims{}
                token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
                    return jwtKey, nil
                })

                if err != nil || !token.Valid {
					if ve, ok := err.(*jwt.ValidationError); ok {
						if ve.Errors&jwt.ValidationErrorExpired != 0 {
							c.JSON(http.StatusUnauthorized, gin.H{"message": "token expired"})
							c.Abort()
							return
						}
					}
                    c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
                    c.Abort()
                    return
                }

                c.Next()
            }
        }



type Item struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Category_id   int     `json:"category_id"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	QuantityStock int     `json:"quantityStock"`
}

var items = []Item{
	{ID: "1", Name: "Кружка", Category_id: 1, Description: "Крутой подарок", Price: 45.0, QuantityStock: 5},
	{ID: "2", Name: "Тарелка", Category_id: 2, Description: "Крутой подарок", Price: 45.0, QuantityStock: 5},
	{ID: "3", Name: "Мягкая игрушка", Category_id: 3, Description: "Крутой подарок", Price: 45.0, QuantityStock: 5},
}

var itemsBasket = []Item{}

var router = gin.Default()

func main() {
	//router := gin.Default()
	router.POST("/login", login)

    protected := router.Group("/")
    protected.Use(authMiddleware())
        {
			protected.GET("/items", getItems)

			// Получение всех товаров из корзины
			protected.GET("/id/1/basket", getItemsBasket)

			// Получение товара по ID
			protected.GET("/items/:id", getItemByID)

			// Создание нового товара
			protected.POST("/items", createItem)

			// Добавление товара в корзину
			protected.POST("/id/1/basket", createItemBasket)

			// Обновление существующего товара
			protected.PUT("/items/:id", updateItem)

			// Удаление товара
			protected.DELETE("/items/:id", deleteItem)

			// Удаление товара из корзины
			protected.DELETE("/id/1/basket", deleteItemBasket)
        }


	router.Run(":8080")
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func getItemsBasket(c *gin.Context) {
	c.JSON(http.StatusOK, itemsBasket)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func createItem(c *gin.Context) {
	var newItem Item

	if err := c.BindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	items = append(items, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func createItemBasket(c *gin.Context) {
    
    var body Item
    c.ShouldBindJSON(&body)
	var newItem Item
    id := body.ID

    for _, item := range items {
		if item.ID == id {
            newItem = item
			break
		}
	}

	if newItem.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	itemsBasket = append(itemsBasket, newItem)
	c.JSON(http.StatusOK, newItem)
}

func updateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem Item

	if err := c.BindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i] = updatedItem
			c.JSON(http.StatusOK, updatedItem)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "item deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "item not found"})
}

func deleteItemBasket(c *gin.Context) {
	var body Item
    c.ShouldBindJSON(&body)
    id := body.ID

	for i, item := range itemsBasket {
		if item.ID == id {
			itemsBasket = append(itemsBasket[:i], itemsBasket[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "item from basket deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "item in basket not found"})
}
