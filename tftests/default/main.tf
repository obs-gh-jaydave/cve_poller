resource "random_pet" "test" {}

data "observe_workspace" "default" {
  name = "Default"
}

data "observe_datastream" "default" {
  workspace = data.observe_workspace.default.oid
  name      = var.data_stream_name
}

module "default" {
  source      = "../.."

}
