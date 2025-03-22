package http

import (
	"ecommerce_clean/internals/product/controller/dto"
	"ecommerce_clean/internals/product/entity"
	"ecommerce_clean/internals/product/usecase"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/response"
	"ecommerce_clean/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	usecase usecase.IProductUseCase
}

func NewProductHandler(usecase usecase.IProductUseCase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

//		@Summary	 Retrieve a list of products
//	 @Description Fetches a paginated list of products based on the provided filter parameters.
//		@Tags		 Products
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the list of products"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	var req dto.ListProductRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get query", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	products, pagination, err := h.usecase.ListProducts(c, &req)
	if err != nil {
		logger.Error("Failed to get products", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get products")
		return
	}

	var res dto.ListProductResponse
	utils.MapStruct(&res.Products, products)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Retrieve a product by its ID
//	 @Description Fetches the details of a specific product based on the provided product ID.
//		@Tags		 Products
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Successfully retrieved the product"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 404	{object}	response.Response	"Not Found - Event with the specified ID not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	var res entity.Product

	productId := c.Param("id")

	product, err := h.usecase.GetProductById(c, productId)
	if err != nil {
		logger.Error("Failed to get product detail: ", err)
		logger.Info(err.Error())
		switch err.Error() {
		case "record not found":
			response.Error(c, http.StatusNotFound, err, err.Error())
		default:
			response.Error(c, http.StatusInternalServerError, err, err.Error())
			return
		}

		return
	}

	utils.MapStruct(&res, product)
	response.JSON(c, http.StatusOK, res)
}

//		@Summary	 Create a new product
//	 @Description Creates a new product based on the provided details.
//		@Tags		 Products
//		@Produce	 json
//		@Param		 _	body	dto.CreateProductRequest	  true	"Body"
//		@Success	 201	{object}	response.Response	"Product created successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req dto.CreateProductRequest

	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	if err := h.usecase.CreateProduct(c, &req); err != nil {
		logger.Error("Failed to create product", err)

		switch utils.ExtractConstraintName(err) {
		case "unique_product_code":
			response.Error(c, http.StatusConflict, err, "Code already in use")
		case "unique_product_name":
			response.Error(c, http.StatusConflict, err, "Name already in use")
		default:
			response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		}
		return
	}

	response.JSON(c, http.StatusCreated, "Create product successfully")
}

//		@Summary	 Update a product
//	 @Description Update a new product based on the provided details.
//		@Tags		 Products
//		@Produce	 json
//		@Param		 _	body	dto.UpdateProductRequest	  true	"Body"
//		@Success	 201	{object}	response.Response	"Product updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	var req dto.UpdateProductRequest

	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	productId := c.Param("id")

	if err := h.usecase.UpdateProduct(c, productId, &req); err != nil {
		logger.Error("Failed to update product", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	response.JSON(c, http.StatusOK, "Update product successfully")
}

//		@Summary	 Delete a product
//	 @Description Delete a new product based on the provided details.
//		@Tags		 Products
//		@Produce	 json
//		@Success	 201	{object}	response.Response	"Delete updated successfully"
//		@Failure	 401	{object}	response.Response	"Unauthorized - User not authenticated"
//		@Failure	 403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	productId := c.Param("id")

	err := h.usecase.DeleteProduct(c, productId)

	if err != nil {
		logger.Error("Failed to delete products: ", err)
		response.Error(c, http.StatusNotFound, err, "Not found")
		return
	}

	response.JSON(c, http.StatusOK, "Delete products successfully")
}
