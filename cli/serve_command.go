package cli

import (
	"fmt"
	"lita"
	"lita/endpoint"
	"lita/service"
	"lita/transport"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	serverPort string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: fmt.Sprintf("Start %s server", lita.FirstUpper(lita.Name)),
	Long:  fmt.Sprintf("Start %s server.", lita.FirstUpper(lita.Name)),
	Run: func(cmd *cobra.Command, args []string) {
		serve(serverPort)
	},
}

func init() {
	serveCmd.Flags().StringVarP(&serverPort, "port", "p", lita.Port,
		fmt.Sprintf("optional %s server port, default %s", lita.Name, lita.Port))
}

func serve(optionalPort ...string) {
	var err error
	defer func() {
		fmt.Println(lita.NewError(lita.ErrServing, "failed to serve", err))
	}()

	port := lita.Port
	if len(optionalPort) != 0 {
		port = optionalPort[0]
	}

	// uncomment below in production environment
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	baseService := service.NewBaseService()
	{
		group := server.Group("")
		transport.NewBaseHTTPHandler(endpoint.NewBaseEndpoint(baseService), group)
	}

	if err = server.Run(fmt.Sprintf(":%s", port)); err != nil {
		return
	}
}
