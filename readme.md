# SimpleRedirect for Traefik
This plugin provides a middleware that can perform redirects without a
backend server.

## Installation via static configuration
```yaml
experimental:
  plugins:
    simpleRedirect:
      moduleName: "github.com/daniels0056/traefik-simpleredirect"
      version: "v1.0.0"
```

## Example usage configuration
The following request redirects all users accessing the host `outdated.example.com` to `google.com`.

```yaml
http:
  middlewares:
    redirect:
      plugin:
        simpleRedirect:
          redirectTo: "https://google.com"
          redirectCode: 302
          
  routers:
    outdated-host:
      rule: "Host(`outdated.example.com`)"
      service: noop@internal
      middlewares:
        - redirect@file
```
