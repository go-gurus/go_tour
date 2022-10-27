variable "deployment_target" {
  type        = string
  description = "Deployment Target for project."
  default     = "main"
}

variable "location" {
  type        = string
  description = "Azure region where the resource group will be created."
  default     = "westeurope"
}

variable "sku_instance_type" {
  type        = string
  description = "Azure SKU instance type"
  default     = "B1"
}

variable "sku_service_plan" {
  type        = string
  description = "Azure SKU service plan"
  default     = "Basic"
}

variable "docker_registry_server_username" {
  type        = string
  description = "The docker container registry username will be set by the gitlab pipeline"
}

variable "docker_registry_server_password" {
  type        = string
  description = "The docker container registry password will be set by the gitlab pipeline"
}

variable "docker_registry" {
  type        = string
  description = "The docker container registry address"
}

variable "main_service_name" {
  type        = string
  description = "name for main service"
  default     = "golang-for-developers"
}

variable "main_service_container_name" {
  type        = string
  description = "name for main service"
  default     = "golang_for_developers"
}

variable "main_service_container_tag" {
  type        = string
  description = "container tag for main service"
}