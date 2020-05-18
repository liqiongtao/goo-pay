package gooSms

type AliyunConfig struct {
	Appid        string `yaml:"appid"`
	Secret       string `yaml:"secret"`
	SignName     string `yaml:"sign_name"`
	TemplateCode string `yaml:"template_code"`
}
