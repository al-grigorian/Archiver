package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands

// Создание команды
var rootCmd = &cobra.Command{
	Use:   "Archiver",
	Short: "A brief description of your application", // Краткое пояснение, которое пользователь видит при запуске программы
	Long:  "",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleErr(err)
	}
}

// Обработка ошибки
func handleErr(err error) {
	_, _ = fmt.Fprint(os.Stderr, err)
	os.Exit(1)
}

func init() {

}
