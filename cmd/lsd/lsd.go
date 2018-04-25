package lsd

import (
	"os"

	"github.com/ncw/rclone/cmd"
	"github.com/ncw/rclone/cmd/ls/lshelp"
	"github.com/ncw/rclone/fs"
	"github.com/ncw/rclone/fs/operations"
	"github.com/spf13/cobra"
)

var (
	recurse bool
)

func init() {
	cmd.Root.AddCommand(commandDefintion)
	commandDefintion.Flags().BoolVarP(&recurse, "recursive", "R", false, "Recurse into the listing.")
}

var commandDefintion = &cobra.Command{
	Use:   "lsd remote:path",
	Short: `List all directories/containers/buckets in the path.`,
	Long: `
Lists the directories in the source path to standard output. Does not
recurse by default.  Use the -R flag to recurse.
` + lshelp.Help,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 1, command, args)
		if recurse {
			fs.Config.MaxDepth = 0
		}
		fsrc := cmd.NewFsSrc(args)
		cmd.Run(false, false, command, func() error {
			return operations.ListDir(fsrc, os.Stdout)
		})
	},
}
