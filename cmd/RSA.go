/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	// "RootCmd"

	"fmt"
	helper "rsa/helper"

	"github.com/spf13/cobra"
)

// RSACmd represents the RSA command
var RSACmd = &cobra.Command{
	Use:   "RSA",
	Short: "This RSA algorithm converts the string to encrypted message and we can decrypt using private key",
	Long:  `This RSA algorithm converts the string to encrypted message and we can decrypt using private key`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("hello")
		res := helper.RSA(args[0])
		fmt.Println("The input passed is ", args[0])
		fmt.Println("the result we got is", res)
	},
}

func init() {
	RootCmd.AddCommand(RSACmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// RSACmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// RSACmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
