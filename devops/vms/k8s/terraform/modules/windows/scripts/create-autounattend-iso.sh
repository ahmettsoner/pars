#!/bin/bash
set -e

SRC_DIR="$1"
OUTPUT_ISO="$2"

if [[ ! -f "$SRC_DIR/Autounattend.xml" ]]; then
  echo "❌ Autounattend.xml not found in $SRC_DIR"
  exit 1
fi

# Eğer ISO dosyası varsa sil
if [[ -f "$OUTPUT_ISO" ]]; then
  echo "ℹ️  Existing ISO found at $OUTPUT_ISO, removing..."
  rm -f "$OUTPUT_ISO"
fi

genisoimage -o "$OUTPUT_ISO" -V "AUTOUNATTEND" -r -J "$SRC_DIR"
echo "✅ Created $OUTPUT_ISO"