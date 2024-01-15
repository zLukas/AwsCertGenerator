data "aws_iam_policy_document" "cognito_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Federated"
      identifiers = ["cognito-identity.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}


data "aws_iam_policy_document" "policy_doc" {
  statement {
    effect    = "Allow"
    actions   = var.policy.actions
    resources = var.policy.resources
  }
}
resource "aws_iam_policy" "cognito_policy" {
  name        = var.policy.name
  path        = "/cognito/${var.pool_name}/"
  description = "Cognito user policy"

  policy = data.aws_iam_policy_document.policy_doc.json
}

resource "aws_iam_role" "cognito_role"{

    name = "Cognito${var.pool_name}UserRoles"
    assume_role_policy = data.aws_iam_policy_document.cognito_assume_role.json
}

resource "aws_iam_policy_attachment" "cognito_attachment" {
  name       = "cognitoAttachment"
  roles      = [aws_iam_role.cognito_role.name] 
  policy_arn = aws_iam_policy.cognito_policy.arn
}
