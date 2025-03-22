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

//		@Summary	 Signup a new user
//	 @Description Registers a new user with the provided details. Returns a sign-in response upon successful registration.
//		@Tags		 Auth
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User successfully registered"
//		@Failure	 400	{object}	response.Response	"Invalid user input"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/signup [post]
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

//		@Summary	 Signin a user
//	 @Description Authenticates the user based on the provided credentials and returns a sign-in response if successful.
//		@Tags		 Auth
//		@Produce	 json
//		@Param		 _	body	dto.SignInRequest	  true	"Body"
//		@Success	 200	{object}	response.Response	"Successfully signed in"
//		@Failure	 401	{object}	response.Response	"Unauthorized - Invalid credentials"
//		@Failure	 404	{object}	response.Response	"Not Found - User not found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/signin [post]
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

//		@Summary	 Signout a user
//	 @Description Registers a new user with the provided details. Returns a sign-in response upon successful registration.
//		@Tags		 Auth
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User successfully logout"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/auth/signout [post]
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

//		@Summary	 Get users
//	 @Description Get list of users.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Get users successfully"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users [get]
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

//		@Summary	 Get user
//	 @Description Get user detail.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"Get user successfully"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/{id} [get]
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

//		@Summary	 Delete user
//	 @Description Delete user from database.
//		@Tags		 Users
//		@Produce	 json
//		@Success	 200	{object}	response.Response	"User successfully deleted"
//		@Failure	 404	{object}	response.Response	"Not Found"
//		@Failure	 500	{object}	response.Response	"Internal Server Error - An error occurred while processing the request"
//		@Router		 /api/v1/users/{id} [delete]
func (h *AuthHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	err := h.usecase.DeleteUser(c, userId)
	if err != nil {
		logger.Error("Failed to delete ", err)
		response.Error(c, http.StatusInternalServerError, err, "Failed to sign up")
		return
	}

	response.JSON(c, http.StatusOK, true)
}
