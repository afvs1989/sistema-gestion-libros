# Sistema de Gestión de Libros

Avance del **Proyecto Integrador** (Autónomo 2 — POO) en **Go** con **arquitectura hexagonal** y principios **SOLID**. Incluye menú interactivo en consola.

> Avance significativo — no es la versión final.

## Requisitos

- [Go 1.22+](https://go.dev/dl/)

```bash
brew install go   # macOS
go version
```

## Ejecución

```bash
cd sistema-gestion-libros
go run ./cmd/main.go
```

Opción **13** carga datos de demostración.

```bash
go build -o sistema-libros ./cmd/main.go
./sistema-libros
```

## Menú de consola

| # | Acción | # | Acción |
|---|--------|---|--------|
| 1 | Registrar libro | 8 | Listar usuarios |
| 2 | Registrar revista | 9 | Prestar recurso |
| 3 | Registrar usuario | 10 | Devolver recurso |
| 4 | Buscar por título | 11 | Préstamos activos |
| 5 | Buscar por autor | 12 | Resumen catálogo |
| 6 | Listar libros | 13 | Cargar demo |
| 7 | Recursos disponibles | 0 | Salir |

## Estructura (SOLID + Hexagonal)

```
internal/
├── domain/
│   ├── entity/          # Una entidad por archivo
│   │   ├── recurso.go   # Interfaz Recurso
│   │   ├── autor.go
│   │   ├── libro.go
│   │   ├── revista.go
│   │   ├── usuario.go
│   │   └── prestamo.go
│   └── errores/         # Errores por contexto
├── port/
│   ├── inbound/         # Puertos primarios (ISP)
│   └── outbound/        # Puertos secundarios
├── application/         # Un servicio por responsabilidad (SRP)
│   ├── catalogo_servicio.go
│   ├── usuario_servicio.go
│   ├── prestamo_servicio.go
│   ├── consulta_servicio.go
│   └── gestor_biblioteca.go
└── adapter/
    ├── inbound/cli/     # Handlers separados por dominio
    └── outbound/persistence/
```

Documentación completa: [docs/ARQUITECTURA.md](docs/ARQUITECTURA.md)

## Principios SOLID

| Principio | Cómo se aplica |
|-----------|----------------|
| **SRP** | Un archivo por entidad/servicio/handler |
| **OCP** | Extensible con nuevos adaptadores |
| **LSP** | `Libro`/`Revista` implementan `Recurso` |
| **ISP** | Interfaces pequeñas por contexto |
| **DIP** | Servicios → puertos → adaptadores |

## Documentación

| Archivo | Descripción |
|---------|-------------|
| [docs/ARQUITECTURA.md](docs/ARQUITECTURA.md) | Hexagonal + SOLID |
| [docs/Autonomo2POOAFVS.pdf](docs/Autonomo2POOAFVS.pdf) | Informe Autónomo 2 |
| [docs/Autonomo1POOAFVS.pdf](docs/Autonomo1POOAFVS.pdf) | Planificación Autónomo 1 |

## Autor

**Valenzuela Saavedra Andrés Fernando** — UIDE, Ingeniería en Ciberseguridad

## Licencia

Proyecto académico. Uso educativo.
