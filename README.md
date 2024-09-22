# cash_register

Este proyecto fue diseñado para funcionar como una caja registradora, implementando diversas funcionalidades de gestión de transacciones, pagos y control de cambio.

## Descripción

El proyecto sigue la arquitectura limpia (Clean Architecture), separando las responsabilidades en capas independientes para mejorar la mantenibilidad y escalabilidad del sistema. 

### Tecnologías y herramientas utilizadas:

- **GORM**: ORM para la interacción con la base de datos.
- **Gin**: Framework web ligero para manejar las solicitudes HTTP.
- **PostgreSQL**: Base de datos relacional.
- **Docker**: Contenedores para facilitar la implementación del proyecto.

## Arquitectura

- **Patrón Singleton**: Utilizado para gestionar la conexión a la base de datos, asegurando una única instancia activa.
- **Patrón Factory**: Encapsula la lógica de creación de la base de datos, facilitando la extensión y mantenimiento del código.

El proyecto incluye la creación de tres CRUDs completos:

- **Denominations**: Gestión de denominaciones monetarias.
- **MoneyType**: Tipos de dinero aceptados.
- **TransactionType**: Tipos de transacciones (pagos, devoluciones, etc.).

Además, se implementaron funcionalidades completas para la creación de transacciones, pagos y manejo de cambio.

## Instalación

Para clonar e instalar el proyecto localmente, sigue los siguientes pasos:

1. Clona este repositorio:

    ```bash
    git clone git@github.com:JoseAyala97/cash_register.git
    ```

2. Navega al directorio del proyecto:

    ```bash
    cd cash_register
    ```

3. Ejecuta el proyecto usando Docker:

    ```bash
    docker-compose up --build
    ```

   Esto construirá y ejecutará el contenedor del proyecto, incluyendo las migraciones de la base de datos.

### Ejecución manual

Si prefieres ejecutar el proyecto localmente sin Docker, puedes hacerlo usando Go:

1. Asegúrate de tener **Go** instalado en tu sistema.

2. Ejecuta el siguiente comando:

    ```bash
    go run cmd/main.go
    ```

   Esto ejecutará las migraciones automáticamente y conectará el proyecto a la base de datos PostgreSQL.

## Conexión a la base de datos

El proyecto se conecta a una base de datos PostgreSQL utilizando la configuración proporcionada en los archivos de entorno. Si es necesario, ajusta las variables de conexión en el archivo `.env` o `docker-compose.yml`.

## Uso

Una vez ejecutado el proyecto, puedes interactuar con las diferentes rutas y endpoints. Las colecciones de **Postman** con ejemplos de uso están adjuntas al correo electrónico que recibirás. Estas colecciones incluyen las solicitudes para:

- Crear transacciones.
- Consultar denominaciones.
- Registrar pagos.
- Obtener cambios de transacciones, entre otras.