/**
 * Copyright 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

variable "project_id" {
  description = "The project ID to create the database in."
  type = string
}

variable "database_id" {
  description = "The database ID to create the database with."
  type = string
}

variable "location" {
  description = "The location ID to create the database in."
  type = string
}

variable "database_type" {
  description = "The database type to create the database with."
  type = string
  default = "FIRESTORE_NATIVE"

  validation {
    condition = var.database_type == "FIRESTORE_NATIVE" || var.database_type == "DATASTORE_MODE"
    error_message = "Invalid database type. Database type can be either FIRESTORE_NATIVE (or) DATASTORE_MODE."
  }
}

variable "concurrency_mode" {
  description = "The concurrency mode to create the database with."
  type = string
  default = "OPTIMISTIC"

  validation {
    condition = var.concurrency_mode == "OPTIMISTIC" || var.concurrency_mode == "PESSIMISTIC" || var.concurrency_mode == "OPTIMISTIC_WITH_ENTITY_GROUPS"
    error_message = "Invalid concurrency mode. Concurrency mode can be either OPTIMISTIC (or) PESSIMISTIC (or) OPTIMISTIC_WITH_ENTITY_GROUPS."
  }
}

variable "delete_protection_state" {
  description = "The deletion protection state to create the database with."
  type = string
  default = "DELETE_PROTECTION_ENABLED"

  validation {
    condition = var.delete_protection_state == "DELETE_PROTECTION_ENABLED" || var.delete_protection_state == "DELETE_PROTECTION_DISABLED"
    error_message = "Invalid deletion protection state. Deletion protection state can be either DELETE_PROTECTION_ENABLED (or) DELETE_PROTECTION_DISABLED."
  }
}

variable "kms_key_name" {
  description = "The resource ID of the Customer-managed Encryption Key (CMEK) using which the created database will be encrypted."
  type = string
  default = null
}

variable "point_in_time_recovery_enablement" {
  description = "Determines whether point-in-time recovery is enabled for the database."
  type = string
  default = "POINT_IN_TIME_RECOVERY_ENABLED"

  validation {
    condition = var.point_in_time_recovery_enablement == "POINT_IN_TIME_RECOVERY_ENABLED" || var.point_in_time_recovery_enablement == "POINT_IN_TIME_RECOVERY_DISABLED"
    error_message = "Invalid point in time recovery configuration. Valid values are POINT_IN_TIME_RECOVERY_ENABLED (or) POINT_IN_TIME_RECOVERY_DISABLED."
  }
}

variable "deletion_policy" {
  description = "The deletion policy to create the database with."
  type = string
  default = "DELETED"
}

variable "backup_schedule_configuration" {
  description = "Describes the backup configuration for the database."
  type = object({
    weekly_recurrence = optional(object({
      day = string
      retention = string
    }))

    daily_recurrence = optional(object({
      retention = string
    }))
  })
  default = null
}

variable "composite_index_configuration" {
  description = "List of composite indexes to create for the database."
  type = list(object({
    index_id = string
    collection = string
    query_scope = optional(string, "COLLECTION")
    api_scope = optional(string, "ANY_API")
    fields = list(object({
      field_path = string
      order = optional(string)
      array_config = optional(string)
      vector_config = optional(object({
        dimension = number
      }))
    }))
  }))
  default = []

  validation {
    condition = alltrue(flatten([
      for item in var.composite_index_configuration : alltrue(flatten([
        for field in item.fields : length([for v in [field.order, field.array_config, field.vector_config] : v if v != null]) == 1
      ]))
    ]))
    error_message = "For each 'field' object must have exactly one of 'order', 'array_config', or 'vector_config' set."
  }

  validation {
    condition = alltrue(flatten([
      for item in var.composite_index_configuration : (item.query_scope == "COLLECTION" || item.query_scope == "COLLECTION_GROUP" || item.query_scope == "COLLECTION_RECURSIVE")
    ]))
    error_message = "Invalid query scope. Query scope can be either COLLECTION (or) COLLECTION_GROUP (or) COLLECTION_RECURSIVE."
  }

    validation {
    condition = alltrue(flatten([
      for item in var.composite_index_configuration : (item.api_scope == "ANY_API" || item.api_scope == "DATASTORE_MODE_API")
    ]))
    error_message = "Invalid API scope. API scope can be either ANY_API (or) DATASTORE_MODE_API."
  }
}

variable "field_configuration" {
  description = "Describes the single fields that are exempt from default indexing."
  type = list(object({
    collection = string
    field = string
    ttl_enabled = optional(bool, false)
    ascending_index_query_scope = optional(set(string), [])
    descending_index_query_scope = optional(set(string), [])
    array_index_query_scope = optional(set(string), [])
  }))
  default = []
  
  validation {
    condition = alltrue(flatten([
      for field in var.field_configuration : alltrue(flatten([
        for scope in field.ascending_index_query_scope : (scope == "COLLECTION" || scope == "COLLECTION_GROUP")
      ]))
    ]))
    error_message = "Invalid query scope provided for ascending index. Query scope can either be COLLECTION (or) COLLECTION_GROUP"
  }

  validation {
    condition = alltrue(flatten([
      for field in var.field_configuration : alltrue(flatten([
        for scope in field.descending_index_query_scope : (scope == "COLLECTION" || scope == "COLLECTION_GROUP")
      ]))
    ]))
    error_message = "Invalid query scope provided for descending index. Query scope can either be COLLECTION (or) COLLECTION_GROUP"
  }

  validation {
    condition = alltrue(flatten([
      for field in var.field_configuration : alltrue(flatten([
        for scope in field.array_index_query_scope : (scope == "COLLECTION" || scope == "COLLECTION_GROUP")
      ]))
    ]))
    error_message = "Invalid query scope provided for array index. Query scope can either be COLLECTION (or) COLLECTION_GROUP"
  }
}

