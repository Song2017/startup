#!/usr/bin/env sh

# configuration
input="nomad_red_bbc.yml"
pkg_path="com.samarkand.red.bbc"
output="tests/nomad_java_client"
type="java"
pkg_name="nomad_red_bbc_cli"
pkg_version="0.1.0"

# https://openapi-generator.tech/docs/generators/java
docker_image="openapitools/openapi-generator-cli:v5.0.0-beta"

function generate_swagger_lib() {
  input=$1
  output=$2
  type=$3
  pkg_name=$4
  pkg_version=$5
  pkg_path=$6

  echo "${green}Step1===${reset} remove old '$output'"
  rm -rf $output

  echo "${green}Step2===${reset} generate new template code based on $input"

  validate_cmd="docker run -v ${PWD}:/local $docker_image  validate -i /local/$input"
  echo "RUN --- $validate_cmd"
  $validate_cmd

  gen_cmd="docker run -v ${PWD}:/local $docker_image generate \
    --input-spec /local/$input \
    --generator-name $type \
    --output /local/$output"

  # Lang: Java
  # bash string manipulation: http://tldp.org/LDP/abs/html/string-manipulation.html
  artifact_id="${pkg_name//_/-}"
  echo $artifact_id
  gen_cmd="$gen_cmd -p artifactId=$artifact_id -p artifactVersion=$pkg_version"
  gen_cmd="$gen_cmd -p groupId=com.gitlab.samarkand-nomad -p invokerPackage=$pkg_path"
  gen_cmd="$gen_cmd -p apiPackage=$pkg_path.api -p modelPackage=$pkg_path.model"

  echo "RUN --- $gen_cmd"
  $gen_cmd
  # deploy package to maven staging repo https://oss.sonatype.org/#welcome
  sed -i -e 's?</project>? ?' $output/pom.xml
  # add .asc files: https://central.sonatype.org/pages/apache-maven.html#gpg-signed-components
  export temp="<plugin><groupId>org.sonatype.plugins</groupId><artifactId>nexus-staging-maven-plugin</artifactId><version>1.6.7</version><extensions>true</extensions><configuration><serverId>ossrh</serverId><nexusUrl>https://oss.sonatype.org/</nexusUrl><autoReleaseAfterClose>true</autoReleaseAfterClose></configuration></plugin></plugins>"
  sed -i -e "s|</plugins>|$temp|" $output/pom.xml
  export temp='<plugin><groupId>org.apache.maven.plugins</groupId><artifactId>maven-gpg-plugin</artifactId><version>1.5</version><executions><execution><id>sign-artifacts</id><phase>verify</phase><goals><goal>sign</goal></goals></execution></executions></plugin></plugins>'
  sed -i -e "s|</plugins>|$temp|" $output/pom.xml
  echo  "
    <distributionManagement>
        <snapshotRepository>
            <id>ossrh</id>
            <url>https://oss.sonatype.org/content/repositories/snapshots/</url>
        </snapshotRepository>
        <repository>
            <id>ossrh</id>
            <url>
                https://oss.sonatype.org/service/local/staging/deploy/maven2/
            </url>
        </repository>
    </distributionManagement>
</project>
  " >> $output/pom.xml

    # please find OSSRH_USER/OSSRH_PASS in Pilot CI variables,
    #   and replace your local mvn setting.xml
    # publish java client package need gpg keys:
    #   2576741AA75957B75B171DBC950CE44AFA40A229 has sent to key server
    # pub   rsa2048 2020-06-03 [SC] [有效至：2022-06-03]
    echo  "<!--maven connect nexus need user and password-->
<settings>
    <servers>
        <server>
            <id>ossrh</id>
            <username>${OSSRH_USER}</username>
            <password>${OSSRH_PASS}</password>
        </server>
    </servers>

    <profiles>
        <profile>
            <id>ossrh</id>
            <activation>
                <activeByDefault>true</activeByDefault>
            </activation>
            <properties>
                <gpg.passphrase>2576741AA75957B75B171DBC950CE44AFA40A229
                </gpg.passphrase>
            </properties>
        </profile>
    </profiles>
</settings>
  " > $output/settings.xml

  pkg_test_path="${pkg_path//.//}"
  echo $pkg_test_path
  echo  "/*
 * Nomad Pilot
 */
package $pkg_path.api;

import $pkg_path.ApiException;
import $pkg_path.ApiClient;
import $pkg_path.model.ModelApiResponse;
import org.junit.Test;
import org.junit.Assert;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
" > $output/src/test/java/$pkg_test_path/api/ShipApiWithAppCodeTest.java

  echo  '
/**
 * API tests for ShipApiWithAppCodeTest
 */
public class ShipApiWithAppCodeTest {

    private final ShipApi api = new ShipApi();

    /**
     * queryShip
     */
    @Test
    public void queryShipTest() throws ApiException {
        ApiClient defaultClient = api.getApiClient();
        defaultClient.addDefaultHeader("Authorization","APPCODE APP_CODE_VALUE"); // replace APP_CODE_VALUE with the app code
        api.setApiClient(defaultClient);
        String carrier = "samarkand.sfexpress.prod"; // charge us
        String orderRef = "E202004301638260794000019610GZ";
        String sellerOrderRef = null;
        ModelApiResponse response = api.queryShip(carrier, orderRef, sellerOrderRef);

        System.out.println(response);
        Assert.assertEquals("Should be 200 OK", (long) 200, (long) response.getCode());
    }
}' >> $output/src/test/java/$pkg_test_path/api/ShipApiWithAppCodeTest.java

}

generate_swagger_lib $input $output $type $pkg_name $pkg_version $pkg_path
