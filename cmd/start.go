/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		router := gin.Default()
		router.StaticFS("/", http.Dir(args[0]))

		// router.StaticFS("/", http.Dir("mockui"))
		// router.Static("/assets", "./assets")
		// router.StaticFS("/more_static", http.Dir("my_file_system"))
		// router.StaticFile("/favicon.ico", "./resources/favicon.ico")
		// router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))

		// Listen and serve on 0.0.0.0:8080
		router.Run(":8080")
		// router.Run(":3000")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
