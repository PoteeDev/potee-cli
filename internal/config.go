package internal

type Config struct {
	Host        string `mapstructure:"host"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	AccessToken string `mapstructure:"access_token"`
}
