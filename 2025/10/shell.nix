{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.z3
    pkgs.python3Packages.z3-solver
    pkgs.stdenv.cc.cc.lib
  ];

  shellHook = ''
    export LD_LIBRARY_PATH="${pkgs.z3}/lib:$LD_LIBRARY_PATH"
    export PYTHONPATH="${pkgs.python3Packages.z3-solver}/${pkgs.python3.sitePackages}:$PYTHONPATH"
    export Z3_LIBRARY_PATH="${pkgs.z3}/lib"

    echo "Z3 environment is ready."
  '';
}
