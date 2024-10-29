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

/ Get para proyectos del usuario
    r.GET("/proyectos/:id_usuario", func(c *gin.Context) {
        idUsuario, err := strconv.Atoi(c.Param("id_usuario"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
            return
        }
        proyectos, err := getProyectosByUsuario(db, idUsuario)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, proyectos)
    })

    // Get para proyectos
    r.GET("/proyectos", func(c *gin.Context) {
        proyectos, err := getProyectos(db)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, proyectos)
    })

    // Post para proyectos
    r.POST("/proyectos", func(c *gin.Context) {
        var nuevoProyecto Proyecto
        if err := c.ShouldBindJSON(&nuevoProyecto); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := createProyecto(db, &nuevoProyecto); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, nuevoProyecto)
    })

    // Put para proyectos
    r.PUT("/proyectos/:id", func(c *gin.Context) {
        var proyecto Proyecto
        if err := c.ShouldBindJSON(&proyecto); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
            return
        }
        proyecto.ID = id

        if err := updateProyecto(db, &proyecto); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, proyecto)
    })

    // Delete para proyectos
    r.DELETE("/proyectos/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
            return
        }

        if err := deleteProyecto(db, id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Proyecto eliminado"})
    })

// Post para usuarios

r.POST("/usuarios", func(c *gin.Context) {
    var nuevoUsuario Usuario
    if err := c.ShouldBindJSON(&nuevoUsuario); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := createUsuario(db, &nuevoUsuario); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, nuevoUsuario)
})

// Post para autenticar usuarios
r.POST("/auth", func(c *gin.Context) {
    var credenciales struct {
        Usuario    string json:"usuario"
        Contrasena string json:"contrasena"
    }
    if err := c.ShouldBindJSON(&credenciales); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    usuario, err := authenticateUsuario(db, credenciales.Usuario, credenciales.Contrasena)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
        return
    }

    c.JSON(http.StatusOK, usuario)
})

r.Run(":8080") 
}
