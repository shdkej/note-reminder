resource "aws_cloudwatch_event_rule" "every_day" {
    name = "every-day-cron"
    description = "Execute Every day"
    schedule_expression = "cron(0 20 * * ? *)"
}

resource "aws_cloudwatch_event_target" "check_lambda" {
    rule = aws_cloudwatch_event_rule.every_day.name
    target_id = "lambda"
    arn = aws_lambda_function.cbr.arn
}

resource "aws_lambda_permission" "allow_cloudwatch" {
    statement_id = "AllowExecutionFromCloudWatch"
    action = "lambda:InvokeFunction"
    function_name = aws_lambda_function.cbr.function_name
    principal = "events.amazonaws.com"
    source_arn = aws_cloudwatch_event_rule.every_day.arn
}
