# Arquitectura Hexagonal y SOLID

Este documento describe la organización del sistema aplicando **Arquitectura Hexagonal** y los **principios SOLID**.

## Principios SOLID aplicados

| Principio | Aplicación en el proyecto |
|-----------|---------------------------|
| **S** — Single Responsibility | Cada archivo/clase tiene una sola responsabilidad: una entidad por archivo, un servicio por dominio |
| **O** — Open/Closed | Nuevos adaptadores (PostgreSQL, API REST) sin modificar servicios existentes |
| **L** — Liskov Substitution | `Libro` y `Revista` sustituyen a `Recurso` en operaciones polimórficas |
| **I** — Interface Segregation | Puertos inbound/outbound divididos por contexto (catálogo, usuario, préstamo) |
| **D** — Dependency Inversion | Servicios dependen de interfaces (`port/outbound`), no de `Memoria` |

## Diagrama de capas

```
┌─────────────────────────────────────────────────────────┐
│  adapter/inbound/cli/     (Menú consola)                │
└───────────────────────────┬─────────────────────────────┘
                            │ port/inbound
┌───────────────────────────▼─────────────────────────────┐
│  application/                                           │
│  ├── catalogo_servicio.go                               │
│  ├── usuario_servicio.go                                │
│  ├── prestamo_servicio.go                               │
│  ├── consulta_servicio.go                               │
│  └── gestor_biblioteca.go  (fachada)                    │
└───────────────────────────┬─────────────────────────────┘
                            │ port/outbound
┌───────────────────────────▼─────────────────────────────┐
│  adapter/outbound/persistence/                          │
│  ├── memoria.go                                         │
│  ├── memoria_catalogo.go                                │
│  ├── memoria_usuarios.go                                │
│  └── memoria_prestamos.go                               │
└─────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────┐
│  domain/entity/  (núcleo — sin dependencias externas)   │
│  recurso.go · autor.go · libro.go · revista.go          │
│  usuario.go · prestamo.go                               │
└─────────────────────────────────────────────────────────┘
```

## Estructura de archivos

### Dominio (`internal/domain/`)

| Archivo | Responsabilidad |
|---------|-----------------|
| `entity/recurso.go` | Interfaz `Recurso` |
| `entity/autor.go` | Entidad `Autor` |
| `entity/libro.go` | Entidad `Libro` |
| `entity/revista.go` | Entidad `Revista` |
| `entity/usuario.go` | Entidad `Usuario` |
| `entity/prestamo.go` | Entidad `Prestamo` |
| `errores/catalogo.go` | Errores del catálogo |
| `errores/usuario.go` | Errores de usuarios |
| `errores/prestamo.go` | Errores de préstamos |
| `errores/validacion.go` | Errores de validación |

### Puertos (`internal/port/`)

**Inbound (driving):**

| Archivo | Interfaz |
|---------|----------|
| `inbound/catalogo.go` | `CatalogoCasosUso` |
| `inbound/usuario.go` | `UsuarioCasosUso` |
| `inbound/prestamo.go` | `PrestamoCasosUso` |
| `inbound/biblioteca.go` | `ConsultaBiblioteca` + `Biblioteca` (composición) |

**Outbound (driven):**

| Archivo | Interfaz |
|---------|----------|
| `outbound/catalogo.go` | `Catalogo` |
| `outbound/usuarios.go` | `Usuarios` |
| `outbound/prestamos.go` | `Prestamos` |

### Aplicación (`internal/application/`)

| Archivo | Responsabilidad |
|---------|-----------------|
| `catalogo_servicio.go` | Registro y búsqueda de libros/recursos |
| `usuario_servicio.go` | Gestión de lectores |
| `prestamo_servicio.go` | Préstamos y devoluciones |
| `consulta_servicio.go` | Resúmenes e informes |
| `gestor_biblioteca.go` | Fachada que delega a los servicios |

### Adaptadores (`internal/adapter/`)

**CLI (inbound):**

| Archivo | Responsabilidad |
|---------|-----------------|
| `cli/menu.go` | Bucle principal y enrutamiento |
| `cli/handlers_catalogo.go` | Acciones de catálogo |
| `cli/handlers_usuario.go` | Acciones de usuarios |
| `cli/handlers_prestamo.go` | Acciones de préstamos |
| `cli/demo.go` | Datos de demostración |

**Persistencia (outbound):**

| Archivo | Responsabilidad |
|---------|-----------------|
| `persistence/memoria.go` | Struct e inicialización |
| `persistence/memoria_catalogo.go` | Operaciones de catálogo |
| `persistence/memoria_usuarios.go` | Operaciones de usuarios |
| `persistence/memoria_prestamos.go` | Operaciones de préstamos |

## Flujo de dependencias

```
cmd/main.go
  → adapter/inbound/cli
  → application/gestor_biblioteca
  → application/*_servicio
  → port/outbound (interfaces)
  → adapter/outbound/persistence
  → domain/entity
```

Las dependencias **siempre apuntan hacia el dominio**. El dominio no importa capas externas.

## Ensamblado en `main.go`

```go
repo := persistence.NuevaMemoria()

catalogoSvc := application.NuevoCatalogoServicio(repo)
usuarioSvc := application.NuevoUsuarioServicio(repo)
prestamoSvc := application.NuevoPrestamoServicio(repo, repo, repo)
consultaSvc := application.NuevoConsultaServicio("Biblioteca UIDE", repo, repo, catalogoSvc)

servicio := application.NuevoGestorBiblioteca(catalogoSvc, usuarioSvc, prestamoSvc, consultaSvc)
cli.NuevoMenu(servicio).Ejecutar()
```

## Evolución futura

Gracias a ISP y DIP, se pueden agregar sin romper el núcleo:

- `adapter/inbound/http/` — API REST (solo implementa puertos inbound)
- `adapter/outbound/persistence/postgres/` — Base de datos (solo implementa puertos outbound)

## Documentos relacionados

| Archivo | Descripción |
|---------|-------------|
| `docs/Autonomo1POOAFVS.pdf` | Planificación Etapa 1 |
| `docs/Autonomo2POOAFVS.pdf` | Informe Autónomo 2 |
| `README.md` | Guía de uso |
