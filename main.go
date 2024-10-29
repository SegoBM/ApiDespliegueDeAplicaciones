package main

import (
	_ "github.com/denisenkom/go-mssqldb"
)

// Objeto proyecto
type Proyecto struct {
	ID            int    `json:"id"`
	Titulo        string `json:"titulo"`
	Descripcion   string `json:"descripcion"`
	Estudiante    int    `json:"estudiante"`
	FechaRegistro string `json:"fecha_registro"`
	Estatus       string `json:"estatus"`
}

// Objeto de usuario
type Usuario struct {
	ID         int    `json:"id"`
	Usuario    string `json:"usuario"`
	Nombre     string `json:"nombre"`
	Apellidos  string `json:"apellidos"`
	Contrasena string `json:"contrasena"`
	Carrera    string `json:"carrera"`
	Semestre   int    `json:"semestre"`
}
