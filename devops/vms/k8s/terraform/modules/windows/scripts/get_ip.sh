#!/bin/bash
# get_ip.sh

VM_NAME="$1"


IP=$(sudo virsh domifaddr $VM_NAME | grep ipv4 | awk '{print $4}' | cut -d'/' -f1)
echo "{\"ip\": \"$IP\"}"
