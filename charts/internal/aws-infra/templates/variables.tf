variable "ACCESS_KEY_ID" {
  description = "AWS Access Key ID of technical user"
  type        = "string"
}

variable "SECRET_ACCESS_KEY" {
  description = "AWS Secret Access Key of technical user"
  type        = "string"
}

variable "SESSION_TOKEN" {
  description = "The token that users must pass to the service API to use the temporary credentials."
  type        = "string"
}

