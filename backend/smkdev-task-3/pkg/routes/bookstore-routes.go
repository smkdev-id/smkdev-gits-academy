package routes

import (
	"BookStore/pkg/controllers"
	"BookStore/pkg/middlewares"
	"BookStore/pkg/repository"
	"BookStore/pkg/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	authRepository := repository.NewAuthRepository(db)
	bookRepository := repository.NewBookRepository(db)
	basketRepository := repository.NewBasketRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	bookService := service.NewBookService(bookRepository)
	authService := service.NewAuthService(authRepository)
	basketService := service.NewBasketService(basketRepository)
	transactionService := service.NewTransactionService(transactionRepository, basketService)

	bookController := controllers.NewBookController(bookService)
	authController := controllers.NewAuthController(authService)
	basketController := controllers.NewBasketController(basketService)
	transactionController := controllers.NewTransactionController(transactionService)

	router.POST("/register", authController.RegisterUser)
	router.POST("/login", authController.LoginUser)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middlewares.JWTAuthMiddleware("admin"))
	{
		//book
		adminRoutes.GET("/books", bookController.GetBooks)
		adminRoutes.GET("/books/:id", bookController.GetBooksByID)
		adminRoutes.POST("/books", bookController.CreateBook)
		adminRoutes.PUT("/books/:id", bookController.UpdateBook)
		adminRoutes.DELETE("/books/:id", bookController.DeleteBook)
		// Transaction management
		adminRoutes.GET("/transactions", transactionController.GetAllTransactions)
		adminRoutes.GET("/transactions/:id", transactionController.GetTransaction)
		adminRoutes.PATCH("/transactions/:id/status", transactionController.UpdateTransactionStatus)
	}

	// Routes for customer (requires customer authentication)
	customerRoutes := router.Group("/customer")
	customerRoutes.Use(middlewares.JWTAuthMiddleware("customer"))
	{
		//book
		customerRoutes.GET("/books", bookController.GetBooks)
		customerRoutes.GET("/books/:id", bookController.GetBooksByID)
		//basket
		customerRoutes.GET("/basket", basketController.GetBasket)
		customerRoutes.POST("/basket", basketController.AddToBasket)
		customerRoutes.PUT("/basket", basketController.UpdateBasketItem)
		customerRoutes.DELETE("/basket/:book_id", basketController.RemoveFromBasket)
		customerRoutes.DELETE("/basket", basketController.ClearBasket)
		// Transaction
		customerRoutes.POST("/transactions", transactionController.CreateTransaction)
		customerRoutes.GET("/transactions", transactionController.GetCustomerTransactions)
		customerRoutes.GET("/transactions/:id", transactionController.GetTransaction)
	}

	return router
}

// graceful shutdon
func RunServer(router *gin.Engine) {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server could not start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
