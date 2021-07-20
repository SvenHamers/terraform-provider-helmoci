---
page_title: "helmoci_registry Data Source - terraform-provider-helmoci"
subcategory: ""
description: |-
  
---

# Data Source `helmoci_registry`


```
data "helmoci_registry" "example" {
    name = "just a name"
    chart_url = "registry.example.com/helm/test"
    version_tag = "" // optional defaults to latest
    registry_username = "" //optional
    registry_password = "" //optional
}
```


all charts are saved under the working dir /charts