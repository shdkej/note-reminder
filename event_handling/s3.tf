resource "aws_s3_bucket" "upload" {
  bucket = "${var.s3_bucket_upload_name}-${var.s3_version}"
  acl = "private"

  tags = {
    Name = "Website"
  }

  versioning {
    enabled = true
  }

  policy = <<EOF
{
  "Id": "bucket_policy_site",
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "bucket_policy_site_main",
      "Action": [
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::${var.s3_bucket_upload_name}-${var.s3_version}/*",
      "Principal": "*"
    }
  ]
}
EOF

  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}

resource "aws_s3_bucket_object" "python" {
  bucket = aws_s3_bucket.upload.bucket
  key = var.s3-key
  source = var.s3-source
  etag = filemd5(var.s3-source)
}

resource "aws_s3_bucket_notification" "upload" {
  bucket = aws_s3_bucket.upload.id

  topic {
    topic_arn = aws_sns_topic.upload.arn
    events = ["s3:ObjectCreated:*"]
  }
}
