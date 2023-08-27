resource "MyTest_zone" "my_zone" {
  account_id = 6
  code       = "...my_code..."
  config = {
    zone_vcenter_config = {
      api_url       = "...my_api_url..."
      appliance_url = "...my_appliance_url..."
      datacenter    = "...my_datacenter..."
    }
  }
  credential = {
    id = 6
  }
  description    = "...my_description..."
  enabled        = true
  group_id       = 3
  name           = "Ismael Little"
  scale_priority = 6
  visibility     = "...my_visibility..."
  zone_type = {
    code = "vmware"
  }
}