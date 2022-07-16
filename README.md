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

data "rds_db_instances" "rds" {
}

output "rds" {
  description = "rds"
  value       = data.rds_db_instance_memory.rds
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
```
