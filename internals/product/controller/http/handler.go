package http

import (
	"ecommerce_clean/configs"
	"ecommerce_clean/internals/product/controller/dto"
	"ecommerce_clean/internals/product/entity"
	"ecommerce_clean/internals/product/usecase"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/redis"
	"ecommerce_clean/pkgs/response"
	"ecommerce_clean/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	usecase usecase.IProductUseCase
	cache   redis.IRedis
}

func NewProductHandler(usecase usecase.IProductUseCase, cache redis.IRedis) *ProductHandler {
	return &ProductHandler{usecase: usecase, cache: cache}
}

//	@Summary		Retrieve a list of products
//	@Description	Fetches a paginated list of products based on the provided filter parameters.
//	@Tags			Products
//	@Produce		json
//	@Param			search		query	string	false	"Search keyword for products"
//	@Param			page		query	int		false	"Page number (default: 1)"
//	@Param			size		query	int		false	"Number of items per page (default: 10)"
//	@Param			order_by	query	string	false	"Field to sort by"
//	@Param			order_desc	query	bool	false	"Sort in descending order (true/false)"
//	@Param			take_all	query	bool	false	"Retrieve all products without pagination"
//	@Success		200			{object}	response.Response	"Successfully retrieved the list of products"
//	@Failure		400			{object}	response.Response	"Bad Request - Invalid query parameters"
//	@Failure		500			{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router			/api/v1/products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	var req dto.ListProductRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get query", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	var res dto.ListProductResponse
	cacheKey := c.Request.URL.RequestURI()
	err := h.cache.Get(cacheKey, &res)
	if err == nil {
		response.JSON(c, http.StatusOK, res)
		return
	}

	products, pagination, err := h.usecase.ListProducts(c, &req)
	if err != nil {
		logger.Error("Failed to get products", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get products")
		return
	}

	utils.MapStruct(&res.Products, products)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
	_ = h.cache.SetWithExpiration(cacheKey, res, configs.ProductCachingTime)
}

//	@Summary		Retrieve a product by its ID
//	@Description	Fetches the details of a specific product based on the provided product ID.
//	@Tags			Products
//	@Produce		json
//	@Param			id	path	string	true	"Product ID"
//	@Success		200	{object}	response.Response	"Successfully retrieved the product"
//	@Failure		400	{object}	response.Response	"Bad Request - Invalid product ID"
//	@Failure		401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure		403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure		404	{object}	response.Response	"Not Found - Product with the specified ID not found"
//	@Failure		500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router			/api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	var res entity.Product

	cacheKey := c.Request.URL.RequestURI()
	err := h.cache.Get(cacheKey, &res)
	if err == nil {
		response.JSON(c, http.StatusOK, res)
		return
	}

	productId := c.Param("id")

	product, err := h.usecase.GetProductById(c, productId)
	if err != nil {
		logger.Error("Failed to get product detail: ", err)
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
	_ = h.cache.SetWithExpiration(cacheKey, res, configs.ProductCachingTime)
}

//	@Summary		Create a new product
//	@Description	Creates a new product based on the provided details.
//	@Tags			Products
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			name		formData	string		true	"Product Name"
//	@Param			description	formData	string		true	"Product Description"
//	@Param			image		formData	file		true	"Product Image"
//	@Param			price		formData	number		true	"Product Price (must be greater than 0)"
//	@Success		201	{object}	response.Response	"Product created successfully"
//	@Failure		400	{object}	response.Response	"Bad Request - Invalid parameters"
//	@Failure		401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure		403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure		409	{object}	response.Response	"Conflict - Code or Name already in use"
//	@Failure		500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router			/api/v1/products [post]
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

//	@Summary		Update a product
//	@Description	Update an existing product based on the provided details.
//	@Tags			Products
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id			path		string		true	"Product ID"
//	@Param			name		formData	string		false	"Product Name"
//	@Param			description	formData	string		false	"Product Description"
//	@Param			image		formData	file		false	"Product Image"
//	@Param			price		formData	number		false	"Product Price (must be greater than or equal to 0)"
//	@Success		200	{object}	response.Response	"Product updated successfully"
//	@Failure		400	{object}	response.Response	"Bad Request - Invalid parameters"
//	@Failure		401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure		403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure		404	{object}	response.Response	"Not Found - Product with the specified ID not found"
//	@Failure		500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router			/api/v1/products/{id} [put]
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

//	@Summary		Delete a product
//	@Description	Deletes an existing product by its ID.
//	@Tags			Products
//	@Produce		json
//	@Param			id	path	string	true	"Product ID"
//	@Success		200	{object}	response.Response	"Product deleted successfully"
//	@Failure		400	{object}	response.Response	"Bad Request - Invalid parameters"
//	@Failure		401	{object}	response.Response	"Unauthorized - User not authenticated"
//	@Failure		403	{object}	response.Response	"Forbidden - User does not have the required permissions"
//	@Failure		404	{object}	response.Response	"Not Found - Product with the specified ID not found"
//	@Failure		500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//	@Router			/api/v1/products/{id} [delete]
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
