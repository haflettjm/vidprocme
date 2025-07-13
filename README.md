# vidprocme

**Distributed Video Processing Pipeline in Go**
Accelerated video analysis and LLM-powered batch processing, designed for real-world scale.

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](LICENSE)
[![Build](https://github.com/haflettjm/vidprocme/actions/workflows/build.yml/badge.svg)](https://github.com/haflettjm/vidprocme/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/haflettjm/vidprocme.svg)](https://pkg.go.dev/github.com/haflettjm/vidprocme)

---

## Project Purpose

While interviewing with Sieve Data, I wanted to move beyond a simple demo and build a complete, production-ready video pipeline. This project shows how to run a distributed, cloud-ready video system that is practical and scalable—not just an academic toy. The focus is on:

* Reliable transcoding and frame extraction
* Priority-based job scheduling, including GPU-awareness, backoff, and retries
* End-to-end observability with Prometheus and Grafana
* Automated, reproducible deployments using CDKTF and Docker
* Full CI/CD and compliance pipeline

---

## System Overview

* **Video Ingestion:** Upload video files via HTTP API (either JSON jobspec or file upload)
* **Transcoding and Frame Extraction:** FFmpeg, running in parallel with Go workers
* **LLM Batch Analysis:** Frames are processed through an LLM (e.g., GPT-4o), and the results are stored for later use
* **Queue and Scheduling:** Robust, retryable task runner supporting priorities and efficient resource allocation
* **Observability:** Metrics and dashboards track every stage of processing
* **Hybrid Deployments:** The stack works both locally (KinD) and in Google Cloud (Autopilot, CloudSQL, GCS)
* **Developer Experience:** Scripts, a local development stack, CI/CD pipelines, and clean infrastructure-as-code patterns make it easy to get started or extend

### High-Level Data Flow

```text
┌──────────────┐    ┌───────────┐    ┌────────────┐    ┌─────────────┐
│  Frontend    │───▶│ HTTP API  │───▶│ Scheduler  │───▶│   Worker    │
└──────────────┘    └───────────┘    └────────────┘    └─────┬───────┘
                                                            │
                                                            ▼
┌─────────────┐◀───┐
│ GCS (in)    │    │
└─────────────┘    │
 ┌───────┐   ┌───────────┐   ┌─────────────┐   ┌─────────────┐
 │ ffmpeg│◀──│ LLM Batch │◀──│ Frame PNGs  │◀──│   Worker    │
 └───────┘   └───────────┘   └─────────────┘   └─────────────┘
         │             │
         ▼             ▼
┌─────────────┐   ┌─────────────┐
│ GCS (out)   │   │ Postgres    │
│ Encoded mp4 │   │ Frame Ctx   │
└─────────────┘   └─────────────┘
```

---

## Project Structure

```text
vidprocme/
├── cmd/               # Main entry points (server, CLI planned)
├── internal/          # Core packages: queue, scheduler, transcoder, monitor, config, API handlers
├── web/               # Static dashboard UI (Bulma/JavaScript)
├── infra/             # CDKTF app (TypeScript) for GCP and local
│   ├── gcp/           # Google Cloud stack (GKE, CloudSQL, GCS)
│   └── local/         # KinD/k3d stack for development, Redis, Prometheus
├── .github/workflows/ # GitHub Actions for CI/CD and infrastructure
├── scripts/           # Local dev tools: infra up/down, load test, ffmpeg helpers
├── Dockerfile         # Multi-stage build for containers
├── docker-compose.yml # Local orchestration
└── README.md
```

---

## Key Features

* **Fast Video Transcoding**
  Parallelized Go workers use ffmpeg for efficient encoding.

* **Intelligent GPU Scheduling**
  The job queue supports FIFO, priorities, bin-packing, and backoff, scaling to real workloads.

* **LLM Frame Analysis**
  The LLM client is flexible, supporting OpenAI now and designed for future support of local models.

* **Reliability and Observability**
  Metrics are exported for Prometheus and Grafana. Logging uses structured output with Zap. Health checks, status endpoints, retries, and a dead letter queue improve resilience.

* **Infrastructure Automation**
  CDKTF (TypeScript) automates deployments for both GCP and KinD local.
  CI/CD pipelines handle build, test, infrastructure plan/apply, and end-to-end tests.
  Container security and compliance are built in (Trivy, SLSA, OIDC).

* **Web Dashboard**
  Live job tracking, frame-by-frame LLM output, and dark mode are available out of the box.

* **Postgres Job Store**
  All metadata and LLM context are stored in a normalized, queryable schema.

---

## Getting Started

### Prerequisites

* Go 1.22+
* Node.js 20+
* Docker and Docker Compose
* kubectl
* Terraform/CDKTF

### Quick Start (Local)

```bash
git clone https://github.com/haflettjm/vidprocme.git
cd vidprocme
make infra/local/up
make run-dev    # or docker-compose up --build
# The API is available at http://localhost:8080 and the dashboard at /dashboard
```

### Deployment on Google Cloud

```bash
cd infra/gcp
cdktf deploy
# Configure secrets and update node pools as needed. See infra/main.ts for more details.
```

---

## Roadmap

**Legend:**
\[x] Complete
\[ ] Planned
~~strikethrough~~: Post-MVP or stretch goal

### Core System

* [x] Directory structure, GitHub repo, CI pipeline
* [ ] Connect config, logger, router, queue, and scheduler
* [ ] Unified configuration loader (env + file, using viper)
* [ ] HTTP API: job submission, status, logs, frame context, and full-text search
* [ ] Serve dashboard (static, Bulma/JS)
* [ ] OpenAPI and Swagger UI

### Queue and Scheduling

* [ ] Pluggable queue: in-memory for dev, Redis for production, with exponential backoff and DLQ
* [ ] Scheduler: GPU-aware, priority bins, bin-packing, and job metrics

### Transcoding and LLM

* [ ] Jobspec support (inputs, outputs, codecs, enable LLM)
* [ ] FFmpeg progress tracking and presigned URLs
* [ ] LLM frame context extraction, batching, rate limiting, retries, and Postgres storage

### Monitoring

* [ ] Prometheus metrics, Grafana dashboards, structured logging
* [ ] End-to-end and integration test pipelines

### Infrastructure

* [ ] CDKTF for GCP (GKE, CloudSQL, GCS, Secret Manager) and KinD local
* [ ] One-command local stack; scripts for development and load testing

### Security

* [ ] Container scanning, SLSA provenance, network policies, and GCP workload identity

### Documentation

* [ ] Architecture diagrams (C4 model)
* [ ] OpenAPI spec
* [ ] Runbooks (on-call, build, and cost analysis)

### Final Validation

* [ ] End-to-end tests, both locally and in the cloud (process 5s video: 150+ frames)
* [ ] Chaos and DLQ validation; load test: 100 concurrent jobs in under five minutes

---

## Stretch Goals

* CLI tool (`vidctl`): for job submission, status, and live watch (Cobra, GoReleaser)
* GPU node simulation and multiple worker orchestration

---

## Design Choices and Tradeoffs

* **Go for the core:** Simple concurrency, robust ecosystem, and easy deployments.
  The main limitation is the lack of a native HA/actor model like Erlang or Elixir, which is mitigated by relying on external queues and aggressive monitoring.

* **Redis as a queue:** Fast and reliable, but persistence tuning is needed for production use.

* **Hybrid GCP and KinD deployment:** High local/prod parity, but some vendor lock-in for managed secrets and CloudSQL.

* **LLM integration:** OpenAI is used initially for convenience, but the system is designed to support pluggable backends, such as local LLMs (llama.cpp, ollama) in the future.

---

## Contributing

Contributions are always welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines (coming soon). All code is released under the GPLv3 license.

---

## License

This project is licensed under the GNU GPLv3. See [LICENSE](LICENSE) for details.

---

## Author

Contact: [github@haflettjm](mailto:github@haflettjm)
Project started: July 8, 2025

---

<details>
  <summary>Legacy / Author Notes</summary>
  <ul>
    <li>This project was built to gain real-world experience with video and LLM processing, and as a robust portfolio reference.</li>
    <li>Feedback, pull requests, and architecture discussions are welcome—especially from infrastructure and ML Ops engineers.</li>
  </ul>
</details>

---

## Additional Resources

* [TODO.md](TODO.md): Full task breakdown and roadmap
* [docs/](docs/): (WIP) Architecture diagrams, OpenAPI spec, runbooks

---

Let me know if you want it even more conversational, or if you have any other tone/style requests.
