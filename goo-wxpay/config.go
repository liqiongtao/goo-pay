package gooWXPay

type Config struct {
	MchId           string `yaml:"mch_id"`
	ApiKey          string `yaml:"api_key"`
	NotifyUrl       string `yaml:"notify_url"`
	NotifyRefundUrl string `yaml:"notify_refund_url"`
}
