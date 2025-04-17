terraform {
  required_providers {
    libvirt = {
      source  = "dmacvicar/libvirt"
      version = "0.6.12" # En son sürümü kontrol edin
    }
  }

  required_version = ">= 1.0"
}

provider "libvirt" {
  uri = "qemu:///system"
}
