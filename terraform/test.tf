resource "google_storage_bucket" "test001" {
    name            = "${local.project_id}-bucket-test001"
    project         = local.project_id
    location        = local.region
    force_destroy   = true
    uniform_bucket_level_access = true
}