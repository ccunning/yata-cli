package cmd

import (
	"github.com/tuxagon/yata-cli/task"
	"github.com/urfave/cli"
)

// List returns the list of tasks/todos that have been recorded
func List(ctx *cli.Context) error {
	m := task.NewFileManager()

	return nil
}
