package cmd

import (
	"fmt"

	screenshotcopy "github.com/jm199seo/screenshot_copy/internal/screenshot_copy"
	"github.com/spf13/cobra"
)

func runScreenshotCopy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "screenshot-copy",
		Short: "sc",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			cli, cleanup, err := screenshotcopy.NewClient(ctx)
			if err != nil {
				return err
			}

			err = cli.GetConfig()
			if err != nil {
				return err
			}

			err = cli.StartWatcher(ctx)
			if err != nil {
				return err
			}
			err = cli.WatchDir(ctx)
			if err != nil {
				return err
			}
			<-ctx.Done()
			fmt.Println("done!")
			cleanup()

			return nil
		},
	}
	return cmd
}
