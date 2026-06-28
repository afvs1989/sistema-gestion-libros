# 📚 Sistema de Gestión de Libros — Biblioteca UIDE

**Proyecto Integrador Final — Programación Orientada a Objetos (POO)**
Aplicación full‑stack para la gestión de una biblioteca: catálogo de libros y revistas,
registro de lectores y control de préstamos/devoluciones, con **API REST**, **autenticación
JWT**, persistencia **Code First** sobre **SQL Server** y un **frontend Angular**.

---

## 🧑‍💻 Datos del grupo

| Campo | Detalle |
|-------|---------|
| **Integrante** | Valenzuela Saavedra Andrés Fernando |
| **Carrera** | Ingeniería en Ciberseguridad — UIDE |
| **Asignatura** | Programación Orientada a Objetos (POO) |
| **Entregable** | Proyecto Final (integración de las 8 semanas) |
| **Fecha** | 28 de junio de 2026 |

---

## 🎯 Objetivo del programa

Construir un sistema de gestión bibliotecaria que **integre los conocimientos de las 4 unidades**
de la asignatura, aplicando **Programación Orientada a Objetos**, **principios SOLID** y
**arquitectura hexagonal (puertos y adaptadores)**, exponiendo la lógica de negocio mediante
**servicios web REST** protegidos con **JWT**, con persistencia **Code First** en una base de
datos **SQL Server** (contenedor Docker) y una interfaz de usuario **Angular**.

---

## 🏗️ Arquitectura general

```
┌──────────────────────────┐        HTTP/JSON + JWT        ┌───────────────────────────┐
│   FRONTEND (Angular 22)   │  ───────────────────────────▶ │     BACKEND (Go 1.23)     │
│  - Login / token JWT      │ ◀───────────────────────────  │  Arquitectura Hexagonal   │
│  - Catálogo, usuarios,    │            JSON               │  + SOLID                  │
│    préstamos              │                               └────────────┬──────────────┘
└──────────────────────────┘                                            │ GORM (Code First)
                                                                        ▼
                                                            ┌───────────────────────────┐
                                                            │   SQL Server (Docker)     │
                                                            │   base de datos "biblioteca" │
                                                            └───────────────────────────┘
```

El repositorio está organizado en **dos carpetas**:

```
sistema-gestion-libros/
├── backend/     # API REST en Go (hexagonal + SOLID + GORM + JWT)
├── frontend/    # SPA en Angular 22 (login JWT + CRUD)
├── docs/        # Documentación e informes
└── README.md
```

### Backend — capas hexagonales

```
backend/internal/
├── domain/                 # Núcleo: entidades y reglas de negocio (sin dependencias externas)
│   ├── entity/             # Libro, Revista, Usuario, Préstamo, Autor, Cuenta, Recurso
│   └── errores/            # Errores de dominio por contexto
├── port/
│   ├── inbound/            # Puertos primarios (casos de uso: catálogo, usuario, préstamo, auth)
│   └── outbound/           # Puertos secundarios (persistencia, hasher)
├── application/            # Servicios de aplicación (un servicio por responsabilidad — SRP)
└── adapter/
    ├── inbound/
    │   ├── cli/            # Adaptador de consola (menú interactivo original)
    │   └── http/           # Adaptador REST con Gin + JWT + DTOs JSON  ← NUEVO
    └── outbound/
        ├── persistence/    # Adaptador en RAM (memoria_*.go, usado por la CLI)
        │   └── sqlserver/  # Adaptador GORM/SQL Server (Code First)   ← NUEVO
        └── seguridad/      # Hasher bcrypt                            ← NUEVO
```

> El mismo núcleo de dominio se reutiliza para **dos adaptadores de entrada** (CLI y API REST)
> y **dos adaptadores de persistencia** (memoria y SQL Server), demostrando el principio
> **Abierto/Cerrado**: se añade infraestructura nueva sin modificar la lógica de negocio.

---

## 🔌 Servicios web (REST)

Se exponen **15 servicios web** (el requisito mínimo era 8). La serialización de entrada y
salida se realiza mediante **JSON**.

