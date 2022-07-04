/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"net"
	"net/http"

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
		util.Logger.Info("client starting")

		addr := net.UDPAddr{
			IP:   net.IPv4(239, 255, 90, 90),
			Port: 9090,
		}
		client := n.NewClient(addr, nil)

		util.Logger.Info("Watching for clipboard events")
		ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
		for data := range ch {
			util.Logger.Info("Received clipboard event: ", data)
			bytes := []byte(data)

			client.SendEvent(model.ClipboardUpdated{
				Content:     bytes,
				ContentType: http.DetectContentType(bytes),
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
