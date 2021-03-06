package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceApplicationProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSApplicationProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "name", "test-System-Secure-HTTP-abc"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_port", "false"),
					resource.TestCheckResourceAttr(
						"avi_applicationprofile.testApplicationProfile", "preserve_client_ip", "false"),
				),
			},
		},
	})

}

const testAccAVIDSApplicationProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_stringgroup" "system_compressiblestringgroup" {
    name= "System-Compressible-Content-Types"
}
data "avi_stringgroup" "system_cacheablestringgroup" {
    name= "System-Cacheable-Resource-Types"
}
resource "avi_applicationprofile" "testApplicationProfile" {
	name = "test-System-Secure-HTTP-abc"
	type = "APPLICATION_PROFILE_TYPE_HTTP"
	tenant_ref = data.avi_tenant.default_tenant.id
	http_profile {
		max_rps_uri = "0"
		keepalive_header = false
		max_rps_cip_uri = "0"
		x_forwarded_proto_enabled = true
		connection_multiplexing_enabled = true
		websockets_enabled = true
		enable_request_body_buffering = false
		hsts_enabled = true
		compression_profile {
			compressible_content_ref = data.avi_stringgroup.system_compressiblestringgroup.id
			type = "AUTO_COMPRESSION"
			compression = false
			remove_accept_encoding_header = true
		}
		xff_enabled = true
		disable_keepalive_posts_msie6 = true
		keepalive_timeout = "30000"
		ssl_client_certificate_mode = "SSL_CLIENT_CERTIFICATE_NONE"
		http_to_https = true
		spdy_enabled = false
		respond_with_100_continue = true
		client_body_timeout = "30000"
		httponly_enabled = true
		hsts_max_age = "365"
		client_max_header_size = "12"
		server_side_redirect_to_https = true
		max_bad_rps_cip = "0"
		client_max_request_size = "48"
		cache_config {
			min_object_size = "100"
			query_cacheable = false
			xcache_header = true
			age_header = true
			enabled = false
			default_expire = "600"
			max_cache_size = "0"
			heuristic_expire = false
			date_header = true
			aggressive = false
			max_object_size = "4194304"
			mime_types_group_refs = [data.avi_stringgroup.system_cacheablestringgroup.id]
		}
		max_rps_unknown_uri = "0"
		spdy_fwd_proxy_mode = false
		allow_dots_in_header_name = false
		client_header_timeout = "10000"
		post_accept_timeout = "30000"
		secure_cookie_enabled = true
		max_response_headers_size = "48"
		xff_alternate_name = "X-Forwarded-For"
		max_rps_cip = "0"
		enable_fire_and_forget = false
		max_rps_unknown_cip = "0"
		max_bad_rps_cip_uri = "0"
		max_bad_rps_uri = "0"
		use_app_keepalive_timeout = false
	}
	preserve_client_port = false
	preserve_client_ip = false
}

data "avi_applicationprofile" "testApplicationProfile" {
    name= "${avi_applicationprofile.testApplicationProfile.name}"
}
`
