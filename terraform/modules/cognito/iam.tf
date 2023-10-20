data "aws_iam_policy_document" "policy_doc" {
  statement {
    effect    = "Allow"
    actions   = var.policy.actions
    resources = var.policy.resources
  }
}
resource "aws_iam_policy" "cognito_policy" {
  name        = var.policy.name
  path        = "/cognito/${var.pool}"
  description = "Cognito user policy"

  policy = var.aws_iam_policy_document.policy_doc.json
}

resource "aws_iam_role" "cognito_role"{

    name = "Cognito${var.pool}UserRoles"
    assume_role_policy = jsonencode({
        Version = "2012-10-17"
        Statement = [{
            Action = "sts:AssumeRole"
            Effect = "Allow"
            Sid    = ""
            Principal = {
            Service = "cognito.amazonaws.com"
            }
        },]
    })
    inline_policy {
        name = "IAM policy"
        policy = aws_iam_policy.cognito_policy.json
    }
}


