{ pkgs ? import <nixpkgs> {} }:

with pkgs; pkgs.mkShell rec {
  buildInputs = with pkgs; [
    alsa-lib
    xorg.libX11
    # xorg.libX11.dev
    libGL

    raylib
    glfw
    xorg.libXrandr
    xorg.libXinerama
    xorg.libXcursor
    xorg.libXi
    pkg-config

    wayland
    wayland-protocols
    libxkbcommon
    glfw-wayland
  ];
  LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath buildInputs;
}
