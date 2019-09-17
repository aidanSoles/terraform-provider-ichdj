resource "ichdj_random_joke" "first" {
  name = "first"
}

output "joke" {
  value = "${ichdj_random_joke.first.joke}"
}
