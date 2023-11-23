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
  name               = "${var.lambda_name}LambdaIAM"
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
  name        = "lambda${var.lambda_name}permissions"
  path        = "/"
  description = "IAM policy for Lambda"
  policy      = data.aws_iam_policy_document.lambda_policy_doc.json
}

resource "aws_iam_policy_attachment" "lambda_attachment" {
  name       = "${var.lambda_name}lambdaAttachment"
  roles      = [aws_iam_role.iam_for_lambda.name] 
  policy_arn = aws_iam_policy.lambda_permissions.arn
}
