// Punto de entrada de la API REST: ensambla los adaptadores (SQL Server + HTTP/JWT)
// con los casos de uso de la aplicación (arquitectura hexagonal + SOLID).
//
//	@title			API Sistema de Gestión de Libros — Biblioteca UIDE
//	@version		1.0
//	@description	API REST del proyecto final de POO. Catálogo de libros y revistas,
//	@description	usuarios (lectores) y préstamos, con autenticación JWT y persistencia
//	@description	Code First sobre SQL Server. La serialización es JSON.
//	@contact.name	Valenzuela Saavedra Andrés Fernando — UIDE
//	@host			localhost:8081
//	@BasePath		/api
//	@schemes		http
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Autenticación JWT. Escribe: "Bearer {token}" (obtén el token en POST /auth/login).
package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "sistema-gestion-libros/docs" // spec OpenAPI generada por swag (Swagger UI)
	apphttp "sistema-gestion-libros/internal/adapter/inbound/http"
	"sistema-gestion-libros/internal/adapter/outbound/persistence/sqlserver"
	"sistema-gestion-libros/internal/adapter/outbound/seguridad"
	"sistema-gestion-libros/internal/application"
	"sistema-gestion-libros/internal/domain/entity"
	"sistema-gestion-libros/internal/domain/errores"
)

func main() {
	_ = godotenv.Load() // carga variables desde .env si existe (opcional)

	// ---- Adaptador outbound: persistencia SQL Server (Code First) ----
	cfg := sqlserver.Config{
		Host:     env("DB_HOST", "localhost"),
		Puerto:   env("DB_PORT", "1433"),
		Usuario:  env("DB_USER", "sa"),
		Password: env("DB_PASSWORD", "TuPasswordSeguro123!"),
		Base:     env("DB_NAME", "biblioteca"),
	}
	db, err := sqlserver.Conectar(cfg)
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a la base de datos: %v", err)
	}
	log.Printf("✅ Conectado a SQL Server (%s/%s) y esquema migrado (Code First).", cfg.Host, cfg.Base)

	repoCatalogo := sqlserver.NuevoRepoCatalogo(db)
	repoUsuarios := sqlserver.NuevoRepoUsuarios(db)
	repoPrestamos := sqlserver.NuevoRepoPrestamos(db)
	repoCuentas := sqlserver.NuevoRepoCuentas(db)
	hasher := seguridad.NuevoBcryptHasher()

	// ---- Capa de aplicación: casos de uso ----
	catalogoSvc := application.NuevoCatalogoServicio(repoCatalogo)
	usuarioSvc := application.NuevoUsuarioServicio(repoUsuarios)
	prestamoSvc := application.NuevoPrestamoServicio(repoCatalogo, repoUsuarios, repoPrestamos)
	consultaSvc := application.NuevoConsultaServicio("Biblioteca UIDE", repoCatalogo, repoPrestamos, catalogoSvc)
	gestor := application.NuevoGestorBiblioteca(catalogoSvc, usuarioSvc, prestamoSvc, consultaSvc)
	authSvc := application.NuevoAuthServicio(repoCuentas, hasher)

	sembrarAdmin(authSvc)

	// ---- Adaptador inbound: API HTTP con JWT ----
	jwt := apphttp.NuevoGestorJWT(env("JWT_SECRET", "cambia-esta-clave-secreta-en-produccion"), 24)
	servidor := apphttp.NuevoServidor(gestor, authSvc, jwt)

	puerto := env("API_PORT", "8081")
	log.Printf("🚀 API escuchando en http://localhost:%s/api", puerto)
	if err := servidor.Router().Run(":" + puerto); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}

// sembrarAdmin crea una cuenta administradora por defecto si aún no existe,
// para poder iniciar sesión la primera vez (admin / admin123).
func sembrarAdmin(auth *application.AuthServicio) {
	_, err := auth.Registrar("admin", "admin123", entity.RolAdmin)
	switch {
	case err == nil:
		log.Println("👤 Cuenta admin creada por defecto (usuario: admin / contraseña: admin123).")
	case errors.Is(err, errores.ErrCuentaYaExiste):
		// Ya existe, nada que hacer.
	default:
		log.Printf("⚠️  No se pudo sembrar la cuenta admin: %v", err)
	}
}

func env(clave, porDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}
