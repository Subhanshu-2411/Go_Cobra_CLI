/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Go_Cobra_CLI/util"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

// visualStudioCmd represents the visualStudio command
var visualStudioCmd = &cobra.Command{
	Use:   "visualStudio <Zipped File Name>",
	Short: "Open a Zip Folder in Visual Studio",
	Long: `It opens a Zipped Folder Present in your working directory 
in Visual Studio with simple Arguments. Do Note to install Visual
Studio on Your System`,
	//Args:                  cobra.ExactArgs(1),
	Args: func(cmd *cobra.Command, args []string) error {
		if File == "" && len(args) < 1 {
			return errors.New("Accepts 1 argument received 0")
		}
		return nil
	},
	DisableFlagsInUseLine: true,
	Example:               `Go_Cobra_CLI visualStudio hello.zip (or File Path)`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string

		if File != "" {
			argument = File
		} else {
			argument = args[0]
		}

		fileExists, err := util.FileExists(argument)

		if err != nil {
			fmt.Println(err.Error())
		}

		if fileExists {
			fileName, err = filepath.Abs(argument)

			if err != nil {
				fmt.Println(err.Error())
			}

		} else {
			fmt.Printf("File %v doen not exist", argument)
			return
		}

		wd, err := os.Getwd()

		if err != nil {
			fmt.Println(err.Error())
		}

		util.Unzip(fileName, wd)

		os.Chdir(util.FilenameWithoutExtension(fileName))

		wd, err = os.Getwd()

		if err != nil {
			fmt.Println(err.Error())
		}

		commandCode := exec.Command("visualStudio", wd)
		err = commandCode.Run()

		if err != nil {
			fmt.Println("Visual Studio Executable File not found in %PATH%")
		}
	},
}

func init() {
	rootCmd.AddCommand(visualStudioCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// visualStudioCmd.PersistentFlags().String("foo", "", "A help for foo")
	visualStudioCmd.PersistentFlags().StringVarP(&File, "file", "f", "", "File Name to be Open")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// visualStudioCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
