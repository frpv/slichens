/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"slichens/pkg/attenuation"
)

// attenuationCmd represents the attenuation command
var attenuationCmd = &cobra.Command{
	Use:   "attenuation",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := cmd.Flags().GetString("outfile")
		in, _ := cmd.Flags().GetString("infile")
		if out != "" && in != "" {
			attenuation.Attenuation(out, in, Verbose, Freq, true)
		} else {
			fmt.Println("survey files name requiered")
		}
	},
}

func init() {
	rootCmd.AddCommand(attenuationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// attenuationCmd.PersistentFlags().String("foo", "", "A help for foo")
	attenuationCmd.PersistentFlags().String("outfile", "", "Outdoor siretta filename Lxxxxx.csv")
	attenuationCmd.PersistentFlags().String("infile", "", "Indoor siretta filename Lxxxxx.csv")
	attenuationCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose statistic output")
	attenuationCmd.PersistentFlags().BoolVarP(&Freq, "band", "b", false, "Sorted by frequency band")
	attenuationCmd.PersistentFlags().BoolVarP(&Sample, "exclude", "s", false, "Sorted by frequency band")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// attenuationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}