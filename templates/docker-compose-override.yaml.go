package templates

const ROOT_DOCKER_COMPOSE_OVERRIDE = `
{{ $config := . }}

version: '3.7'

services:
  
{{- range .ActorClasses}}

{{- if $config.HasAiByActorClass .Id }}
  {{.Id|kebabify}}:
    volumes:
      - ./:/app
{{end}}
{{end}}

  env:
    volumes:
      - ./:/app
    environment: 
      - COGMENT_GRPC_REFLECTION=1
      - PYTHONUNBUFFERED=1

  orchestrator:
    volumes:
      - .:/app
    ports:
      - "9000:9000"

  client:
    volumes:
      - ./:/app

# log exporter with postgres db - need to add lan IP address
# 
#  setup_db:
#      environment:
#        - IP= # Please specify your computer's LAN!
#  truncate_db:
#      environment:
#        - IP= # Please specify your computer's LAN!

`
