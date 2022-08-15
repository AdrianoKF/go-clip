/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"net"
	"strings"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"

	"github.com/AdrianoKF/go-clip/internal/util"
	"github.com/AdrianoKF/go-clip/pkg/model"
	n "github.com/AdrianoKF/go-clip/pkg/net"
)

func PrintMessage(_ *net.UDPAddr, _ int, buf []byte) {
	var ev model.ClipboardUpdated
	err := json.NewDecoder(bytes.NewReader(buf)).Decode(&ev)
	if err != nil {
		util.Logger.Error(err)
	}
	if strings.HasPrefix(ev.ContentType, "text/") {
		util.Logger.Info(string(ev.Content))
		clipboard.Write(clipboard.FmtText, ev.Content)
	} else {
		util.Logger.Info(ev)
	}

}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(_ *cobra.Command, _ []string) {
		addr := net.UDPAddr{
			IP:   net.IPv4(239, 255, 90, 90),
			Port: 9090,
		}
		server := n.NewServer(addr, PrintMessage)
		util.Logger.Info("Starting server: %+v", server)
		server.Listen()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// server.goCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// server.goCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
