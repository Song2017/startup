#!/usr/bin/env sh

# configuration
input="nomad_z.yml"
pkg_path="com.samarkand.z"
output="tests/nomad_python_client"
type="python"
pkg_name="nomad_z_cli"
pkg_version="0.1.0"

source bin/generate_swagger_lib.sh
generate_swagger_lib $input $output $type $pkg_name $pkg_version