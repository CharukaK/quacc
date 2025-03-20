package quacc

import (
	"fmt"

	"github.com/spf13/cobra"
)

type RootCmdOpts struct {
	isEdit bool
}

var cmdOpts = &RootCmdOpts{}

var rootCmd = &cobra.Command{
	Use:   `quacc`,
	Short: `quick access notes`,
	Long:  `Quick Access Notes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("no topic provided")
		}

		if cmdOpts.isEdit {
			// handle  editing
			return handleEdit(args)
		} else {
			// handle viewing
		}

		return nil
	},
}

func handleEdit(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing arguments")
	}

	return nil
}


func RunCmd() {
	if err := rootCmd.Root().Execute(); err != nil {
		// TODO: Handle error centrally
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&cmdOpts.isEdit, "edit", "e", false, "")
}
