# prueba_tecnica_golang

# User Service API

Este proyecto es una API de servicio de usuario simple escrita en Go. Proporciona funcionalidades básicas para registrar usuarios, hacer login y verificar la existencia de un usuario por su correo electrónico, teléfono o nombre de usuario.

## Características

- Registro de usuario.
- Login de usuario.
- Validación de datos de entrada.
- Verificación de la existencia del correo electrónico, teléfono y nombre de usuario.
- Generación de token JWT para sesiones de login.

## Tecnologías Utilizadas

- Go
- Gorilla Mux (para enrutamiento)
- JWT para autenticación

## Requisitos

Para ejecutar este proyecto necesitas tener Go instalado en tu máquina. Visita [la página oficial de Go](https://golang.org/) para descargarlo e instalarlo.

## Configuración y Ejecución

1. **Clonar el Repositorio**

git clone https://github.com/RamonVP28/prueba_tecnica_golang.git


2. **Ejecutar el Servicio**

go run .

Esto lanzará el servidor en `localhost:8080` y estará listo para recibir peticiones.

## Uso

### Registrar un Usuario

Para registrar un usuario, envía una petición POST a `/register` con el siguiente cuerpo JSON:

```json
{
"username": "nuevoUsuario",
"email": "usuario@example.com",
"phone": "1234567890",
"password": "Password123!"
}
```

### Hacer Login

Para hacer login, envía una petición POST a `/login` con el siguiente cuerpo JSON:
```json
 {
  "login": "nuevoUsuario",
  "password": "Password123!"
 }
```

### Verificación de Disponibilidad de Usuario

Esta funcionalidad se ejecuta automáticamente al intentar registrar un nuevo usuario. Verifica si el correo electrónico, el teléfono o el nombre de usuario ya están en uso.

### Desarrollo y Pruebas

El proyecto incluye pruebas básicas. Para ejecutar las pruebas:

go test ./...

