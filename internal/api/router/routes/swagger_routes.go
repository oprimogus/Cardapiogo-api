package routes

import (
	"net/http"
	"os"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/oprimogus/cardapiogo/docs"
	"github.com/oprimogus/cardapiogo/internal/config"
	"github.com/oprimogus/cardapiogo/internal/errors"
)

func SwaggerRoutes(router *gin.Engine) {
	basePath := config.GetInstance().Api.BasePath()
	v1 := router.Group(basePath)

	v1.GET("/v1/reference/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1.GET("/v2/reference", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Cardapiogo",
			},
			DarkMode: true,
		})
		if err != nil {
			c.JSON(500, xerrors.InternalServerError(err.Error()))
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
}
