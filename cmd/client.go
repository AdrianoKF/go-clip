/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/AdrianoKF/go-clip/internal/util"
	"github.com/AdrianoKF/go-clip/pkg/model"
	n "github.com/AdrianoKF/go-clip/pkg/net"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		hostname, _ := os.Hostname()
		port, _ := cmd.Flags().GetInt("port")
		secret, _ := cmd.Flags().GetString("secret")

		util.Logger.Infof("client starting, port=%d, hostname=%s", port, hostname)

		addr := net.UDPAddr{
			IP:   net.IPv4(239, 255, 90, 90),
			Port: port,
		}

		client := n.NewClient(addr, secret, nil)

		util.Logger.Info("Watching for clipboard events")
		ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
		for data := range ch {
			util.Logger.Info("Received clipboard event: ", string(data))
			bytes := []byte(data)

			client.SendEvent(model.ClipboardUpdated{
				Source:      hostname,
				Content:     bytes,
				ContentType: http.DetectContentType(bytes),
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
