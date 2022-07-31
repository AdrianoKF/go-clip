/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.design/x/clipboard"

	"github.com/AdrianoKF/go-clip/internal/util"
	"github.com/AdrianoKF/go-clip/pkg/model"
	n "github.com/AdrianoKF/go-clip/pkg/net"
)

var hostname, _ = os.Hostname()

func PrintMessage(_ *net.UDPAddr, _ int, buf []byte) {
	var ev model.ClipboardUpdated
	err := json.NewDecoder(bytes.NewReader(buf)).Decode(&ev)
	if err != nil {
		util.Logger.Error(err)
	}

	if ev.Source == hostname {
		util.Logger.Info("Ignoring event from myself")
		return
	}

	if strings.HasPrefix(ev.ContentType, "text/") {
		util.Logger.Info(strings.TrimSpace(string(ev.Content)))
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
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		secret, _ := cmd.Flags().GetString("secret")

		addr := net.UDPAddr{
			IP:   net.IPv4(239, 255, 90, 90),
			Port: port,
		}
		server := n.NewServer(addr, secret, PrintMessage)
		util.Logger.Infof("Starting server, address=%s:%d, hostname=%s", addr.IP, addr.Port, hostname)
		server.Listen()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
