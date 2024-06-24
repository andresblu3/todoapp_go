# Lista de Tareas (Todo List)

# Lista de Tareas (Todo List)

Esta es una aplicación de lista de tareas desarrollada con **Vue.js**, **Vuetify** y **Go** (Golang), utilizando **MySQL** como base de datos.

## Requisitos

- **Node.js** y **npm**
- **Go**
- **MySQL**

## Instalación

### Clonar el repositorio

```bash
git clone https://github.com/andresblu3/todoapp_go.git
```

### Instala las dependencias de Go:

```bash
go mod tidy
```


Actualiza la configuración de la base de datos en config/config.go con tus credenciales de MySQL.

Ejecuta el servidor backend:

```bash
go run main.go
```

## Configuración del Frontend

Instala las dependencias de npm:

```bash
cd todo-app-frontend
npm install
```

Ejecuta la aplicación frontend:

```bash
npm run dev
```

Abre tu navegador y navega a http://localhost:8080 para ver la aplicación en funcionamiento.

