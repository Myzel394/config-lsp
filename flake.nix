{
  description = "Build config-lsp";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    gomod2nix = {
      url = "github:tweag/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.utils.follows = "utils";
    };
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils, gomod2nix }: 
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            (final: prev: {
              go = prev.go_1_22;
              buildGoModule = prev.buildGo122Module;
            })
            gomod2nix.overlays.default
          ];
        };
        inputs = [
          pkgs.go_1_22
        ];
      in {
        packages = {
          default = pkgs.buildGoModule {
            nativeBuildInputs = inputs;
            pname = "github.com/Myzel394/config-lsp";
            version = "v0.0.1";
            src = ./.;
            vendorHash = "sha256-PUVmhdbmfy1FaSzLt3SIK0MtWezsjb+PL+Z5YxMMhdw=";
            checkPhase = ''
              go test -v $(pwd)/...
            '';
          };
        };
        devShells.default = pkgs.mkShell {
          buildInputs = inputs ++ (with pkgs; [
            mailutils
            postfix
            wireguard-tools
          ]);
        };
      }
    );
}
