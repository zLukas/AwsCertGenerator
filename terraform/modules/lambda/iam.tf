data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name               = "LambdaIam"
  assume_role_policy = data.aws_iam_policy_document.lambda_assume_role.json
}

data "aws_iam_policy_document" "lambda_policy_doc" {
  statement {
    effect = "Allow"
    actions = var.lambda_iam_actions
    resources = var.lambda_iam_resources
  }
}

resource "aws_iam_policy" "lambda_permissions" {
  name        = "lambda_permissions"
  path        = "/"
  description = "IAM policy for Lambda"
  policy      = data.aws_iam_policy_document.lambda_policy_doc.json
}

resource "aws_iam_policy_attachment" "lambda_attachment" {
  name       = "lambdaAttachment"
  roles      = [aws_iam_role.iam_for_lambda.name] 
  policy_arn = aws_iam_policy.lambda_permissions.arn
}
