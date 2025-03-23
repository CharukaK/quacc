package quacc

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

type RootCmdOpts struct {
	isEdit bool
}

var cmdOpts = &RootCmdOpts{}
var noteBase = ""
var editor = ""

var rootCmd = &cobra.Command{
	Use:   `quacc`,
	Short: `quick access notes`,
	Long:  `Quick Access Notes`,
	Args:  cobra.ExactArgs(1),
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

	p, _ := parseArguments(args[0])

	fp := path.Join(noteBase, p)

	if _, err := os.Stat(p); os.IsNotExist(err) {
		nFile := fmt.Sprintf(`%s.md`, fp)
		fmt.Println(nFile)
		file, err := os.Create(nFile)

		if err != nil {
			return err
		}

		file.Close()
	}

	cmd := exec.Command(editor, fp)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func parseArguments(input string) (path string, searchQuery []string) {
	segments := strings.Split(input, "~")

	if len(segments) > 0 {
		path = segments[0]
	}

	if len(segments) > 1 {
		searchQuery = strings.Split(segments[1], "+")
	}

	return
}

func RunCmd() {
	if err := rootCmd.Root().Execute(); err != nil {
		// TODO: Handle error centrally
		fmt.Println(err.Error())
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&cmdOpts.isEdit, "edit", "e", false, "")

	// set note base
	// todo: make this to be read from a config file later

	if dir, err := setupBaseDir(); err != nil {
		// todo: handle error
	} else {
		noteBase = dir
	}

	editor, _ = os.LookupEnv("EDITOR")
}
