# Generate API stubs with Swagger/OpenAPI

Swagger (or OpenAPI) is awesome. Create your API spec file in YAML, then use
the `swagger-codegen` tool to generate server stubs for a variety of languages!

After generating the stubs, add edited files to `.swagger-codegen-ignore` (or
`.openapi-generator-ignore`) so they won't be ovewritten when the generator
command is re-run (for instance, to add a new endpoint stub).

## `swagger-codegen` vs `openapi-generator`

`openapi-generator` is a community fork of `swagger-codegen`. It has more
frequent updates and some better language support with some different design
choices. `swagger-codegen` can be used to convert a Swagger v2 spec to an
OpenAPI v3 spec, but `openapi-generator` doesn't have that feature.

## Examples

View available languages/templates:
```bash
swagger-codegen langs

# openapi-generator equivalent
openapi-generator list
```

Create go stubs in the current directory (and format `api/swagger.yaml`):
```bash
swagger-codegen -l go-server -i api/swagger.yaml

# openapi-generator equivalent
openapi-generator -g go-server -i api/openapi.yaml
```

Convert v2 to v3:
```bash
swagger-codegen openapi-yaml -i api/swagger.yaml
```

## Related

* https://github.com/swagger-api/swagger-codegen
* https://github.com/OpenAPITools/openapi-generator

