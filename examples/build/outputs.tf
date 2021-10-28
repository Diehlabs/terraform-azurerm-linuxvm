output "ssh_key" {
  value     = tls_private_key.terratest.private_key_pem
  sensitive = true
}
