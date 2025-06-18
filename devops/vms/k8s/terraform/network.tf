
resource "libvirt_network" "example_network" {
  name      = "terraform-network"
  bridge    = "br-terraform"
  mode      = "nat"
  domain    = "terraform.local"
  addresses = ["192.168.200.0/24"]

  dhcp {
    enabled = true
  }

  autostart = true
}