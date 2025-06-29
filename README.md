# GoCleanApp Generator

GoCleanApp Generator is a minimal Go backend scaffold generator that maintains a go-clean-template-style hierarchy and adds generic Docker deployment files.

## Features

- Interactive CLI for project setup
- Supports PostgreSQL, MySQL, and RabbitMQ
- Generates Dockerfile and docker-compose.yml for containerized deployment
- Includes Makefile for easy local development

## Getting Started

### Prerequisites

- Go 1.22 or later
- Docker
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/uwu-octane/create-goclean-app.git
   cd create-goclean-app
   ```

2. Install the generator:

   ```bash
   go install ./gogogo
   ```

3. Ensure `$GOBIN` is in your PATH:
   ```bash
   export PATH="$HOME/go/bin:$PATH"
   ```

### Usage

Run the generator to create a new project:

```bash
cd /path/to/your/workspace
gogogo
```

Follow the interactive prompts to set up your project.

### Running the Project

#### Locally

```bash
cd your-project-name
make run
```

#### With Docker Compose

```bash
make compose-up
```

## Contributing

Feel free to submit issues or pull requests for improvements and bug fixes.

## License

This project is licensed under the Apache License.
