{
  description = "tinyFaaS-cli";

  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        tinyFaaSOverlay = final: prev: {
          tinyFaaS-cli = self.packages.${system}.tinyFaaS-cli;
        };
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShell = import ./shell.nix { inherit pkgs; };
        packages = {
          tinyFaaS-cli = import ./package.nix { inherit pkgs; };
          default = self.packages.${system}.tinyFaaS-cli;
        };
        overlay = tinyFaaSOverlay;
      });
}
