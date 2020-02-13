package config

type ConfigTypes struct {
	Settings struct {
		Key     string `yaml:"key"`
		Api_Key string `yaml:"Api_Key"`
	} `yaml:"settings"`
}
