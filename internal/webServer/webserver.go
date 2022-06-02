package webserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Listen() {

	router := gin.Default()

	router.GET("/", Default)
	router.POST("/addmonitor", AddMonitor)
	router.StaticFS("/static", gin.Dir("./internal/webApp", false))

	srv := &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: router,
	}

	fmt.Println(fmt.Sprintf("Server is listening to %s", srv.Addr))

	var err error
	//err = srv.ListenAndServe()
	err = srv.ListenAndServeTLS(os.Getenv("CERT"), os.Getenv("CERT_KEY"))
	// service connections
	if err != nil && err != http.ErrServerClosed {
		fmt.Println(fmt.Sprintf("listen: %s\n", err))
	} else {
		fmt.Println(fmt.Sprintf("Server is listening to %s", srv.Addr))
	}

}

func Default(c *gin.Context) {
	c.Writer.WriteHeader(204)
}
func AddMonitor(c *gin.Context) {
	c.Writer.WriteHeader(204)
}
