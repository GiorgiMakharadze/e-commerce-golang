package server

import (
	"github.com/GiorgiMakharadze/e-commerce-golang/internal/dto"
	"github.com/GiorgiMakharadze/e-commerce-golang/internal/services"
	"github.com/GiorgiMakharadze/e-commerce-golang/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data", err)
		return
	}

	authService := services.NewAuthService(s.db, s.config)
	response, err := authService.Register(&req)
	if err != nil {
		utils.BadRequestResponse(ctx, "Registration failed", err)
		return
	}

	utils.CreatedResponse(ctx, "User registered successfully", response)
}

func (s *Server) login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data", err)
		return
	}

	authService := services.NewAuthService(s.db, s.config)
	response, err := authService.Login(&req)
	if err != nil {
		utils.BadRequestResponse(ctx, "Login failed", err)
		return
	}

	utils.SuccessResponse(ctx, "Login successfully", response)
}

func (s *Server) refreshToken(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data", err)
		return
	}

	authService := services.NewAuthService(s.db, s.config)
	response, err := authService.RefreshToken(&req)
	if err != nil {
		utils.BadRequestResponse(ctx, "Login failed", err)
		return
	}

	utils.SuccessResponse(ctx, "Token refreshed successfully", response)
}

func (s *Server) logout(ctx *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data", err)
		return
	}

	authService := services.NewAuthService(s.db, s.config)
	err := authService.Logout(req.RefreshToken)
	if err != nil {
		utils.BadRequestResponse(ctx, "Login failed", err)
		return
	}

	utils.SuccessResponse(ctx, "Logout successfully", nil)
}

func (s *Server) getProfile(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")
	userService := services.NewUserService(s.db)
	profile, err := userService.GetProfile(userID)
	if err != nil {
		utils.NotFoundResponse(ctx, "User not found")
		return
	}

	utils.SuccessResponse(ctx, "Profile retrieved successfully", profile)
}

func (s *Server) updateProfile(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	var req dto.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(ctx, "Invalid request data", err)
		return
	}

	userService := services.NewUserService(s.db)
	profile, err := userService.UpdateProfile(userID, &req)
	if err != nil {
		utils.InternalServerErrorResponse(ctx, "Failed to update profile", err)
		return
	}
	utils.SuccessResponse(ctx, "Profile updated successfully", profile)
}
