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
    
}