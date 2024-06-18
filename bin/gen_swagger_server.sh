#!/usr/bin/env sh

# configuration
input="order.yml"
output="swagger-server"
type="go-gin-server"
pkg_name="order_service"
pkg_version="0.1.0"

source ./bin/generate_swagger_lib.sh
generate_swagger_lib $input $output $type $pkg_name $pkg_version

# sed -e 's/sw ".\/order_service"/sw "server\/swagger-server\/order_service"/' $output/main.go >$output/main.go.tmp
# mv -f $output/main.go.tmp $output/main.go
# sed -e 's/	Message string/	Message interface{}/' $output/order_service/model_pilot_response.go >$output/order_service/model_pilot_response.go.tmp
# mv -f $output/order_service/model_pilot_response.go.tmp $output/order_service/model_pilot_response.go

# sed -e 's/address2,omitempty/address2/' $output/order_service/model_address.go >$output/order_service/model_address.go.tmp
# mv -f $output/order_service/model_address.go.tmp $output/order_service/model_address.go
# sed -e 's/lastName,omitempty/lastName/' $output/order_service/model_address.go >$output/order_service/model_address.go.tmp
# mv -f $output/order_service/model_address.go.tmp $output/order_service/model_address.go 

# array=("model_item" "model_customer" "model_odoo_response" "model_odoo_order_dto" "model_product")  # "model_order"
# for element in "${array[@]}"  
# do  
#     echo "$element remove omitempty"  
#     sed -e 's/,omitempty//' $output/order_service/$element.go >$output/order_service/$element.go.tmp
#     mv -f $output/order_service/$element.go.tmp $output/order_service/$element.go    
# done

# swagger file
# rm -f ./resources/openapi.yaml
# cp -f $output/api/openapi.yaml ./resources


  echo "${green}generate server Done!${reset}"