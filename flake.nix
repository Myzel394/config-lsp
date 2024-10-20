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
    utils.lib.eachSystem [
      "x86_64-linux"
      "aarch64-linux"

      "x86_64-darwin"
      "aarch64-darwin"

      "x86_64-windows"
      "aarch64-windows"
    ] (system: 
      let
        version = "0.1.0"; # CI:CD-VERSION
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
        serverUncompressed = pkgs.buildGoModule {
          nativeBuildInputs = inputs;
          pname = "github.com/Myzel394/config-lsp";
          version = version;
          src = ./server;
          vendorHash = "sha256-eO1eY+2XuOCd/dKwgFtu05+bnn/Cv8ZbUIwRjCwJF+U=";
          ldflags = [ "-s" "-w" ];
          checkPhase = ''
            go test -v $(pwd)/...
          '';
        };
        server = pkgs.stdenv.mkDerivation {
          name = "config-lsp-${version}";
          src = serverUncompressed;
          buildInputs = [
            pkgs.upx
          ];
          buildPhase = ''
            mkdir -p $out/bin
            cp $src/bin/config-lsp $out/bin/
            chmod +rw $out/bin/config-lsp
            upx --ultra-brute $out/bin/config-lsp
          '';
        };
      in {
        packages = {
          default = server;
          "vs-code-extension" = let
            name = "config-lsp";
            node-modules = pkgs.mkYarnPackage {
              src = ./vs-code-extension;
              name = name;
              packageJSON = ./vs-code-extension/package.json;
              yarnLock = ./vs-code-extension/yarn.lock;
              yarnNix = ./vs-code-extension/yarn.nix;

              buildPhase = ''
                yarn --offline run compile
              '';
              installPhase = ''
                mkdir -p extension

                # No idea why this is being created
                rm deps/${name}/config-lsp

                cp -rL deps/${name}/. extension
                cp ${server}/bin/config-lsp extension/out/config-lsp

                cd extension && ${pkgs.vsce}/bin/vsce package
                mkdir -p $out
                cp *.vsix $out
              '';
              distPhase = "true";

              buildInputs = [
                pkgs.vsce
              ];
            };
          in node-modules;
        };
        devShells.default = pkgs.mkShell {
          buildInputs = inputs ++ (with pkgs; [
            mailutils
            wireguard-tools
            antlr
          ]) ++ (if pkgs.stdenv.isLinux then with pkgs; [
            postfix
          ] else []);
        };
        devShells."vs-code-extension" = pkgs.mkShell {
          buildInputs = [
            pkgs.nodejs
            pkgs.vsce
            pkgs.yarn2nix
          ];
        };
      }
    );
}
