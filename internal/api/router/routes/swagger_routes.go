package routes

import (
	"net/http"
	"os"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/oprimogus/cardapiogo/internal/errors"
)

func SwaggerRoutes(router *gin.Engine) {

	basePath := os.Getenv("API_BASE_PATH")
	v1 := router.Group(basePath)

	// SWAGGER
	v1.StaticFile("/docs/swagger.json", "./docs/swagger.json")
	v1.GET("/v1/reference/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/api/docs/swagger.json")))
	
	v1.GET("/v2/reference", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			Theme: scalar.ThemeKepler,
			Layout: scalar.LayoutModern,
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Cardapiogo",
			},
			DarkMode: true,
			Authentication: "Bearer",
		})
		if err != nil {
			c.JSON(500, errors.InternalServerError(err.Error()))
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})

}
