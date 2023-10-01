package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PoteeDev/potee-cli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ModuleArgs struct {
	Admin string
}

var config = &internal.Config{}

func initConfig() {
	//viper.AutomaticEnv()
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".potee" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".potee")

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}
}

func init() {
	initConfig()
}

type ModuleResponse struct {
	Msg     string `json:"msg,omitempty"`
	Failed  bool   `json:"failed,omitempty"`
	Changed bool   `json:"changed,omitempty"`
}

func ReturnResponse(response ModuleResponse) {
	val, _ := json.Marshal(response)
	fmt.Println(string(val))
}

func main() {
	moduleArgs := ModuleArgs{}
	args, _ := os.ReadFile(os.Args[1])
	json.Unmarshal(args, &moduleArgs)

	response := ModuleResponse{}

	switch moduleArgs.Admin {
	case "teams":
		msg, _ := json.Marshal(config.ShowTeams())
		response.Msg = string(msg)
		ReturnResponse(response)
	}
	os.Exit(0)
}
