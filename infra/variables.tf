variable "project_id" {
  description = "The project ID to deploy to"
  type        = string
  default     = "gilbert-learning-gcp-113"
}

variable "region" {
  type    = string
  default = "us-central1"
}

variable "service_image" {
  type    = string
  default = "us-central1-docker.pkg.dev/gilbert-learning-gcp-113/gilbert-test-repo/edt"
}

variable "service_image_version" {
  type    = string
  default = "local"
}

variable "cpu" {
  type    = string
  default = "2"
}

variable "mem" {
  type    = string
  default = "4096Mi"
}
