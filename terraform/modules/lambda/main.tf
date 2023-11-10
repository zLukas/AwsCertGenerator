resource "aws_lambda_function" "lambda" {
   
  # If the file is not in the current working directory you will need to include a
  # path.module in the filename.
  filename      = var.zip_file
  function_name = var.lambda_name
  role          = aws_iam_role.iam_for_lambda.arn
  handler = var.handler
  runtime = var.runtime
  timeout = var.timeout
  memory_size = var.memory_size


  environment {
    variables = var.env_vars
      }
}