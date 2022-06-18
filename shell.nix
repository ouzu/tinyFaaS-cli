{ pkgs ? import <nixpkgs> { } }:
with pkgs;
mkShell {
  name = "tinyFaaS-cli-shell";

  buildInputs = [
    go

    nixfmt

    tcpdump
    wireshark
  ];

  shellHook = ''
    # ...
  '';
}
