# GCR TCP Proxy (Fixed)

A simple TCP proxy for Google Cloud Run with support for custom backend ports.

## Deployment to Google Cloud Run

1. Build and push the image to Docker Hub or Google Container Registry.
2. Deploy to Cloud Run:
   - Set `V2RAY_SERVER_IP` to your server address (e.g., `1.2.3.4:2028`).
   - Port should be `8080`.
   - Ensure the backend port is open on your server.
