package categories

import (
	"disspace/business/categories"
	"disspace/controllers"

	// "disspace/controllers/categories/request"
	"disspace/controllers/categories/request"
	"disspace/controllers/categories/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoriesController struct {
	CategoriesUseCase categories.UseCase
}

func NewCategoriesController(categoriesUseCase categories.UseCase) *CategoriesController {
	return &CategoriesController{
		CategoriesUseCase: categoriesUseCase,
	}
}

func (controller *CategoriesController) GetAll(c echo.Context) error {
	categories := []response.Categories{}
	ctx := c.Request().Context()

	result, err := controller.CategoriesUseCase.GetAll(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	for _, i := range result {
		categories = append(categories, response.FromDomain(i))
	}

	return controllers.NewSuccessResponse(c, categories)
}

func (controller *CategoriesController) Create(c echo.Context) error {
	createCategory := request.Categories{}
	c.Bind(&createCategory)

	ctx := c.Request().Context()

	_, err := controller.CategoriesUseCase.Create(ctx, createCategory.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, "successfully created new category")
}

func (controller *CategoriesController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	result, err := controller.CategoriesUseCase.GetByID(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomain(result))
}

func (controller *CategoriesController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	err := controller.CategoriesUseCase.Delete(ctx, id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully deleted category")
}

func (controller *CategoriesController) Update(c echo.Context) error {
	ctx := c.Request().Context()
	data := request.Categories{}
	c.Bind(&data)
	id := c.Param("id")
	// fmt.Println(data)
	err := controller.CategoriesUseCase.Update(ctx, data.ToDomain(), id)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, "successfully updated category")

}
