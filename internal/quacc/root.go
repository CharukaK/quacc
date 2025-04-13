package quacc

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/CharukaK/quacc/internal/quacc/cmdargs"
	"github.com/CharukaK/quacc/internal/quacc/errors"
	"github.com/CharukaK/quacc/internal/quacc/fileutils"
	"github.com/CharukaK/quacc/internal/quacc/logger"
	"github.com/CharukaK/quacc/internal/quacc/render"
	"github.com/spf13/cobra"
)

type RootCmdOpts struct {
	isEdit bool
}

var cmdOpts = &RootCmdOpts{}
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
	ValidArgsFunction: argCompletionFunc,
}

func argCompletionFunc(
	cmd *cobra.Command,
	args []string,
	toComplete string,
) (comps []cobra.Completion, directive cobra.ShellCompDirective) {
	directive = cobra.ShellCompDirectiveNoSpace
	comps, _ = fileutils.GetTopicListForCompletion(toComplete)
	logger.Info("to complete >>>", toComplete)
	logger.Info("dir list >>>", fmt.Sprint("%v", comps))
	return
}

func handleEdit(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing arguments")
	}

	p, _ := cmdargs.ParseArguments(args[0])

	err := fileutils.CreateTopicIfNotExists(p)
	if err != nil {
		return err
	}

	return openEditor(fileutils.GenTopicFilePath(p))
}

func openEditor(filePath string) error {
	cmd := exec.Command(editor, filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func handleView(opts []string) error {
	if len(opts) == 0 {
		return fmt.Errorf("missing arguments")
	}

	p, _ := cmdargs.ParseArguments(opts[0])
	fp, err := fileutils.ResolveTopicFilePath(p)
	if err != nil {
		return err
	}

	content, err := fileutils.GetFileContent(fp)
	if err != nil {
		return err
	}

	return render.RenderNoteContent(content)
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
	if err := fileutils.SetupBaseDir(); err != nil {
		errors.HandleError(err)
	}

	editor, _ = os.LookupEnv("EDITOR")
}