| # | Método | Ruta | Descripción | Protegido |
|---|--------|------|-------------|-----------|
| 1 | POST | `/api/auth/register` | Registrar cuenta de acceso | Público |
| 2 | POST | `/api/auth/login` | Iniciar sesión y emitir token JWT | Público |
| 3 | GET | `/api/libros` | Listar catálogo de libros | 🔒 JWT |
| 4 | POST | `/api/libros` | Registrar un libro | 🔒 JWT |
| 5 | GET | `/api/libros/buscar?titulo=&autor=` | Buscar libros por título o autor | 🔒 JWT |
| 6 | GET | `/api/libros/:id` | Obtener un libro por ID | 🔒 JWT |
| 7 | DELETE | `/api/libros/:id` | Eliminar un libro | 🔒 JWT |
| 8 | POST | `/api/revistas` | Registrar una revista | 🔒 JWT |
| 9 | GET | `/api/recursos/disponibles` | Recursos disponibles para préstamo | 🔒 JWT |
| 10 | GET | `/api/usuarios` | Listar lectores | 🔒 JWT |
| 11 | POST | `/api/usuarios` | Registrar lector | 🔒 JWT |
| 12 | POST | `/api/prestamos` | Prestar un recurso a un usuario | 🔒 JWT |
| 13 | PUT | `/api/prestamos/devolver/:recursoId` | Devolver un recurso | 🔒 JWT |
| 14 | GET | `/api/prestamos/activos` | Listar préstamos activos | 🔒 JWT |
| 15 | GET | `/api/catalogo/resumen` | Resumen agregado del sistema | 🔒 JWT |

Las rutas protegidas requieren la cabecera `Authorization: Bearer <token>`.

### 📖 Documentación interactiva (Swagger / OpenAPI)

La API está documentada con **Swagger (OpenAPI 2.0)** mediante `swaggo/swag`. Con el backend
en marcha, abre:

**http://localhost:8081/swagger/index.html**

Desde ahí puedes ver y **probar** cada servicio. Para usar las rutas protegidas:
1. Ejecuta `POST /auth/login` con `admin` / `admin123` y copia el `token`.
2. Pulsa **Authorize** (🔒) y escribe `Bearer <token>`.
3. Ya puedes invocar el resto de endpoints desde la propia UI.

La especificación se sirve en `http://localhost:8081/swagger/doc.json`. Si modificas las
anotaciones de los handlers, regenérala con:

```bash
cd backend
swag init -g cmd/api/main.go --parseDependency --parseInternal -o docs
```

> Requiere la CLI: `go install github.com/swaggo/swag/cmd/swag@latest`.

---

## ⚙️ Funcionalidades principales

- **Autenticación y seguridad:** registro/login con contraseñas cifradas (bcrypt) y emisión de
  **tokens JWT** firmados (HS256), validados por un *middleware* en cada petición.
- **Catálogo polimórfico:** libros y revistas son `Recurso` (polimorfismo/LSP); se registran,
  buscan, listan y eliminan.
- **Gestión de usuarios (lectores).**
- **Préstamos y devoluciones** con control de disponibilidad del recurso.
- **Persistencia Code First:** los modelos definidos en código generan automáticamente el
  esquema de la base de datos (`AutoMigrate` de GORM).
- **Doble interfaz:** API REST (consumida por Angular) y menú de consola original.

---

## 🚀 Cómo ejecutar

