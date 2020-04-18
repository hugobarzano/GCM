package commons

const PropertyTemplate = `{{range $key, $val := .}}
      {{$key}}:
        type: {{$val}}
{{end}}`

const ModelPropertiesTemplate = `{{define "ModelProperties"}}
    type: object
    required:
      - id
    properties:
      id:
        type: string
        description: Id of the element
      createdAt:
        type: string
      updatedAt:
        type: string
      {{.}}
{{end}}`

const SpecTemplate = `
{{define "Spec"}}
swagger: "2.0"
info:
  description: This is the swagger file that defines generated api
  version: "dev/0.0.1"
  title: {{.Title}}
consumes:
  - application/json
produces:
  - application/json

definitions:
  {{.ModelName}}:{{.ModelData}}

#base path for server application
basePath: /api
# Paths supported by the server application
paths:
  /{{.ApiName}}:
    get:
      operationId: {{.ApiPkg}}listAll
      tags:
        - {{.ModelName}}
      summary: Read the entire list of {{.ApiName}}
      description: Read the list of {{.ApiName}}
      responses:
        200:
          description: Successfully read {{.ApiName}} list operation
          schema:
            $ref: '#/definitions/{{.ModelName}}'
    post:
      operationId: {{.ApiPkg}}createOne
      tags:
        - {{.ModelName}}
      summary: Create a {{.ModelName}} and add it to the {{.ApiName}} list
      description: Create a new {{.ModelName}} in the {{.ApiName}} list
      parameters:
        - name: object
          in: body
          description: {{.ModelName}} to create
          required: True
          schema:
            $ref: '#/definitions/{{.ModelName}}'
      responses:
        201:
          description: Successfully created {{.ModelName}} in list

  /{{.ApiName}}/{id}:
    get:
      operationId: {{.ApiPkg}}getOne
      tags:
        - {{.ModelName}} 
      summary: Read one {{.ModelName}} from the {{.ApiName}} list
      description: Read one {{.ModelName}} from the {{.ApiName}} list
      parameters:
        - name: id
          in: path
          description: Unique model identifier for {{.ModelName}} 
          type: string
          required: True
      responses:
        200:
          description: Successfully read person from people list operation
          schema:
            $ref: '#/definitions/{{.ModelName}}'
    put:
      operationId: {{.ApiPkg}}updateOne
      tags:
        - {{.ModelName}}
      summary: Update a {{.ModelName}} in {{.ApiName}} 
      description: Update a {{.ModelName}} in {{.ApiName}} 
      parameters:
        - name: id
          in: path
          description: Unique model identifier for {{.ModelName}}
          type: string
          required: True
        - name: object
          in: body
          schema:
            $ref: '#/definitions/{{.ModelName}}'
      responses:
        200:
          description: Successfully updated {{.ModelName}}
    delete:
      operationId: {{.ApiPkg}}deleteOne
      tags:
        - {{.ModelName}}
      summary: Delete a {{.ModelName}} from the {{.ApiName}} list
      description: Delete a person
      parameters:
        - name: id
          in: path
          type: string
          required: True
          description: Unique model identifier for {{.ModelName}}
      responses:
        200:
          description: Successfully deleted a {{.ModelName}}
{{end}}`
