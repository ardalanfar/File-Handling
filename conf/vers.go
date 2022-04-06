package conf

type Config struct {
	Address_text_file string
	Web_address       string
	Allowed_domain    string
}

//Set configs
func GetConfig() *Config {
	return &Config{
		Address_text_file: "docs/text",
		Web_address:       "https://en.wikipedia.org/wiki/Web_scraping",
		Allowed_domain:    "en.wikipedia.org",
	}
}
