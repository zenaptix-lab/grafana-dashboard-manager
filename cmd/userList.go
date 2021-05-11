package cmd

import (
	"github.com/jedib0t/go-pretty/table"
	"github.com/netsage-project/grafana-dashboard-manager/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listUserCmd = &cobra.Command{
	Use:   "list",
	Short: "list users",
	Long:  `list users`,
	Run: func(cmd *cobra.Command, args []string) {

		tableObj.AppendHeader(table.Row{"id", "login", "name", "email", "admin"})
		users := api.ListUsers(client)
		if len(users) == 0 {
			log.Info("No users found")
		} else {
			for _, user := range users {
				tableObj.AppendRow(table.Row{user.ID, user.Login, user.Name, user.Email, user.IsGrafanaAdmin})
			}
			tableObj.Render()
		}

	},
}

func init() {
	userCmd.AddCommand(listUserCmd)
}
