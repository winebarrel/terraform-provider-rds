# terraform-provider-rds

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

data "rds_db_instance_memory" "rds" {
  for_each       = { for i in data.rds_db_instances.test.instances : i.name => i.instance_class }
  instance_class = each.value
}

date "rds_db_instance_memory_map" "rds" {
}

data "rds_db_instances" "rds" {
  # filter {
  #   name   = "db-instance-id"
  #   values = ["database-1"]
  # }
}

output "rds" {
  description = "rds"
  value       = data.rds_db_instance_memory.rds
}

output "memory_map" {
  description = "rds"
  value       = data.rds_db_instance_memory_map.rds.memory_by_instance_class
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
