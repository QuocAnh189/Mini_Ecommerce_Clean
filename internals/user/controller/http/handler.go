package http

import (
	"ecommerce_clean/internals/user/controller/dto"
	"ecommerce_clean/internals/user/usecase"
	"ecommerce_clean/pkgs/logger"
	"ecommerce_clean/pkgs/response"
	"ecommerce_clean/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	usecase usecase.IUserUseCase
}

func NewAuthHandler(usecase usecase.IUserUseCase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}

//	@Summary		User Sign-Up
//	@Description	Registers a new user with the provided details and returns access tokens along with user info if successful.
//	@Tags			Auth
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			email		formData	string					true	"User email (must be unique)"
//	@Param			name		formData	string					true	"User name (must be unique)"
//	@Param			password	formData	string					true	"User password"
//	@Param			avatar		formData	file					false	"User avatar file"
//	@Success		200			{object}	dto.SignUpResponse		"User successfully registered"
//	@Failure		400			{object}	response.Response		"Bad Request - Invalid parameters"
//	@Failure		409			{object}	response.Response		"Conflict - Email or Name already in use"
//	@Failure		500			{object}	response.Response		"Internal Server Error - Failed to sign up"
//	@Router			/auth/signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	accessToken, refreshToken, user, err := h.usecase.SignUp(c, &req)
	if err != nil {
		logger.Error("Failed to sign up ", err)
		switch utils.ExtractConstraintName(err) {
		case "unique_user_email":
			response.Error(c, http.StatusConflict, err, "Email already in use")
		case "unique_user_name":
			response.Error(c, http.StatusConflict, err, "Name already in use")
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		}
		return
	}

	var res dto.SignUpResponse
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	utils.MapStruct(&res.User, user)

	response.JSON(c, http.StatusOK, res)
}

//	@Summary		User Sign-In
//	@Description	Authenticates the user based on the provided credentials and returns access tokens and user info if successful.
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.SignInRequest	true	"User sign-in request"
//	@Success		200		{object}	dto.SignInResponse	"Successfully signed in"
//	@Failure		400		{object}	response.Response	"Bad Request - Invalid parameters"
//	@Failure		409		{object}	response.Response	"Conflict - Wrong password or Email does not exist"
//	@Failure		500		{object}	response.Response	"Internal Server Error - Failed to sign in"
//	@Router			/auth/signin [post]
func (h *AuthHandler) SignIn(c *gin.Context) {
	var req dto.SignInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error("Failed to get body ", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}
	accessToken, refreshToken, user, err := h.usecase.SignIn(c, &req)

	if err != nil {
		logger.Error("Failed to sign up ", err)
		switch err.Error() {
		case "wrong password":
			response.Error(c, http.StatusConflict, err, "Wrong password")
		case "record not found":
			response.Error(c, http.StatusConflict, err, "Email does not exist")
		default:
			response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		}
		return
	}

	var res dto.SignInResponse
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	utils.MapStruct(&res.User, user)
	response.JSON(c, http.StatusOK, res)
}

//	@Summary		User Sign-Out
//	@Description	Logs out the authenticated user by invalidating the current session token.
//	@Tags			Auth
//	@Produce		json
//	@Success		200				{object}	response.Response	"User successfully logged out"
//	@Failure		400				{object}	response.Response	"Bad Request - Missing Authorization header"
//	@Failure		401				{object}	response.Response	"Unauthorized - Invalid or missing user ID"
//	@Failure		500				{object}	response.Response	"Internal Server Error - Failed to sign out"
//	@Router			/auth/signout [post]
//	@Security		ApiKeyAuth
func (h *AuthHandler) SignOut(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		response.Error(c, http.StatusBadRequest, nil, "Missing Authorization header")
		return
	}

	jit := c.GetString("jit")

	userID, exists := c.Get("userId")
	if !exists {
		response.Error(c, http.StatusNotFound, nil, "Unauthorized")
		return
	}

	err := h.usecase.SignOut(c, userID.(string), jit)
	if err != nil {
		logger.Error("Failed to sign out", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to sign out")
		return
	}

	response.JSON(c, http.StatusOK, "Logout successfully")
}

//	@Summary		Get Users List
//	@Description	Retrieves a paginated list of users based on search criteria.
//	@Tags			Users
//	@Produce		json
//	@Param			search		query	string	false	"Search keyword for filtering users"
//	@Param			page		query	int		false	"Page number for pagination"
//	@Param			size		query	int		false	"Number of users per page"
//	@Param			order_by	query	string	false	"Column name to sort by"
//	@Param			order_desc	query	bool	false	"Sort in descending order (true/false)"
//	@Param			take_all	query	bool	false	"Retrieve all users without pagination"
//	@Success		200			{object}	response.Response	"Successfully retrieved users list"
//	@Failure		400			{object}	response.Response	"Bad Request - Invalid query parameters"
//	@Failure		500			{object}	response.Response	"Internal Server Error - Failed to get users"
//	@Router			/users [get]
//	@Security		ApiKeyAuth
func (h *AuthHandler) GetUsers(c *gin.Context) {
	var req dto.ListUserRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("Failed to get query", err)
		response.Error(c, http.StatusBadRequest, err, "Invalid parameters")
		return
	}

	users, pagination, err := h.usecase.ListUsers(c, &req)
	if err != nil {
		logger.Error("Failed to get users", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to get users")
		return
	}

	var res dto.ListUserResponse
	utils.MapStruct(&res.Users, users)
	res.Pagination = pagination
	response.JSON(c, http.StatusOK, res)
}

//	@Summary		Get User Detail
//	@Description	Retrieves detailed information of a specific user by ID.
//	@Tags			Users
//	@Produce		json
//	@Param			id		path	string	true	"User ID"
//	@Success		200		{object}	response.Response	"Successfully retrieved user details"
//	@Failure		400		{object}	response.Response	"Bad Request - Invalid user ID format"
//	@Failure		404		{object}	response.Response	"Not Found - User does not exist"
//	@Failure		500		{object}	response.Response	"Internal Server Error - Failed to retrieve user details"
//	@Router			/users/{id} [get]
//	@Security		ApiKeyAuth
func (h *AuthHandler) GetUser(c *gin.Context) {
	userId := c.Param("id")
	user, err := h.usecase.GetUserById(c, userId)
	if err != nil {
		logger.Error("Failed to get user detail: ", err)
		response.Error(c, http.StatusInternalServerError, err, "Something went wrong")
		return
	}

	response.JSON(c, http.StatusOK, user)
}

//	@Summary		Delete User
//	@Description	Deletes a user from the database by ID.
//	@Tags			Users
//	@Produce		json
//	@Param			id		path	string	true	"User ID"
//	@Success		200		{object}	response.Response	"User successfully deleted"
//	@Failure		400		{object}	response.Response	"Bad Request - Invalid user ID format"
//	@Failure		404		{object}	response.Response	"Not Found - User does not exist"
//	@Failure		500		{object}	response.Response	"Internal Server Error - Failed to delete user"
//	@Router			/users/{id} [delete]
//	@Security		ApiKeyAuth
func (h *AuthHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	err := h.usecase.DeleteUser(c, userId)
	if err != nil {
		logger.Error("Failed to delete ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		return
	}

	response.JSON(c, http.StatusOK, "Delete user successfully")
}
