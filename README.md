# terraform-google-firestore

## Description
### Tagline
This is an auto-generated module.

### Detailed
This module was generated from [terraform-google-module-template](https://github.com/terraform-google-modules/terraform-google-module-template/), which by default generates a module that simply creates a GCS bucket. As the module develops, this README should be updated.

The resources/services/activations/deletions that this module will create/trigger are:

- Create a GCS bucket with the provided name

### PreDeploy
To deploy this blueprint you must have an active billing account and billing permissions.

## Architecture
![alt text for diagram](https://www.link-to-architecture-diagram.com)
1. Architecture description step no. 1
2. Architecture description step no. 2
3. Architecture description step no. N

## Documentation
- [Hosting a Static Website](https://cloud.google.com/storage/docs/hosting-static-website)

## Deployment Duration
Configuration: X mins
Deployment: Y mins

## Cost
[Blueprint cost details](https://cloud.google.com/products/calculator?id=02fb0c45-cc29-4567-8cc6-f72ac9024add)

## Usage

Basic usage of this module is as follows:

```hcl
module "firestore" {
  source  = "terraform-google-modules/firestore/google"
  version = "~> 0.1"

  project_id  = "<PROJECT ID>"
  bucket_name = "gcs-test-bucket"
}
```

Functional examples are included in the
[examples](./examples/) directory.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| backup\_schedule\_configuration | Backup schedule configuration for the Firestore Database. | <pre>object({<br>    weekly_recurrence = optional(object({<br>      day = string<br>      retention = string<br>    }))<br><br>    daily_recurrence = optional(object({<br>      retention = string<br>    }))<br>  })</pre> | `null` | no |
| composite\_index\_configuration | Composite index configuration for the Firestore Database. | <pre>list(object({<br>    index_id = string<br>    collection = string<br>    query_scope = optional(string, "COLLECTION")<br>    api_scope = optional(string, "ANY_API")<br>    fields = list(object({<br>      field_path = string<br>      order = optional(string)<br>      array_config = optional(string)<br>      vector_config = optional(object({<br>        dimension = number<br>      }))<br>    }))<br>  }))</pre> | `[]` | no |
| concurrency\_mode | Concurrency control mode to be used for the Firestore Database. | `string` | `"OPTIMISTIC"` | no |
| database\_id | Unique identifier of the Firestore Database. | `string` | n/a | yes |
| database\_type | Database type used to created the Firestore Database. | `string` | `"FIRESTORE_NATIVE"` | no |
| delete\_protection\_state | Determines whether deletion protection is enabled or not for the Firestore Database. | `string` | `"DELETE_PROTECTION_ENABLED"` | no |
| deletion\_policy | Deletion policy enforced when Firestore Database is destroyed via Terraform. | `string` | `"DELETED"` | no |
| field\_configuration | Single field configurations for the Firestore Database. | <pre>list(object({<br>    collection = string<br>    field = string<br>    ttl_enabled = optional(bool, false)<br>    ascending_index_query_scope = optional(set(string), [])<br>    descending_index_query_scope = optional(set(string), [])<br>    array_index_query_scope = optional(set(string), [])<br>  }))</pre> | `[]` | no |
| kms\_key\_name | The resource ID of the Customer-managed Encryption Key (CMEK) using which the created database will be encrypted. | `string` | `null` | no |
| location | The location in which the Firesotre Database is created. | `string` | n/a | yes |
| point\_in\_time\_recovery\_enablement | Determines whether point-in-time recovery is enabled for the Firestore Database. | `string` | `"POINT_IN_TIME_RECOVERY_ENABLED"` | no |
| project\_id | The ID of the project in which the Firestore resources are created. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| composite\_index\_ids | List of composite indices for the firestore database. |
| daily\_backup\_schedule\_id | The unique backup schedule identifier across all locations and databases for the given project. |
| database\_id | The database id of the firestore database. |
| field\_ids | List of firestore fields created for the database. |
| weekly\_backup\_schedule\_id | The unique backup schedule identifier across all locations and databases for the given project. |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Requirements

These sections describe requirements for using this module.

### Software

The following dependencies must be available:

- [Terraform][terraform] v0.13
- [Terraform Provider for GCP][terraform-provider-gcp] plugin v3.0

### Service Account

A service account with the following roles must be used to provision
the resources of this module:

- Storage Admin: `roles/storage.admin`

The [Project Factory module][project-factory-module] and the
[IAM module][iam-module] may be used in combination to provision a
service account with the necessary roles applied.

### APIs

A project with the following APIs enabled must be used to host the
resources of this module:

- Google Cloud Storage JSON API: `storage-api.googleapis.com`

The [Project Factory module][project-factory-module] can be used to
provision a project with the necessary APIs enabled.

## Contributing

Refer to the [contribution guidelines](./CONTRIBUTING.md) for
information on contributing to this module.

[iam-module]: https://registry.terraform.io/modules/terraform-google-modules/iam/google
[project-factory-module]: https://registry.terraform.io/modules/terraform-google-modules/project-factory/google
[terraform-provider-gcp]: https://www.terraform.io/docs/providers/google/index.html
[terraform]: https://www.terraform.io/downloads.html

## Security Disclosures

Please see our [security disclosure process](./SECURITY.md).
