variable "region" {
  default = "eu-central-1"
}

variable "s3_bucket_upload_name" {
  default = "s3-test"
}

variable "s3_version" {
  default = "0.0.1"
}

variable "s3-key" {
  default = "s3-py"
}

variable "s3-source" {
  default = "test_handler.zip"
}

variable "lambda-handler" {
  default = "recommender_system.getRecommend"
}
