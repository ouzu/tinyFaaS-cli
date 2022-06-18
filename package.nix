{ pkgs ? import <nixpkgs> { } }:
with pkgs;
buildGoModule {
  name = "tinyFaaS-cli";
  src = lib.cleanSource ./.;
  vendorSha256 = "sha256-+84USjFcyOgk6hP6yzAPP2MQzo8Fb+GOI07j1qYmDBw=";
}
