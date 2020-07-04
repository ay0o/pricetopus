variable "password" {
  description = "Sender's password"
  type        = string
}

variable "products" {
  description = "A map of products"
  default = {
    "p1" : {
      "url" : "url",
      "price" : "price"
    },
    "p2" : {
      "url" : "url",
      "price" : "price"
    }
  }
}

variable "port" {
  description = "Email server port"
  type        = number
  default     = 587
}

variable "price" {
  description = "Max product price"
  type        = number
}

variable "recipient" {
  description = "Recipient's email address"
  type        = string
}

variable "sender" {
  description = "Sender's email address"
  type        = string
}

variable "server" {
  description = "Email server"
  type        = string
  default     = "smtp.gmail.com"
}

variable "trigger" {
  description = "Trigger to run the task. Accepts cron(...) or rate(time unit)"
  type        = string
  default     = "rate(1 hour)"
}

variable "url" {
  description = "Product URL"
  type        = string
}
