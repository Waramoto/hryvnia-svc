package config

type EmailConfig struct {
	From     string `fig:"from,required"`
	Password string `fig:"password,required"`
	Host     string `fig:"host,required"`
	Port     string `fig:"port,required"`
	Identity string `fig:"identity"`
}
