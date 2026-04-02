# GCR TCP Proxy (Fixed)

A simple TCP proxy for Google Cloud Run with support for custom backend ports.

## Docker Image
**`docker.io/zuklida/gcr-tcp-proxy:latest`**

## Deployment to Google Cloud Run

1. Build and push the image to Docker Hub or Google Container Registry.
2. Deploy to Cloud Run:
   - Set `V2RAY_SERVER_IP` to your server address (e.g., `<YOUR_SERVER_IP>:<PORT>`).
   - Port should be `8080`.
   - Ensure the backend port is open on your server.
