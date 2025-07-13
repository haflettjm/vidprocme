# VidProcMe – Master TODO (clickable in Obsidian)
---

## High-Level Data Flow

```text
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│  User CLI   │───▶│  HTTP API   │───▶│  Scheduler  │───▶│   Worker    │
│   vidctl    │    │  /submit    │    │  + Queue    │    │  (GKE Pod)  │
└─────────────┘    └─────────────┘    └─────────────┘    └─────┬───────┘
                                                               │
                                                               │ (1) fetch
                                                               ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   GCS IN    │◀───┤   FFmpeg    │◀───┤  LLM Batch  │◀───┤ Frame PNGs  │
│  original   │    │  splitter   │    │  (GPT-4o)   │    │  (tmpfs)    │
└─────────────┘    └─────┬───────┘    └─────┬───────┘    └─────────────┘
                         │                  │
                         ▼                  ▼
                ┌─────────────┐    ┌─────────────┐
                │  GCS OUT    │    │  Postgres   │
                │ encoded mp4 │    │ frame_ctx   │
                └─────────────┘    └─────────────┘
```

---

## 0. Meta & Setup
- [x] `mkdir -p vidcompute && cd vidcompute`
- [x] `git init`
- [x] Add `.gitignore` (Go, Node, Terraform, OS)
- [x] Commit initial scaffold
- [x] Push to GitHub (private repo)
- [x] Enable GitHub Actions
- [x] Install prerequisites (Docker, kubectl, Terraform, Go 1.22+, Node 20+)

---

## 1. Root Directory Skeleton
- [x] `mkdir -p cmd/ internal/ web/ infra/ scripts/ .github/workflows`
- [x] `touch README.md LICENSE Makefile Dockerfile docker-compose.yml`

---

## 2. Core Entry Points (`cmd/`)
- [ ] `cmd/main.go` – bootstrap HTTP server, graceful shutdown
- [ ] Wire config → logger → router → queue → scheduler
- [ ] `cmd/vidctl/main.go` – Cobra root
- [ ] `cmd/vidctl/submit.go` – POST job JSON
- [ ] `cmd/vidctl/status.go` – GET job status
- [ ] `cmd/vidctl/watch.go` – SSE logs
- [ ] Embed version via `ldflags`

---

## 3. Configuration (`internal/config/`)
- [ ] `config.go` – struct with env tags
- [ ] Load via Viper (.yaml + env override)
- [ ] Validation (required, regex, numeric ranges)
- [ ] Unit tests with table-driven cases

---

## 4. HTTP API (`internal/api/`)
- [ ] Router (`chi`) + middleware (request-id, CORS, recovery, metrics)
- [ ] `POST /api/v1/jobs` – accept JobSpec JSON
- [ ] `GET /api/v1/jobs/{id}` – job status & progress
- [ ] `GET /api/v1/jobs/{id}/logs` – SSE log tail
- [ ] `GET /api/v1/jobs/{id}/frames` – paginated frame context
- [ ] `GET /api/v1/search?q=...` – full-text on frame JSON
- [ ] Serve static `/dashboard` (from `web/`)
- [ ] OpenAPI spec + Swagger UI

---

## 5. Queue (`internal/queue/`)
- [ ] Interface `Queue` with enqueue, dequeue, ack, nack
- [ ] In-memory impl (slice + mutex) for local dev
- [ ] Redis impl (`go-redis/v9`) with TTL, retry, dead-letter
- [ ] Exponential back-off policy config
- [ ] Unit tests with `miniredis`

---

## 6. Scheduler (`internal/scheduler/`)
- [ ] GPU node discovery interface
- [ ] FIFO + priority queue (high/normal/low)
- [ ] Bin-packing (first-fit decreasing)
- [ ] Metrics: `scheduler_queue_depth`, `scheduler_jobs_scheduled_total`
- [ ] Integration test on KinD

---

