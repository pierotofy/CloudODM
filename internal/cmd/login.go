// Copyright © 2018 CloudODM Contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/OpenDroneMap/CloudODM/internal/config"
	"github.com/OpenDroneMap/CloudODM/internal/logger"
	"github.com/spf13/cobra"
)

var username string
var password string

var loginCmd = &cobra.Command{
	Use:   "login [--node default]",
	Short: "Login with a node",
	Run: func(cmd *cobra.Command, args []string) {
		config.Initialize()

		if config.CheckLogin(nodeName, username, password) != nil {
			logger.Info("Logged in")
		}
	},
}

func init() {
	loginCmd.Flags().StringVarP(&nodeName, "node", "n", "default", "Processing node to use")
	loginCmd.Flags().StringVar(&username, "username", "", "Username")
	loginCmd.Flags().StringVar(&password, "password", "", "Password")

	rootCmd.AddCommand(loginCmd)
}