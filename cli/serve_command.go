package cli

import (
	"fmt"
	"lita"
	"os"

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

	if err := serveCmd.MarkFlagRequired("port"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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

	// uncomment below when release
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	if err = server.Run(fmt.Sprintf(":%s", port)); err != nil {
		return
	}
}
