resource "azurerm_resource_group" "main-service-rg" {
  name     = "${var.main_service_name}-rg"
  location = var.location
}

resource "azurerm_app_service_plan" "main-service-sp" {
  name                = "${var.main_service_name}-sp"
  location            = var.location
  resource_group_name = azurerm_resource_group.main-service-rg.name
  kind                = "Linux"
  reserved            = true
  sku {
    tier = var.sku_service_plan
    size = var.sku_instance_type
  }
}

resource "azurerm_app_service" "main-service" {
  name                = var.main_service_name
  location            = var.location
  resource_group_name = azurerm_resource_group.main-service-rg.name
  app_service_plan_id = azurerm_app_service_plan.main-service-sp.id
  https_only          = "true"
  site_config {
    always_on        = "true"
    linux_fx_version = "DOCKER|${var.docker_registry}/${var.main_service_container_name}:${var.main_service_container_tag}"
  }
  auth_settings {
    enabled = false
  }
  app_settings = {
    "WEBSITES_ENABLE_APP_SERVICE_STORAGE" = false
    "DOCKER_REGISTRY_SERVER_URL"          = "https://${var.docker_registry}"
    "DOCKER_REGISTRY_SERVER_USERNAME"     = var.docker_registry_server_username
    "DOCKER_REGISTRY_SERVER_PASSWORD"     = var.docker_registry_server_password
    "PORT"                                = 8080
  }
}