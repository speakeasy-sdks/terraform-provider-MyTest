resource "MyTest_zone" "my_zone" {
  account_id = 8
  code       = "...my_code..."
  config = {
    api_url       = "...my_api_url..."
    appliance_url = "...my_appliance_url..."
    datacenter    = "...my_datacenter..."
  }
  credential = {
    id = 5
  }
  description    = "...my_description..."
  enabled        = true
  group_id       = 3
  name           = "Ada Kling"
  scale_priority = 9
  visibility     = "...my_visibility..."
  zone_type = {
    code = "vmware"
  }
}