package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "potee-cli",
		Short: "Cli tool to manage Potee Platform",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.potee.yaml)")
	rootCmd.PersistentFlags().StringP("username", "u", "admin", "platform admin username")
	rootCmd.PersistentFlags().StringP("host", "", "potee.local", "platform url name")
	rootCmd.PersistentFlags().StringP("password", "p", "admin", "password")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.SetDefault("username", "admin")
	viper.SetDefault("host", "potee.local")
	// viper.SetDefault("license", "apache")

	// rootCmd.AddCommand(initCmd)
	// rootCmd.AddCommand(initCmd)

}

type Config struct {
	Host        string `mapstructure:"host"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	AccessToken string `mapstructure:"access_token"`
}

var config = &Config{}

func initConfig() {
	//viper.AutomaticEnv()
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".potee" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".potee")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}
}
