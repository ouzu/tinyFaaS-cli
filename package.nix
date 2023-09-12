{ pkgs ? import <nixpkgs> { } }:
with pkgs;
buildGoModule {
  name = "tinyFaaS-cli";
  src = lib.cleanSource ./.;
  vendorSha256 = "sha256-bgsPI5FNK2AECm29cYxfyUFF/E3/k6XgaoTUhbPvvXk=";
}
