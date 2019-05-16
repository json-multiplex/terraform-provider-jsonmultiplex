provider "jsonmultiplex" {
  account_id = "abe7b9b9-4da5-4eed-9d5a-fb28b7ca5a61"
  user_id    = "root"
  password   = "letmein"
  iam_uri    = "localhost:3000"
}

resource "jsonmultiplex_user" "foo" {
  name     = "foo"
  password = "bar"
}
