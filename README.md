# Docker Credentail helper for Azure Container Registry

## Build Windows

```shell 
# windows
go build -o docker-credential-acr.exe ./go

# linux
go build -o docker-credential-acr ./go
```

Add exec to your path variable 

update .docker\config.json
[Docker Documentation](https://docs.docker.com/reference/cli/docker/login/#credential-helpers)  
eg.
```json
{
	"auths": {},
	"credStore": "desktop",	
	"credHelpers": {
		"registry.azurecr.io": "acr"
	},
	"currentContext": "default"
}
```
