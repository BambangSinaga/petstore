package http

import (
	"net/http"
	"strconv"

	models "petstore/pet"

	petUcase "petstore/pet/usecase"
	"github.com/labstack/echo"
)

type HttpPetHandler struct {
	AUsecase petUcase.PetUsecase
}

func (a *HttpPetHandler) FetchPet(c echo.Context) error {

	numS := c.QueryParam("num")
	num, _ := strconv.Atoi(numS)

	cursor := c.QueryParam("cursor")

	listAr, nextCursor, err := a.AUsecase.Fetch(cursor, int64(num))

	statusCode := getStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	c.Response().Header().Set(`X-Cursor`, nextCursor)
	return c.JSON(http.StatusOK, listAr)
}

func (a *HttpPetHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	art, err := a.AUsecase.GetByID(id)
	statusCode := getStatusCode(err)

	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.JSON(http.StatusOK, art)
}

func (a *HttpPetHandler) Store(c echo.Context) error {
	var pet models.Pet
	err := c.Bind(&pet)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ar, err := a.AUsecase.Store(&pet)
	statusCode := getStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.JSON(http.StatusCreated, ar)
}

func (a *HttpPetHandler) Update(c echo.Context) error {
	var pet models.Pet
	err := c.Bind(&pet)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	
	idP, err := strconv.Atoi(c.Param("id"))
	pet.ID = int64(idP)

	ar, err := a.AUsecase.Update(&pet)
	statusCode := getStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.JSON(http.StatusCreated, ar)
}

func (a *HttpPetHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	id := int64(idP)

	_, err = a.AUsecase.Delete(id)
	statusCode := getStatusCode(err)
	if err != nil {
		return c.JSON(statusCode, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

func getStatusCode(err error) int {
	switch err {
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.CONFLICT_ERROR:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func NewPetHttpHandler(e *echo.Echo, us petUcase.PetUsecase) {
	handler := &HttpPetHandler{
		AUsecase: us,
	}

	e.GET("/pet", handler.FetchPet)
	e.POST("/pet", handler.Store)
	e.GET("/pet/:id", handler.GetByID)
	e.PUT("/pet/:id", handler.Update)
	e.DELETE("/pet/:id", handler.Delete)
	e.POST("/pet/:id/uploadImage", handler.Update)

}
