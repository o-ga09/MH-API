variable "image_tag" {
  description = "The image tag for containers"
  type        = string
}

variable "service_name" {
  description = "The service name for the API container"
  type        = string
  default     = "mh-api"
}

variable "project_id" {
  description = "The GCP project ID"
  type        = string
}
