# Create the Cloud Run service
resource "google_cloud_run_service" "run_service" {
  name     = "fulcrum"
  location = var.region

  template {
    spec {
      containers {
        image = "${var.service_image}:${var.service_image_version}"
        resources {
          limits = {
            cpu    = var.cpu
            memory = var.mem
          }
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/minScale"      = 1
        "autoscaling.knative.dev/maxScale"      = 10
        #"run.googleapis.com/cloudsql-instances" = "${var.project_id}:${var.region}:${var.database_instance_name}"
        "run.googleapis.com/client-name"        = "terraform"
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  # Waits for the Cloud Run API to be enabled from main.tf
  depends_on = [google_project_service.run_api]
}

# Allow unauthenticated users to invoke the service
resource "google_cloud_run_service_iam_member" "run_all_users" {
  service  = google_cloud_run_service.run_service.name
  location = google_cloud_run_service.run_service.location
  role     = "roles/run.invoker"
  member   = "allUsers"
}
