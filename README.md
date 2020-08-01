# minecraftboxapi-twitch

**minecraftboxapi-twitch** es un ejemplo de como usar golang y el API de Twitch para listar streams referentes a un videojuego.

Actualmente se puede ver un ejemplo de este servicio en la sección de streaming de [Minecraftbox](https://minecraftbox.net/).

## Requerimientos

- Tener una cuenta en Twitch.
- Crear una app en Twitch.

## Configuración desarrollo

Los siguientes parámetros son los que se pueden usar en el flujo.

  *Todos son obligatorios*

Parámetro | Descripción
--------- | -----------
`CLIENTID` | ID del cliente de la app de Twitch.
`CLIENTSECRET` | Secret de la app de Twitch.
`ENVAPP` | Environment en el que se va a trabajar.
`APPPORT` | Puerto en el que se ejecutara el servicio.

## Configuración servicio

Los siguientes parámetros son los que se pueden usar en el flujo.

  *Todos son obligatorios*

Parámetro | Descripción
--------- | -----------
`game` | ID del juego de Twitch.
`lang` | Idioma.

### Ejemplo

1. `go run api.go`

2. `http://localhost/api/streaming/twitch?game=27471&lang=es`

## Adicional

Son bienvenido a sugerir cambios o proponer nuevas características. La idea es extender la funcionalidad de este servicio.
