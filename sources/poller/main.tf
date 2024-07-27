resource "observe_poller" "cve_poller" {
  workspace  = data.observe_workspace.default.id
  datastream = var.datastream.oid
  name       = var.name
  interval   = var.interval_duration

  http {
    template {
      headers = {
        accept = "application/json"
      }
    }
    request {
      url    = var.bucket_url
      method = "GET"
    }
  }
}
