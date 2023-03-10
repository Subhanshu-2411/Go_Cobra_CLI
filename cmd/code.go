/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"Go_Cobra_CLI/util"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code <Zipped File Name>",
	Short: "Open a Zip Folder in Visual Studio Code",
	Long: `It opens a Zipped Folder Present in your working directory 
in Visual Studio Code with simple Arguments. Do Note to install Visual
Studio Code on Your System`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Example:               `Go_Cobra_CLI code hello.zip (or File Path)`,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string

		argument = args[0]

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

		commandCode := exec.Command("code", wd)
		err = commandCode.Run()

		if err != nil {
			fmt.Println("VS Code Executable File not found in %PATH%")
		}
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
