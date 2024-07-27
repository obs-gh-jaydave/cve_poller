variable "name" {
  type        = string
  description = "Poller name. Should be unique per datastream."
  default     = "CVE Poller"
}


variable "workspace" {
  type        = object({ oid = string })
  description = "Workspace to apply module to."
}

variable "datastream" {
  type = object({
    oid = string
  })
  description = <<-EOF
    Datastream to derive resources from.
  EOF
}

variable "interval_duration" {
  type        = string
  default     = "5m0s"
  description = <<-EOF
      How often the poller should poll S3. The default interval should be appropriate for what you want.
    EOF
}


# tflint-ignore: terraform_unused_declarations
variable "description" {
  type        = string
  description = "CVE Poller pulling MITRE data"
  default     = "CVE Poller"
}

variable "bucket_url" {
  type        = string
  description = "s3 bucket url for GET"
  default     = "https://s3.us-west-2.amazonaws.com/observeinc.com/icons/apps/terraform-observe-field-enablement-1/sample_data.json"
}