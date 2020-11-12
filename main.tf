provider "aws" {
  region = "eu-central-1"
}

module "note-reminder" {
  source = "./event_handling"
  s3-source = "./event_handling/recommend.zip"
  s3_bucket_upload_name = "my-note"
  s3-key = "recommend.zip"
}

terraform {
  backend "s3" {
    bucket = "shdkej-personal-state.shdkej.com"
    key = "state-note-reminder"
    region = "eu-central-1"
  }
}
