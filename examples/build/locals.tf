locals {
  tags = {
    test_for   = var.gh_repo
    gh_run_id  = var.gh_run_id
    created_by = "Terratest"
    location   = "westus"
  }
}
