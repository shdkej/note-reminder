resource "aws_sns_topic" "upload" {
  name = "sns-sqs-upload-topic"
}

resource "aws_sns_topic_policy" "upload" {
  arn = aws_sns_topic.upload.arn
  policy = data.aws_iam_policy_document.sns_upload.json
}

resource "aws_sns_topic_subscription" "sqs" {
  topic_arn = aws_sns_topic.upload.arn
  protocol = "sqs"
  endpoint = aws_sqs_queue.upload.arn
}

resource "aws_sns_topic_subscription" "lambda" {
  topic_arn = aws_sns_topic.upload.arn
  protocol = "lambda"
  endpoint = aws_lambda_function.cbr.arn
}

data "aws_iam_policy_document" "sns_upload" {
  policy_id = "snssqssns"
  statement {
    actions = [
      "SNS:Publish",
    ]
    condition {
      test = "ArnLike"
      variable = "aws:SourceArn"

      values = [
        "arn:aws:s3:::${var.s3_bucket_upload_name}-${var.s3_version}",
      ]
    }
    effect = "Allow"
    principals {
      type = "AWS"
      identifiers = [
        "*"
      ]
    }
    resources = [
      aws_sns_topic.upload.arn,
    ]
    sid = "snssqssnss3upload"
  }
}
