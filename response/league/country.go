package league

type Country struct {
	Name string `mapstructure:"name"`
	Code string `mapstructure:"code"`
	Flag string `mapstructure:"flag"`
}