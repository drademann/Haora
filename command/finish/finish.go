package finish

import (
	"github.com/drademann/haora/app/data"
	"github.com/drademann/haora/command/internal/parsing"
	"github.com/drademann/haora/command/root"
	"github.com/spf13/cobra"
)

var (
	endFlag string
)

func init() {
	Command.Flags().StringVarP(&endFlag, "end", "e", "", "The finish time, like 17:00, for the day")
	root.Command.AddCommand(Command)
}

var Command = &cobra.Command{
	Use:   "finish",
	Short: "Mark the day as done",
	RunE: func(cmd *cobra.Command, args []string) error {
		time, _, err := parsing.Time(endFlag, args)
		if err != nil {
			return err
		}
		day := data.State.WorkingDay()
		day.Finish(time)
		return nil
	},
	PostRun: func(cmd *cobra.Command, args []string) { // reset flag so tests can rerun!
		endFlag = ""
	},
}
