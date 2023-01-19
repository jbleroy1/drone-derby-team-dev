resource "google_app_engine_application" "app" {
  project     = google_project.my_project.project_id
  location_id = "location"
  database_type = "CLOUD_FIRESTORE"
}