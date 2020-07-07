resource "aws_sqs_queue" "upload" {
  name = "sns-sqs-upload"
}

resource "aws_sqs_queue_policy" "test" {
  queue_url = aws_sqs_queue.upload.id
  policy = data.aws_iam_policy_document.sqs_upload.json
}

data "aws_iam_policy_document" "sqs_upload" {
  policy_id = "snssqssqs"
  statement {
    actions = [
      "sqs:SendMessage",
    ]
    condition {
      test = "ArnEquals"
      variable = "aws:SourceArn"

      values = [
        aws_sns_topic.upload.arn,
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
      aws_sqs_queue.upload.arn,
    ]
    sid = "snssqssqssns"
  }
}