### Requisitos
- [Go 1.23+](https://go.dev/dl/)
- [Node.js 20+](https://nodejs.org/) y npm
- Docker con un contenedor **SQL Server 2022** escuchando en `localhost:1433`

### 1) Base de datos (SQL Server en Docker)

Si aún no existe la base de datos, créala (el esquema/tablas se generan solos vía Code First):

```bash
docker exec sql_server /opt/mssql-tools18/bin/sqlcmd \
  -S localhost -U sa -P "TuPasswordSeguro123!" -No \
  -Q "IF DB_ID('biblioteca') IS NULL CREATE DATABASE biblioteca;"
```

### 2) Backend (API REST)

```bash
cd backend
cp .env.example .env        # ajusta credenciales si es necesario
go run ./cmd/api
```

La API queda en **http://localhost:8081/api**. Al arrancar:
- Conecta a SQL Server y **migra el esquema** (tablas `libros`, `revistas`, `usuarios`,
  `prestamos`, `cuentas`).
- Crea una cuenta administradora por defecto: **`admin` / `admin123`**.

> La CLI original sigue disponible: `go run ./cmd/cli`

### 3) Frontend (Angular)

```bash
cd frontend
npm install        # solo la primera vez
npm start
```

La aplicación queda en **http://localhost:4200**. Inicia sesión con `admin / admin123`.

### Variables de entorno (`backend/.env`)

```env
DB_HOST=localhost
DB_PORT=1433
DB_USER=sa
DB_PASSWORD=TuPasswordSeguro123!
DB_NAME=biblioteca
API_PORT=8081
JWT_SECRET=cambia-esta-clave-secreta-en-produccion
```

---

## 🧩 Integración de las 4 unidades de la asignatura

| Unidad | Concepto | Dónde se evidencia |
|--------|----------|--------------------|
| **1 — Fundamentos de POO** | Clases, encapsulamiento, constructores con validación | `domain/entity/*` (campos privados + getters) |
| **2 — Herencia y polimorfismo** | Interfaz `Recurso` implementada por `Libro` y `Revista` (LSP) | `entity/recurso.go`, préstamo polimórfico |
| **3 — Abstracción e interfaces / SOLID** | Puertos y adaptadores, inyección de dependencias (DIP) | `port/inbound`, `port/outbound`, `application/*` |
| **4 — Persistencia, servicios web y seguridad** | GORM Code First, API REST JSON, JWT | `adapter/outbound/sqlserver`, `adapter/inbound/http` |

### Principios SOLID

| Principio | Aplicación |
|-----------|-----------|
| **SRP** | Un servicio/handler por responsabilidad |
| **OCP** | Nuevos adaptadores (SQL Server, HTTP) sin tocar el dominio |
| **LSP** | `Libro`/`Revista` sustituyen a `Recurso` |
| **ISP** | Puertos pequeños y específicos por caso de uso |
| **DIP** | La aplicación depende de puertos, no de implementaciones |

---

## 🔭 Visualización del futuro

Las tecnologías elegidas (Go + arquitectura hexagonal, API REST con JWT, Code First y Angular)
permiten **escalar** el sistema sin reescribir su núcleo:

- **Microservicios y nube:** al estar el dominio aislado tras puertos, el mismo núcleo podría
  desplegarse como microservicios en contenedores (Docker/Kubernetes).
- **Nuevos canales:** además de la web Angular, podrían añadirse apps móviles o asistentes por
  voz consumiendo los mismos servicios web, sin cambiar la lógica.
- **Biblioteca inteligente:** integración de IA para recomendaciones de lectura, búsqueda
  semántica y predicción de demanda de ejemplares.
- **Trazabilidad y datos:** los préstamos históricos habilitan analítica y reportes para la
  toma de decisiones de la institución.

---

## 🧠 Conclusión reflexiva

Este proyecto integra los conceptos de POO con prácticas profesionales de ingeniería de
software. **Lo aprendido:** que una buena separación de responsabilidades (hexagonal + SOLID)
hace que añadir una API REST, JWT o una base de datos real sea un cambio aditivo y no una
reescritura. **Dificultades:** la rehidratación de entidades de dominio desde la base de datos
manteniendo el encapsulamiento, y la integración de la autenticación JWT entre Angular y Go
(CORS, *interceptors*). **Aplicaciones:** el patrón sirve para cualquier sistema de gestión
empresarial (inventarios, reservas, ventas) donde la lógica de negocio deba mantenerse estable
frente a cambios de infraestructura.

---

## 📄 Documentación

| Archivo | Descripción |
|---------|-------------|
| [docs/PresentacionProyectoFinalPOO.pptx](docs/PresentacionProyectoFinalPOO.pptx) | **Presentación del proyecto** (16 diapositivas: funcionales + técnicas) |
| [docs/ARQUITECTURA.md](docs/ARQUITECTURA.md) | Detalle de la arquitectura hexagonal + SOLID |
| [docs/Autonomo2POOAFVS.pdf](docs/Autonomo2POOAFVS.pdf) | Informe Autónomo 2 |
| [docs/Autonomo1POOAFVS.pdf](docs/Autonomo1POOAFVS.pdf) | Planificación Autónomo 1 |

## Licencia

Proyecto académico — UIDE. Uso educativo.
