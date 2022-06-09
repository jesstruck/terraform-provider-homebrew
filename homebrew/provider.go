package homebrew

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:     schema.TypeString,
				Required: true,
			},
			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/usr/local/bin/brew",
			},
			"login": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("USER", ""),
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"homebrew_package": resourceHomebrewPackage(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Host: d.Get("host").(string),
	}

	if login, ok := d.GetOk("login"); ok {
		config.Login = login.(string)
	}

	if publicKey, ok := d.GetOk("public_key"); ok {
		config.PubKey = publicKey.(string)
	}

	if path, ok := d.GetOk("path"); ok {
		config.HomebrewBinaryPath = path.(string)
	}

	log.Println("[INFO] Initializing Homebrew client")
	return config.Client()
}
