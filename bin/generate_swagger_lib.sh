#!/usr/bin/env sh

set -e

if hash tput 2>/dev/null; then
  red=`tput setaf 1`
  green=`tput setaf 2`
  reset=`tput sgr0`
else
  red=""
  green=""
  reset=""
fi

docker_image="openapitools/openapi-generator-cli:v6.0.0"

function generate_swagger_lib() {
  input=$1
  output=$2
  type=$3
  pkg_name=$4
  pkg_version=$5

  echo "${green}Step1===${reset} remove old '$output'"
  rm -rf $output

  echo "${green}Step2===${reset} generate new template code based on $input"

  validate_cmd="docker run --rm -v ${PWD}:/local $docker_image  validate -i /local/$input"
  echo "RUN --- $validate_cmd"
  $validate_cmd

  gen_cmd="docker run --rm -v ${PWD}:/local $docker_image generate \
    --input-spec /local/$input \
    --generator-name $type \
    --output /local/$output"

  gen_cmd="$gen_cmd -p apiPath=$pkg_name -p packageName=$pkg_name -p packageVersion=$pkg_version"
  gen_cmd="$gen_cmd -p additionalProperties=true"

  echo "RUN --- $gen_cmd"
  $gen_cmd

  # Move "swagger_server/swagger_client" to the root directory so that we could import them easily.
  # And we don't need extra step to tell IDE how to locate source code directory.
  # rm -rf $pkg_name
  # cp -rv $output/$pkg_name $pkg_name
  #rm -rf $output

  echo "${green}generate swagger lib Done!${reset}"
}
