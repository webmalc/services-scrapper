package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger          ErrorLogger
	rootCmd         *cobra.Command
	config          *Config
	scrappersRunner Runner
}

// server runs server.
func (r *CommandRouter) scrap(cmd *cobra.Command, args []string) {
	r.scrappersRunner.Run(args)
}

// Run the router.
func (r *CommandRouter) Run() {
	r.rootCmd.AddCommand(
		&cobra.Command{
			Use:       "scrap",
			Short:     "Run the scrappers",
			Run:       r.scrap,
			ValidArgs: r.config.scrappers,
			Args:      cobra.OnlyValidArgs,
		},
	)
	err := r.rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(log ErrorLogger, r Runner) CommandRouter {
	return CommandRouter{
		config:          NewConfig(),
		logger:          log,
		rootCmd:         &cobra.Command{Use: "ss.app"},
		scrappersRunner: r,
	}
}
