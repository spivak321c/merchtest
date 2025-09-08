package api

    import (
       "api-customer-merchant/internal/api/handlers"
       "api-customer-merchant/internal/db/repositories"
       "api-customer-merchant/internal/middleware"
        "api-customer-merchant/internal/domain/merchant"

        "github.com/gin-gonic/gin"
    )

/*
    func RegisterRoutes(r *gin.Engine) {
        merchant := r.Group("/merchant")
        {
           //userRepo := repositories.NewUserRepository()
            // merchantRepo := repositories.NewMerchantRepository()
            appRepo := repositories.NewMerchantApplicationRepository()
            repo := repositories.NewMerchantRepository()
            service := merchant.NewMerchantService(appRepo, repo)
            h := handlers.NewMerchantHandler(service)

            // authHandler := handlers.NewAuthHandler(merchantRepo)
            // merchant.POST("/submitApplication", authHandler.Register)
            // merchant.POST("/login", authHandler.Login)
            //merchant.GET("/auth/google", authHandler.GoogleAuth)
            //merchant.GET("/auth/google/callback", authHandler.GoogleCallback)
            merchant.POST("/apply",  h.Apply)
            merchant.GET("/application/:id",  h.GetApplication)
            
            
            // Merchant account access (once approved by admin via Express API)
            //protected.POST("/me",  h.GetMyMerchant)
    
           

            protected := merchant.Group("/")
            protected.Use(middleware.AuthMiddleware("merchant"))
            protected.POST("/me",  h.GetMyMerchant)
            //protected.POST("/logout", authHandler.Logout)
            
        }
    }
*/

func RegisterRoutes(r *gin.Engine) {
    appRepo := repositories.NewMerchantApplicationRepository()
    repo := repositories.NewMerchantRepository()
    service := merchant.NewMerchantService(appRepo, repo)
    h := handlers.NewMerchantHandler(service)
    
    
    mg := r.Group("/merchant")
    {
    // Application submission and status
    mg.POST("/apply",  h.Apply)
    mg.GET("/application/:id",  h.GetApplication)
    
    
    // Merchant account access (once approved by admin via Express API)
    mg.GET("/me", middleware.AuthMiddleware("merchant"), h.GetMyMerchant)
    }
    }