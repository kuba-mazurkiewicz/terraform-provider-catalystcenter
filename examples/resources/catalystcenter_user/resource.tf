resource "catalystcenter_user" "example" {
  first_name = "john"
  last_name  = "doe"
  username   = "johndoe"
  password   = "C1sco1357"
  email      = "john.doe@example.com"
  role_ids   = ["5f8d9a3a-3c6a-4a1c-9a1b-1a4f5a4a4a4a"]
}
