# Simple Example With Backup

This example illustrates how to use the `firestore` module with backups enabled.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| project\_id | The ID of the project in which to provision resources. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| daily\_backup\_schedule\_id | Unique identifier for the daily backup schedule. |
| database\_id | Unique identifier of the created firestore database. |
| weekly\_backup\_schedule\_id | Unique identifier for the weekly backup schedule. |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

To provision this example, run the following from within this directory:
- `terraform init` to get the plugins
- `terraform plan` to see the infrastructure plan
- `terraform apply` to apply the infrastructure build
- `terraform destroy` to destroy the built infrastructure
