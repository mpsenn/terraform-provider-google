// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package dns

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceDNSResponsePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSResponsePolicyCreate,
		Read:   resourceDNSResponsePolicyRead,
		Update: resourceDNSResponsePolicyUpdate,
		Delete: resourceDNSResponsePolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDNSResponsePolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"response_policy_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The user assigned name for this Response Policy, such as 'myresponsepolicy'.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The description of the response policy, such as 'My new response policy'.`,
				Default:     "Managed by Terraform",
			},
			"gke_clusters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The list of Google Kubernetes Engine clusters that can see this zone.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gke_cluster_name": {
							Type:     schema.TypeString,
							Required: true,
							Description: `The resource name of the cluster to bind this ManagedZone to.
This should be specified in the format like
'projects/*/locations/*/clusters/*'`,
						},
					},
				},
			},
			"networks": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `The list of network names specifying networks to which this policy is applied.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_url": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
							Description: `The fully qualified URL of the VPC network to bind to.
This should be formatted like
'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'`,
						},
					},
				},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceDNSResponsePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	responsePolicyNameProp, err := expandDNSResponsePolicyResponsePolicyName(d.Get("response_policy_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("response_policy_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(responsePolicyNameProp)) && (ok || !reflect.DeepEqual(v, responsePolicyNameProp)) {
		obj["responsePolicyName"] = responsePolicyNameProp
	}
	descriptionProp, err := expandDNSResponsePolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networksProp, err := expandDNSResponsePolicyNetworks(d.Get("networks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("networks"); ok || !reflect.DeepEqual(v, networksProp) {
		obj["networks"] = networksProp
	}
	gkeClustersProp, err := expandDNSResponsePolicyGkeClusters(d.Get("gke_clusters"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gke_clusters"); !tpgresource.IsEmptyValue(reflect.ValueOf(gkeClustersProp)) && (ok || !reflect.DeepEqual(v, gkeClustersProp)) {
		obj["gkeClusters"] = gkeClustersProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new ResponsePolicy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ResponsePolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
	})
	if err != nil {
		return fmt.Errorf("Error creating ResponsePolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/responsePolicies/{{response_policy_name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating ResponsePolicy %q: %#v", d.Id(), res)

	return resourceDNSResponsePolicyRead(d, meta)
}

func resourceDNSResponsePolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy_name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ResponsePolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DNSResponsePolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading ResponsePolicy: %s", err)
	}

	if err := d.Set("response_policy_name", flattenDNSResponsePolicyResponsePolicyName(res["responsePolicyName"], d, config)); err != nil {
		return fmt.Errorf("Error reading ResponsePolicy: %s", err)
	}
	if err := d.Set("description", flattenDNSResponsePolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading ResponsePolicy: %s", err)
	}
	if err := d.Set("networks", flattenDNSResponsePolicyNetworks(res["networks"], d, config)); err != nil {
		return fmt.Errorf("Error reading ResponsePolicy: %s", err)
	}
	if err := d.Set("gke_clusters", flattenDNSResponsePolicyGkeClusters(res["gkeClusters"], d, config)); err != nil {
		return fmt.Errorf("Error reading ResponsePolicy: %s", err)
	}

	return nil
}

func resourceDNSResponsePolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ResponsePolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandDNSResponsePolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networksProp, err := expandDNSResponsePolicyNetworks(d.Get("networks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("networks"); ok || !reflect.DeepEqual(v, networksProp) {
		obj["networks"] = networksProp
	}
	gkeClustersProp, err := expandDNSResponsePolicyGkeClusters(d.Get("gke_clusters"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("gke_clusters"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, gkeClustersProp)) {
		obj["gkeClusters"] = gkeClustersProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy_name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating ResponsePolicy %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "PATCH",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutUpdate),
	})

	if err != nil {
		return fmt.Errorf("Error updating ResponsePolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating ResponsePolicy %q: %#v", d.Id(), res)
	}

	return resourceDNSResponsePolicyRead(d, meta)
}

func resourceDNSResponsePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for ResponsePolicy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy_name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	// if gke clusters are attached, they need to be detached before the response policy can be deleted
	if d.Get("gke_clusters.#").(int) > 0 {
		patched := make(map[string]interface{})
		patched["gkeClusters"] = nil

		url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy_name}}")
		if err != nil {
			return err
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   project,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      patched,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating Policy %q: %s", d.Id(), err)
		}
	}

	// if networks are attached, they need to be detached before the response policy can be deleted
	if d.Get("networks.#").(int) > 0 {
		patched := make(map[string]interface{})
		patched["networks"] = nil

		url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/responsePolicies/{{response_policy_name}}")
		if err != nil {
			return err
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   project,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      patched,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating Policy %q: %s", d.Id(), err)
		}
	}
	log.Printf("[DEBUG] Deleting ResponsePolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "ResponsePolicy")
	}

	log.Printf("[DEBUG] Finished deleting ResponsePolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceDNSResponsePolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"projects/(?P<project>[^/]+)/responsePolicies/(?P<response_policy_name>[^/]+)",
		"(?P<project>[^/]+)/(?P<response_policy_name>[^/]+)",
		"(?P<response_policy_name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/responsePolicies/{{response_policy_name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDNSResponsePolicyResponsePolicyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyNetworks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"network_url": flattenDNSResponsePolicyNetworksNetworkUrl(original["networkUrl"], d, config),
		})
	}
	return transformed
}
func flattenDNSResponsePolicyNetworksNetworkUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSResponsePolicyGkeClusters(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"gke_cluster_name": flattenDNSResponsePolicyGkeClustersGkeClusterName(original["gkeClusterName"], d, config),
		})
	}
	return transformed
}
func flattenDNSResponsePolicyGkeClustersGkeClusterName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDNSResponsePolicyResponsePolicyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSResponsePolicyNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetworkUrl, err := expandDNSResponsePolicyNetworksNetworkUrl(original["network_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["networkUrl"] = transformedNetworkUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSResponsePolicyNetworksNetworkUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		return v, nil
	}
	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
	if err != nil {
		return "", err
	}
	return tpgresource.ConvertSelfLinkToV1(url), nil
}

func expandDNSResponsePolicyGkeClusters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedGkeClusterName, err := expandDNSResponsePolicyGkeClustersGkeClusterName(original["gke_cluster_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedGkeClusterName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["gkeClusterName"] = transformedGkeClusterName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSResponsePolicyGkeClustersGkeClusterName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
