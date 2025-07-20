{
  description = "Build config-lsp";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:tweag/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.utils.follows = "utils";
    };
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
        version = "0.3.0"; # CI:CD-VERSION
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            (final: prev: {
              go = prev.go_1_24;
              buildGoModule = prev.buildGo124Module;
            })
            gomod2nix.overlays.default
          ];
        };
        inputs = [
          pkgs.go_1_24
        ];
        serverUncompressed = pkgs.buildGoModule {
          nativeBuildInputs = inputs;
          pname = "github.com/Myzel394/config-lsp";
          version = version;
          src = ./server;
          vendorHash = "sha256-3hUSxiH9XUKcOWDFV4zyIVEVtltRJHfdzBGpGARTs9I";
          proxyVendor = true;
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

            # upx is currently not supported for darwin
            if [ "${system}" != "x86_64-darwin" ] && [ "${system}" != "aarch64-darwin" ]; then
              upx --ultra-brute $out/bin/config-lsp
            fi
          '';
        };
      in {
        packages = {
          # The server needs SENTRY_DSN to be injected using Go for a full prod build.
          # This is done in the CI:CD for build releases using 
          # `-ldflags="-X main.SENTRY_DSN=$SENTRY_DSN"`
          default = server;
          "server-uncompressed" = serverUncompressed;
          "vs-code-extension-bare" = let
            name = "config-lsp";
            node-modules = pkgs.mkYarnPackage {
              src = ./vs-code-extension;
              name = name;
              packageJSON = ./vs-code-extension/package.json;
              yarnLock = ./vs-code-extension/yarn.lock;
              yarnNix = ./vs-code-extension/yarn.nix;

              buildPhase = ''
                yarn --offline run compile:prod
              '';
              installPhase = ''
                mkdir -p extension

                # No idea why this is being created
                rm deps/${name}/config-lsp

                cp -rL deps/${name}/. extension

                mkdir -p $out
                cp -r extension/. $out
              '';
              distPhase = "true";

              buildInputs = [
                pkgs.vsce
              ];
            };
          in node-modules;
          "vs-code-extension" = let
            name = "config-lsp";
            node-modules = pkgs.mkYarnPackage {
              src = ./vs-code-extension;
              name = name;
              packageJSON = ./vs-code-extension/package.json;
              yarnLock = ./vs-code-extension/yarn.lock;
              yarnNix = ./vs-code-extension/yarn.nix;

              buildPhase = ''
                yarn --offline run compile:prod
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

        devShells.default = let 
          ourGopls = pkgs.gopls;
        in
          pkgs.mkShell {
            buildInputs = inputs ++ (with pkgs; [
              mailutils
              wireguard-tools
              antlr
              just
              ourGopls
              python3
            ]) ++ (if pkgs.stdenv.isLinux then with pkgs; [
              postfix
            ] else []);
          };

        devShells."vs-code-extension" = pkgs.mkShell {
          buildInputs = with pkgs; [
            just
            nodejs
            vsce
            yarn
            yarn2nix
          ];
        };
      }
    );
}
