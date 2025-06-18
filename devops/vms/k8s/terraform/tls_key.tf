
resource "tls_private_key" "private_key" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

resource "local_file" "private_key_pem" {
  filename = "${path.root}/files/ssh_keys/${var.ssh_public_key_filename}"
  content  = tls_private_key.private_key.private_key_pem
  file_permission = "0600"
}

resource "local_file" "public_key_openssh" {
  filename = "${path.root}/files/ssh_keys/${var.ssh_public_key_filename}.pub"
  content  = tls_private_key.private_key.public_key_openssh
}