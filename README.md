# terraform-provider-rds

[![CI](https://github.com/winebarrel/terraform-provider-rds/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/terraform-provider-rds/actions/workflows/ci.yml)

It is a provider to get the amount of memory for each instance class of RDS.

## Usage

```tf
terraform {
  required_providers {
    rds = {
      source = "winebarrel/rds"
    }
  }
}

provider "rds" {
}

data "rds_db_instance_memory" "main" {
  for_each       = { for i in data.rds_db_instances.main.instances : i.name => i.instance_class if i.tags["Env"] == "production" }
  instance_class = each.value
}

data "rds_db_instance_memory_map" "main" {
}

data "rds_db_instances" "main" {
  # filter {
  #   name   = "db-instance-id"
  #   values = ["database-1"]
  # }
}

output "rds" {
  description = "rds"
  value       = data.rds_db_instance_memory.main
}

output "memory_map" {
  description = "rds"
  value       = data.rds_db_instance_memory_map.main.memory_by_instance_class
}
```

```
$ terraform apply

...

Outputs:

rds = {
  "database-1" = {
    "id" = "terraform-20220716115109174100000002"
    "instance_class" = "db.t3.micro"
    "memory" = 1
    "tags" = {
      "Env" = "production"
    }
  }
  ...
}
memory_map = tomap({
  "db.m1.large" = 7.5
  "db.m1.medium" = 3.75
  "db.m1.small" = 1.7
  ...
}
```
