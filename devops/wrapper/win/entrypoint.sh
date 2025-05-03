#!/bin/bash

# SSH bilgileri
SSH_USER="Administrator"
SSH_HOST="192.168.122.118"
SSH_PASS="kcn-6953162"
REMOTE_DIR="C:\Users\Administrator\Desktop\AS\prj2"

# Log
echo "[ACT WRAPPER] Komut: $@" >> /tmp/act_ssh.log

# Komutu Windows makinede çalıştır
sshpass -p "$SSH_PASS" ssh -o StrictHostKeyChecking=no ${SSH_USER}@${SSH_HOST} \
"cd \"$REMOTE_DIR\" && powershell -Command \"$@\""
