package homebrew

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceHomebrewPackage() *schema.Resource {
	return &schema.Resource{
		Create: resourceHomebrewPackageCreate,
		Read:   resourceHomebrewPackageRead,
		Delete: resourceHomebrewPackageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceHomebrewPackageCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*PackageEndpoint)
	name := d.Get("name").(string)

	params := &Package{
		Name: name,
	}

	packageDesc, err := client.Install(params)

	if err == nil {
		log.Printf("[INFO] Install package: %s", name)
		d.SetId(packageDesc.Name)
	}

	return resourceHomebrewPackageRead(d, m)
}

func resourceHomebrewPackageRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*PackageEndpoint)
	packageDesc, err := client.Get(d.Id())

	if err != nil {
		return err
	}

	d.Set("name", packageDesc.Name)

	return nil
}

func resourceHomebrewPackageDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*PackageEndpoint)
	name := d.Get("name").(string)

	params := &Package{
		Name: name,
	}

	_, err := client.Uninstall(params)

	if err == nil {
		log.Printf("[INFO] Uninstall package: %s", name)
		d.SetId("")
	}

	return err
}
