#!/usr/bin/env bash
# Fetch webgpu.h (webgpu-native/webgpu-headers) and wgpu-native shared libraries.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
THIRD="$ROOT/third_party"
HEADERS_TAG="${HEADERS_TAG:-main}"
WGPU_NATIVE_TAG="${WGPU_NATIVE_TAG:-v29.0.1.1}"

mkdir -p "$THIRD"

echo "==> webgpu.h from webgpu-native/webgpu-headers@${HEADERS_TAG}"
mkdir -p "$THIRD/webgpu-headers"
curl -fsSL "https://raw.githubusercontent.com/webgpu-native/webgpu-headers/${HEADERS_TAG}/webgpu.h" \
  -o "$THIRD/webgpu-headers/webgpu.h"

fetch_wgpu_native() {
  local plat="$1" dest="$2"
  local zip="$THIRD/wgpu-native-${plat}.zip"
  echo "==> wgpu-native ${WGPU_NATIVE_TAG} ${plat}"
  curl -fsSL -o "$zip" \
    "https://github.com/gfx-rs/wgpu-native/releases/download/${WGPU_NATIVE_TAG}/wgpu-${plat}-release.zip"
  mkdir -p "$THIRD/wgpu-native/$dest"
  unzip -qo "$zip" -d "$THIRD/wgpu-native/$dest"
  rm -f "$zip"
}

# Host / common targets. Override with FORCE_ALL=1 to pull every platform.
if [[ "$(uname -s)" == "Darwin" ]]; then
  fetch_wgpu_native "macos-aarch64" "macos-aarch64"
  fetch_wgpu_native "macos-x86_64" "macos-x86_64"
else
  fetch_wgpu_native "linux-x86_64" "linux-x86_64"
  fetch_wgpu_native "linux-aarch64" "linux-aarch64"
fi

# Always keep a Linux .so available (requested for non-cgo / CI use).
if [[ ! -f "$THIRD/wgpu-native/linux-x86_64/lib/libwgpu_native.so" ]]; then
  fetch_wgpu_native "linux-x86_64" "linux-x86_64"
fi

echo "==> done"
find "$THIRD" -type f \( -name 'webgpu.h' -o -name 'libwgpu_native.*' \) | sort
