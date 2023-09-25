{ pkgs }:
with pkgs;
mkShell {
  name = "tinyFaaS-cli-shell";

  buildInputs = [
    go

    nixfmt

    tcpdump
  ];

  shellHook = ''
    # ...
  '';
}
