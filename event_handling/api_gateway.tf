resource "aws_api_gateway_rest_api" "api" {
  name = "lambda-api-gateway"
  description = "Proxy to handle requests to our API"
}

resource "aws_api_gateway_resource" "resource" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  parent_id = aws_api_gateway_rest_api.api.root_resource_id
  path_part = "{proxy+}"
}

resource "aws_api_gateway_method" "method" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  resource_id = aws_api_gateway_resource.resource.id
  http_method = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  resource_id = aws_api_gateway_method.method.resource_id
  http_method = aws_api_gateway_method.method.http_method

  integration_http_method = "POST"
  type = "AWS_PROXY"
  uri = aws_lambda_function.cbr.invoke_arn
}

resource "aws_api_gateway_method" "method_root" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  resource_id = aws_api_gateway_rest_api.api.root_resource_id
  http_method = "ANY"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "integration_root" {
  rest_api_id = aws_api_gateway_rest_api.api.id
  resource_id = aws_api_gateway_method.method_root.resource_id
  http_method = aws_api_gateway_method.method_root.http_method

  integration_http_method = "POST"
  type = "AWS_PROXY"
  uri = aws_lambda_function.cbr.invoke_arn
}

resource "aws_api_gateway_deployment" "lambda" {
  depends_on = [
    aws_api_gateway_integration.integration,
    aws_api_gateway_integration.integration_root,
  ]
  rest_api_id = aws_api_gateway_rest_api.api.id
  stage_name = "test"
}

resource "aws_lambda_permission" "apigw" {
  statement_id = "AllowAPIGatewayInvoke"
  action = "lambda:InvokeFunction"
  function_name = aws_lambda_function.cbr.arn
  principal = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_deployment.lambda.execution_arn}/*/*"
}

output "lambda_api" {
  value = aws_api_gateway_deployment.lambda.invoke_url
}
