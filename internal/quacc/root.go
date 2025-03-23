package quacc

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/CharukaK/quacc/internal/quacc/errors"
	"github.com/charmbracelet/glamour"
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
			return handleView(args)
		}

	},
}

func handleEdit(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing arguments")
	}

	p, _ := parseArguments(args[0])

	fp := fmt.Sprintf(`%s.md`, path.Join(noteBase, p))

	if _, err := os.Stat(fp); os.IsNotExist(err) {
		fmt.Println(fp)
		file, err := os.Create(fp)

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

func handleView(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing arguments")
	}

	p, _ := parseArguments(args[0])
	fp := fmt.Sprintf(`%s.md`, path.Join(noteBase, p))
    fmt.Println(fp)
	file, err := os.Open(fp)

	if err != nil {
		return err
	}
	defer file.Close()

	buffer := bytes.NewBuffer(make([]byte, 0))

	_, err = buffer.ReadFrom(file)

	if err != nil {
		return err
	}

    rc, err := glamour.Render(buffer.String(), "dark")

    if err != nil {
        return err
    }

    fmt.Println(rc)

	return nil
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
		errors.HandleError(err)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&cmdOpts.isEdit, "edit", "e", false, "")

	// set note base
	// todo: make this to be read from a config file later

	if dir, err := setupBaseDir(); err != nil {
		errors.HandleError(err)
	} else {
		noteBase = dir
	}

	editor, _ = os.LookupEnv("EDITOR")
}
