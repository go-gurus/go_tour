<!-- .slide: data-background="img/GOLANG_VS_JAVA_SPRINGBOOT/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## golang vs. Java Spring Boot

In this chapter we create 2 services with identical content.
One is built in Golang the other using Spring Boot.
After the services are fully running and close to production, we will contrast them using various metrics.

----

### Complete Source Code
* [github.com/go-gurus/golang_vs_java_springboot](https://github.com/go-gurus/golang_vs_java_springboot)

----

### prerequisits
* golang
* Java 17

----

### Spring Boot Beer Fridge
* use the special link to spring `initializr` [start.spring.io/...](https://start.spring.io/#!type=maven-project&language=java&platformVersion=3.0.4&packaging=jar&jvmVersion=17&groupId=io.grohm&artifactId=beerfridge&name=beerfridge&description=Demo%20project%20for%20the%20big%20comparison%20between%20Golang%20and%20Java%20Spring%20Boot%20in%20the%20GO-Tour%20workshop.&packageName=io.grohm.beerfridge&dependencies=web,devtools,lombok,data-jpa,liquibase,postgresql) to init the project
* unpack the sources
* install dependencies
```shell
cd beerfridge
./mvnw clean install
```

----

* we will get an error message like this
```shell
[ERROR] Errors: 
[ERROR]   BeerfridgeApplicationTests.contextLoads » IllegalState Failed to load Applicat...
[INFO] 
[ERROR] Tests run: 1, Failures: 0, Errors: 1, Skipped: 0
...
[ERROR] Failed to execute goal org.apache.maven.plugins:maven-surefire-plugin:2.22.2:test (default-test) on project beerfridge: There are test failures.
[ERROR] 
[ERROR] Please refer to /Users/grohmio/repos/github/go-gurus/golang_vs_java_springboot/springboot/beerfridge/target/surefire-reports for the individual test results.
...
```
----

* the application context could not be loaded in the test, because the DB connection is missing
* lets fix that
```shell
# docker-compose.yml
version: '3.1'

services:
    beerfridge_db:
        image: postgres
        restart: unless-stopped
        ports:
            - "5432:5432"
        environment:
            POSTGRES_USER: beerfridge_backend
            POSTGRES_PASSWORD: beerfridge_db_pw
            POSTGRES_DB: beerfridge_backend_db
        volumes:
            - beerfridge_db_data:/var/lib/postgresql/data
        networks:
            - beerfridge_db_net

volumes:
    beerfridge_db_data:

networks:
    beerfridge_db_net:

```
----

* configure the application context

```yaml
# springboot/beerfridge/src/main/resources/application.properties
### datasource
spring.datasource.url=jdbc:postgresql://${DATABASE_HOST:localhost}:${DATABASE_PORT:5432}/${DATABASE_NAME:beerfridge_backend_db}
spring.datasource.username=${DATABASE_USER:beerfridge_backend}
spring.datasource.password=${DATABASE_PASSWORD:beerfridge_db_pw}
```
----

* next build an initial changeset for liquibase

```yaml
# springboot/beerfridge/src/main/resources/application.properties
databaseChangeLog:
  - changeSet:
      id: 2023-03-07-init-database
      author: Andreas Grohmann
      changes:
        sqlFile:
          dbms: postgresql
          encoding: utf-8
          path: postgres/2023-03-07-init-database.sql
          relativeToChangelogFile: true
```

* generate an empty file for now
```yaml
# springboot/beerfridge/src/main/resources/db/postgres/2023-03-07-init-database.sql

```

----

* lets setup the db in a new terminal and try again with no errors expected

```shell
docker-compose up beerfridge_db
```
```shell
./mvnw clean install
```

----

* now lets speed up things a little bit by using code generation, add the swagger code gen maven plugin

```xml
<!-- springboot/beerfridge/pom.xml -->
...
	<build>
		<plugins>
    ...
    <plugin>
				<groupId>org.openapitools</groupId>
				<artifactId>openapi-generator-maven-plugin</artifactId>
				<version>6.4.0</version>
				<executions>
					<execution>
						<goals>
							<goal>generate</goal>
						</goals>
						<configuration>
							<inputSpec>
								${project.basedir}/src/main/resources/openapi.yml
							</inputSpec>
							<generatorName>spring</generatorName>
							<apiPackage>io.grohm.beerfridge.SwaggerCodgen.api</apiPackage>
							<modelPackage>io.grohm.beerfridge.SwaggerCodgen.model</modelPackage>
						</configuration>
					</execution>
				</executions>
			</plugin>
		</plugins>
	</build>
</project>
```

----

* lets add the openapi interface file using [editor.swagger.io/](https://editor.swagger.io/) and the `swagger.yml` from our previous section

```yaml
# src/main/resources/openapi.yml
openapi: 3.0.1
info:
  title: A beer fridge service
  description: Beer fridge service build with openapi-generator
  version: 1.0.0
servers:
- url: /
paths:
  /beers:
    get:
      tags:
      - beers
      operationId: getAllBeers
      parameters:
      - name: limit
        in: query
        schema:
          type: integer
          format: int32
          default: 20
      responses:
        200:
          description: list the beer operations
          content:
            application/io.grohm.go-workshop.beer-fridge.v1+json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/beer'
    post:
      tags:
      - beers
      operationId: addOne
      requestBody:
        content:
          application/io.grohm.go-workshop.beer-fridge.v1+json:
            schema:
              $ref: '#/components/schemas/beer'
        required: false
      responses:
        201:
          description: Created
          content:
            application/io.grohm.go-workshop.beer-fridge.v1+json:
              schema:
                $ref: '#/components/schemas/beer'
        default:
          description: error
          content:
            application/io.grohm.go-workshop.beer-fridge.v1+json:
              schema:
                $ref: '#/components/schemas/error'
      x-codegen-request-body-name: body
  /beers/{id}:
    delete:
      tags:
      - beers
      operationId: destroyOne
      parameters:
      - name: id
        in: path
        required: true
        schema:
          type: integer
          format: int64
      responses:
        204:
          description: Deleted
          content: {}
        default:
          description: error
          content:
            application/io.grohm.go-workshop.beer-fridge.v1+json:
              schema:
                $ref: '#/components/schemas/error'
  /temperature:
    get:
      tags:
      - fridge
      operationId: getTemperature
      responses:
        200:
          description: return the current fridge temperature
          content:
            application/io.grohm.go-workshop.beer-fridge.v1+json:
              schema:
                $ref: '#/components/schemas/temperature'
components:
  schemas:
    beer:
      required:
      - origin
      - title
      - volume-percentage
      type: object
      properties:
        id:
          type: integer
          format: int64
          readOnly: true
        title:
          minLength: 1
          type: string
        origin:
          minLength: 1
          type: string
        volume-percentage:
          type: number
          format: float
    temperature:
      type: integer
      format: int64
    error:
      required:
      - message
      type: object
      properties:
        code:
          type: integer
          format: int64
        message:
          type: string

```

----

### What we have learned
* TBD
* TBD

----

### Further readings

* Swagger Editor
  * [editor-next.swagger.io](https://editor-next.swagger.io/)
  * [editor.swagger.io/](https://editor.swagger.io/)
* openapi-generator
  * [github.com/OpenAPITools/openapi-generator](https://github.com/OpenAPITools/openapi-generator)
  * [github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/spring.md](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/spring.md)
* openapi-generator-maven-plugin
  * [github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator-maven-plugin](https://github.com/OpenAPITools/openapi-generator/tree/master/modules/openapi-generator-maven-plugin)

---