## 7. Transcoder (`internal/transcoder/`)
- [ ] `JobSpec` struct (input, outputs, codec, bitrate, enableLLM bool)
- [ ] `Worker.Run()` – download → FFmpeg → upload
- [ ] FFmpeg progress parser (`frame=`, `fps=`, `time=`)
- [ ] Side-car JSON manifest for multi-bitrate HLS
- [ ] Pre-signed URL generation

---

## 8. LLM Frame-Context (NEW)
- [ ] `internal/llm/extractor.go` – FFmpeg scene-detect or fixed fps
- [ ] `internal/llm/client.go` – interface + OpenAI impl
- [ ] Rate-limiting & retry logic
- [ ] `internal/llm/store.go` – batch insert to Postgres
- [ ] Migration `migrations/0001_create_frame_context.sql`
- [ ] Add secrets for LLM key (GCP Secret Manager)
- [ ] Update infra node pool CPU/RAM for LLM

---

## 9. Monitor (`internal/monitor/`)
- [ ] Prometheus registry singleton
- [ ] Counters: `jobs_total`, `jobs_failed_total`
- [ ] Histograms: `job_duration_seconds`, `llm_inference_seconds`
- [ ] Structured logs (Zap)
- [ ] Grafana dashboard JSON (infra auto-import)

---

## 10. Web Dashboard (`web/`)
- [ ] `dashboard.html` – Bulma CSS skeleton
- [ ] `js/jobs.js` – auto-refresh table
- [ ] `js/logs.js` – SSE modal for live logs
- [ ] `js/frames.js` – timeline scrubber with frame-by-frame context
- [ ] Dark-mode toggle
- [ ] Bundle with esbuild

---

## 11. Infrastructure (`infra/`)
- [ ] `cdktf.json` – TypeScript CDKTF init
- [ ] `infra/gcp/main.ts` – GKE Autopilot + CloudSQL + GCS + Secret Manager
- [ ] `infra/local/main.ts` – KinD + Redis + Prometheus + Grafana
- [ ] Shared constructs (`VidComputeService`, `VidComputeQueue`)
- [ ] `make infra/local/up` one-command dev stack

---

## 12. Docker & Compose
- [ ] Multi-stage Dockerfile (distroless)
- [ ] `.dockerignore`
- [ ] `docker-compose.yml` (local dev)
- [ ] `compose.dev.yml` override (volume mounts, hot-reload)

---

## 13. CI/CD (`.github/workflows/`)
- [ ] `build.yml` – Docker build & push
- [ ] `test.yml` – `go test ./...`
- [ ] `infra-plan.yml` – Terraform plan on PR
- [ ] `infra-apply.yml` – manual apply
- [ ] `e2e-test.yml` – KinD + submit job + assert completion
- [ ] `release.yml` – GoReleaser for `vidctl` binaries

---

## 14. Scripts (`scripts/`)
- [ ] `dev-up.sh` – start KinD + infra/local
- [ ] `dev-down.sh` – tear down
- [ ] `load-sample-video.sh` – upload test clip
- [ ] `generate-ffmpeg-cmd.sh` – helper script

---

## 15. Security & Compliance
- [ ] Container scan (Trivy) in CI
- [ ] SLSA provenance via GitHub OIDC
- [ ] NetworkPolicies (deny-all, allow-apiserver)
- [ ] Workload Identity for GKE → GCS/CloudSQL

---

## 16. Docs & Runbooks
- [ ] `docs/architecture.png` – C4 diagram
- [ ] `docs/openapi.yaml` – OpenAPI 3.1
- [ ] `docs/runbook.md` – on-call playbooks
- [ ] `docs/costs.md` – GCP cost calculator
- [ ] `docs/build-log.md` – Obsidian daily log

---

## 17. Final Validation
- [ ] Local e2e: upload 5 s clip → encoded + 150 frame rows
- [ ] Cloud e2e: same on GKE Autopilot
- [ ] Chaos: kill worker, assert retry & DLQ
- [ ] Load: 100 concurrent jobs under 5 min

---


---
---


# AUTHOR NOTES


This is probably good enough for what I want to do for now
