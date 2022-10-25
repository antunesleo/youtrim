/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// trimCmd represents the trim command
var trimCmd = &cobra.Command{
	Use:   "trim",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// url, _ := cmd.Flags().GetString("url")
		// if url == "" {
		// 	fmt.Println("You must provide an URL flag --url string")
		// }
	},
}

func init() {
	rootCmd.AddCommand(trimCmd)
	trimCmd.PersistentFlags().String("url", "", "")
	trimCmd.PersistentFlags().String("start", "", "")
	trimCmd.PersistentFlags().String("end", "", "")
	trimCmd.MarkPersistentFlagRequired("url")
	trimCmd.MarkPersistentFlagRequired("start")
	trimCmd.MarkPersistentFlagRequired("end")
}
