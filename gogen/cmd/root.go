package cmd

import (
	"sync"

	"github.com/spf13/cobra"
)

var (
	rootCmd *cobra.Command
	once    sync.Once
)

// NewCmd new cobra
func NewCmd() *cobra.Command {
	once.Do(func() {
		rootCmd = &cobra.Command{
			Use:   "gogen",
			Short: "ChengFeng code generater artifact by fengtianru117",
			Long:  "This command will generate code by customed golang struct",
			RunE:  nil,
		}
	})
	return rootCmd
}
