resource "aws_iam_role" "lambda_exec" {
  name = "sns-sqs-cbr-lambda"
  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Action": [
          "sts:AssumeRole"
        ],
        "Principal": {
          "Service": "lambda.amazonaws.com"
        },
        "Effect": "Allow",
        "Sid": ""
      }
    ]
}
EOF
}

resource "aws_lambda_function" "cbr" {
  function_name = "sns-sqs-upload-cbr"
  s3_bucket = "${var.s3_bucket_upload_name}-${var.s3_version}"
  s3_key = var.s3-key
  handler = var.lambda-handler
  layers = ["arn:aws:lambda:eu-central-1:292169987271:layer:AWSLambda-Python36-SciPy1x:20"]
  source_code_hash = base64sha256(var.s3-source)
  runtime = "python3.6"
  role = aws_iam_role.lambda_exec.arn
  memory_size = 1024
  timeout = 5
}

resource "aws_lambda_permission" "sns" {
  statement_id = "AllowExecutionFromSNS"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cbr.function_name
  principal = "sns.amazonaws.com"
  source_arn = aws_sns_topic.upload.arn
}
