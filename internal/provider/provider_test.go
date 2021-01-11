package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"domeneshop": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {
	if err := New("dev")().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("DOMENESHOP_TOKEN"); v == "" {
		t.Fatal("DOMENESHOP_TOKEN must be set for acceptance tests")
	}
	if v := os.Getenv("DOMENESHOP_SECRET"); v == "" {
		t.Fatal("DOMENESHOP_SECRET must be set for acceptance tests")
	}
	if v := os.Getenv("DOMENESHOP_DOMAIN"); v == "" {
		t.Fatal("DOMENESHOP_DOMAIN must be set for acceptance tests")
	}
}
