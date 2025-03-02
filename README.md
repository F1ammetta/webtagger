# Yggdrasil

 Yggdrasil is a web application for managing and editing audio files on a server. Built with Svelte on the frontend and Go on the backend, it allows users to edit metadata of songs and upload new audio files to the server.

## Features

- Upload audio files to server storage ( WIP )
- Edit metadata of existing audio files (artist, title, album, etc.)
- Browse and search through your audio collection

## Development Setup

### Prerequisites

- Node.js (v18+)
- npm or yarn
- Go (v1.23.5+)
- [Tageditor](https://github.com/Martchus/tageditor)
- [FFmpeg](https://github.com/FFmpeg/FFmpeg)

All included in the Nix flake used for developmenta.

### Running the Backend

```bash
# From the project root
go run main.go
```

The Go server will start on port 8080.

### Running the Frontend

```bash
# From the project root
cd frontend
npm install
npm run dev
```

The Vite development server will start on port 3000. During development, API requests to `/api` will be automatically proxied to `localhost:8080/`.

A configuration file will be added in the future.

## Deployment

The application is designed to be deployed behind a reverse proxy (e.g., Nginx). The reverse proxy should:

- Serve the Svelte frontend on the root path (`/`)
- Proxy requests to `/api` to the Go backend

### Default Ports

- Frontend: 3000
- Backend API: 8080

### Sample Nginx Configuration

```nginx
server {
    listen 80;
    server_name yourdomain.com;

    # Serve frontend static files
    location / {
        root /path/to/frontend/dist;
        index index.html;
        try_files $uri $uri/ /index.html;  # For SPA routing
    }

    # Proxy API requests
    location /api {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

### Production Deployment Notes

In production, the Vite server is not used to serve the frontend. Instead:

1. Build the Svelte app with `npm run build` in the frontend directory
2. Serve the static files from the `frontend/dist` directory
3. Configure your reverse proxy to route API requests to the Go server

⚠️ **Note:** The current implementation requires manual configuration for production deployment. Future updates may include a Nix flake for simplified deployment.

## Configuration

### Environment Variables

#### Backend (Go)
- `PORT`: Port for the Go server (default: 8080)
- `STORAGE_PATH`: Path where audio files will be stored (default: "./storage")

#### Frontend (Svelte/Vite)
- `PORT`: Port for the development server (default: 3000)

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

[GPLv3+](LICENSE)
