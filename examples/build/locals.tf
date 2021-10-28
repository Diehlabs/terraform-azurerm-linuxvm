locals {
  tags = {
    test_for   = var.gh_repo
    gh_run_id  = var.gh_run_id
    created_by = "terratest"
    location   = "westus"
  }
  # vm_name = replace("${local.tags.created_by}-${var.gh_run_id}-${var.gh_repo}", "/", "-")
  vm_name = "${local.tags.created_by}-${var.gh_run_id}"
}
