package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func UUIDGenerator() *schema.Resource {
	return &schema.Resource{
		Create: func(r *schema.ResourceData, m any) error {
			uc, ok := r.Get("uuid_count").(string)
			if !ok {
				return errors.New("failed to get uuid_count")
			}
			r.SetId(uc)
			url, err := url.JoinPath("https://www.uuidtools.com/api/generate/v1/count/", uc)
			if err != nil {
				return fmt.Errorf("url.JoinPath() err = %w", err)
			}
			resp, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("http.Get() err = %w", err)
			}
			defer resp.Body.Close()
			return nil
		},
		Read: func(r *schema.ResourceData, _ any) error {
			return nil
		},
		Update: func(r *schema.ResourceData, _ any) error {
			return nil
		},
		Delete: func(r *schema.ResourceData, _ any) error {
			r.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"uuid_count": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
