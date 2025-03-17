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

output "database_id" {
  description = "Unique identifier of the created firestore database."
  value       = module.firestore.database_id
}

output "weekly_backup_schedule_id" {
  description = "Unique identifier for the weekly backup schedule."
  value       = module.firestore.weekly_backup_schedule_id
}

output "daily_backup_schedule_id" {
  description = "Unique identifier for the daily backup schedule."
  value       = module.firestore._backup_schedule_id
}