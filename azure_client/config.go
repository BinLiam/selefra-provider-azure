package azure_client

type Config struct {
	Subscriptions []string `yaml:"subscriptions,omitempty" mapstructure:"subscriptions"`
	TenantID      string   `yaml:"tenant_id,omitempty" mapstructure:"tenant_id"`
	ClientID      string   `yaml:"client_id,omitempty" mapstructure:"client_id"`
	ClientSecret  string   `yaml:"client_secret,omitempty" mapstructure:"client_secret"`

	FromFile        string `yaml:"from_file,omitempty" mapstructure:"from_file"`
	ResourceBaseURI string `yaml:"resource_base_uri,omitempty" mapstructure:"resource_base_uri"`
}
