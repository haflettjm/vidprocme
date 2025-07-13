Below is the **single-file, most-detailed, actionable** TODO list you can copy-paste into `TODO.md` and check-off as you work.
Every bullet is a **git-commit-sized** task and references the exact file / directory it lives in.

---

# VidCompute v2 — Ultra-Granular TODO

> Legend
> 🔲 = not started 🚧 = in-progress ✅ = done ⏭️ = blocked / next sprint
> Use `git commit -m "feat(internal): 1.3.1 – add Prometheus histogram for job_duration_seconds"`

---

## 1. Root & Meta Files
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 1.1 | Scaffold root directories (`mkdir -p …`) | `vidcompute/*` | All | <input type="checkbox"> |
| 1.2 | Add `.gitignore` (Go, Node, Terraform, OS) | `.gitignore` | DevOps | <input type="checkbox"> |
| 1.3 | Write top-level README badges + quick-start | `README.md` | Docs | <input type="checkbox"> |
| 1.4 | Create `LICENSE` (Apache-2.0) | `LICENSE` | Legal | <input type="checkbox"> |
| 1.5 | Add `Makefile` with `make dev`, `make test`, `make infra/local/up` | `Makefile` | DevOps | <input type="checkbox"> |

---

## 2. cmd/
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 2.1 | `main.go` skeleton (wire `config.Load`, `log.Init`, graceful shutdown) | `cmd/main.go` | Backend | <input type="checkbox"> |
| 2.2 | HTTP server bootstrap (port 8080) | `cmd/main.go` | Backend | <input type="checkbox"> |
| 2.3 | Mount API routes (`/health`, `/metrics`, `/api/v1/*`) | `cmd/main.go` | Backend | <input type="checkbox"> |
| 2.4 | CLI module `vidctl` cobra scaffolding | `cmd/vidctl/main.go` | Backend | <input type="checkbox"> |
| 2.5 | `vidctl submit` – read JSON job file, POST to API | `cmd/vidctl/submit.go` | Backend | <input type="checkbox"> |
| 2.6 | `vidctl status <id>` – pretty-print JSON | `cmd/vidctl/status.go` | Backend | <input type="checkbox"> |
| 2.7 | `vidctl watch <id>` – SSE stream logs | `cmd/vidctl/watch.go` | Backend | <input type="checkbox"> |
| 2.8 | `vidctl version` – embed build info via `ldflags` | `cmd/vidctl/version.go` | Backend | <input type="checkbox"> |
| 2.9 | `vidctl help` – print usage | `cmd/vidctl/help.go` | Backend | <input type="checkbox"> |
| 2.10 | `vidctl logs <id>` – fetch logs from API | `cmd/vidctl/logs.go` | Backend | <input type="checkbox"> |

---

## 3. internal/config
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 3.1 | Define `Config` struct (env + JSON) | `internal/config/config.go` | Backend | <input type="checkbox"> |
| 3.2 | Viper loader with `config.yaml` fallback | `internal/config/loader.go` | Backend | <input type="checkbox"> |
| 3.3 | Validation (required fields, regex) | `internal/config/validate.go` | Backend | <input type="checkbox"> |
| 3.4 | Unit tests w/ table-driven tests | `internal/config/*_test.go` | Backend | <input type="checkbox"> |

---

## 4. internal/api
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 4.1 | Router (`chi`) with middleware: CORS, request-id, recovery | `internal/api/router.go` | Backend | <input type="checkbox"> |
| 4.2 | `POST /api/v1/jobs` – validate JSON schema | `internal/api/handlers.go` | Backend | <input type="checkbox"> |
| 4.3 | `GET /api/v1/jobs/{id}` – return job status | `internal/api/handlers.go` | Backend | <input type="checkbox"> |
| 4.4 | `GET /api/v1/jobs/{id}/logs` – tail logs (SSE) | `internal/api/handlers.go` | Backend | <input type="checkbox"> |
| 4.5 | `GET /dashboard` – serve `web/dashboard.html` | `internal/api/handlers.go` | Backend | <input type="checkbox"> |
| 4.6 | OpenAPI 3.1 spec generation (`/docs/swagger.json`) | `docs/openapi.yaml` + `internal/api/docs.go` | Backend | <input type="checkbox"> |
| 4.7 | Add HTTP metrics middleware (Prometheus) | `internal/api/middleware.go` | Backend | <input type="checkbox"> |

---

## 5. internal/queue
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 5.1 | Interface `Queue` (Enqueue, Dequeue, Ack, Nack) | `internal/queue/queue.go` | Backend | <input type="checkbox"> |
| 5.2 | In-memory implementation (slice + mutex) | `internal/queue/memory.go` | Backend | <input type="checkbox"> |
| 5.3 | Redis implementation (`go-redis`) | `internal/queue/redis.go` | Backend | <input type="checkbox"> |
| 5.4 | TTL & dead-letter logic | `internal/queue/ttl.go` | Backend | <input type="checkbox"> |
| 5.5 | Retry with exponential backoff | `internal/queue/retry.go` | Backend | <input type="checkbox"> |
| 5.6 | Unit tests with `miniredis` | `internal/queue/*_test.go` | Backend | <input type="checkbox"> |

---

## 6. internal/scheduler
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 6.1 | GPU node discovery interface | `internal/scheduler/cluster.go` | Backend | <input type="checkbox"> |
| 6.2 | Mock GPU inventory (for local) | `internal/scheduler/mock.go` | Backend | <input type="checkbox"> |
| 6.3 | Bin-packing algo (first-fit decreasing) | `internal/scheduler/scheduler.go` | Backend | <input type="checkbox"> |
| 6.4 | Priority queue (high, normal, low) | `internal/scheduler/scheduler.go` | Backend | <input type="checkbox"> |
| 6.5 | Metrics: `scheduler_jobs_scheduled_total`, `scheduler_queue_depth` | `internal/scheduler/metrics.go` | Backend | <input type="checkbox"> |
| 6.6 | Integration tests (KinD) | `tests/scheduler_test.go` | QA | <input type="checkbox"> |

---

## 7. internal/transcoder
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 7.1 | `JobSpec` struct (input, outputs, codec, bitrate) | `internal/transcoder/types.go` | Backend | <input type="checkbox"> |
| 7.2 | `Worker.Run()` – download from GCS, run FFmpeg | `internal/transcoder/worker.go` | Backend | <input type="checkbox"> |
| 7.3 | FFmpeg progress parser (regex frame=…) | `internal/transcoder/progress.go` | Backend | <input type="checkbox"> |
| 7.4 | Upload outputs back to GCS | `internal/transcoder/uploader.go` | Backend | <input type="checkbox"> |
| 7.5 | Pre-signed URL generation for secure downloads | `internal/transcoder/urls.go` | Backend | <input type="checkbox"> |
| 7.6 | Unit tests with sample 1-second MP4 | `internal/transcoder/*_test.go` | Backend | <input type="checkbox"> |

---

## 8. internal/monitor
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 8.1 | Prometheus registry singleton | `internal/monitor/metrics.go` | Backend | 🔲 |
| 8.2 | Counters: jobs_total, jobs_failed_total | `internal/monitor/metrics.go` | Backend | 🔲 |
| 8.3 | Histograms: job_duration_seconds | `internal/monitor/metrics.go` | Backend | 🔲 |
| 8.4 | Zap logger setup (JSON, level, output) | `internal/monitor/logger.go` | Backend | 🔲 |
| 8.5 | Request-ID propagation middleware | `internal/monitor/middleware.go` | Backend | 🔲 |
| 8.6 | Grafana dashboard JSON (import via infra) | `grafana/dashboard.json` | DevOps | 🔲 |

---

## 9. web/
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 9.1 | HTML skeleton (Bulma CSS) | `web/dashboard.html` | Frontend | 🔲 |
| 9.2 | JS: fetch `/api/v1/jobs` + render table | `web/js/jobs.js` | Frontend | 🔲 |
| 9.3 | JS: auto-refresh every 5 s | `web/js/jobs.js` | Frontend | 🔲 |
| 9.4 | JS: SSE for live logs modal | `web/js/logs.js` | Frontend | 🔲 |
| 9.5 | Dark mode toggle | `web/js/theme.js` | Frontend | 🔲 |
| 9.6 | Bundle w/ esbuild (no React needed) | `web/package.json` | Frontend | 🔲 |

---

## 10. infra/
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 10.1 | CDKTF init (TypeScript, GCP provider) | `infra/cdktf.json` | DevOps | 🔲 |
| 10.2 | `gcp/main.ts` – GKE Autopilot cluster | `infra/gcp/main.ts` | DevOps | 🔲 |
| 10.3 | `gcp/main.ts` – CloudSQL Postgres (job metadata) | `infra/gcp/main.ts` | DevOps | 🔲 |
| 10.4 | `gcp/main.ts` – Cloud Storage buckets (input/output) | `infra/gcp/main.ts` | DevOps | 🔲 |
| 10.5 | `gcp/main.ts` – Workload Identity bindings | `infra/gcp/main.ts` | DevOps | 🔲 |
| 10.6 | `local/main.ts` – KinD cluster + registry | `infra/local/main.ts` | DevOps | 🔲 |
| 10.7 | `local/main.ts` – Redis Helm release | `infra/local/main.ts` | DevOps | 🔲 |
| 10.8 | `local/main.ts` – Prometheus + Grafana stack | `infra/local/main.ts` | DevOps | 🔲 |
| 10.9 | Shared constructs (`VidComputeService`) | `infra/lib/constructs.ts` | DevOps | 🔲 |
| 10.10 | Output kubeconfig for local | `infra/local/outputs.ts` | DevOps | 🔲 |

---

## 11. .github/workflows
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 11.1 | Build & push Docker image | `.github/workflows/build.yml` | DevOps | 🔲 |
| 11.2 | Run unit tests (`go test ./...`) | `.github/workflows/test.yml` | DevOps | 🔲 |
| 11.3 | Terraform plan on PR (infra-plan.yml) | `.github/workflows/infra-plan.yml` | DevOps | 🔲 |
| 11.4 | Manual approve + infra apply (infra-apply.yml) | `.github/workflows/infra-apply.yml` | DevOps | 🔲 |
| 11.5 | KinD e2e test (submit job, assert done) | `.github/workflows/e2e-test.yml` | QA | 🔲 |
| 11.6 | Release binaries (vidctl) via GoReleaser | `.github/workflows/release.yml` | DevOps | 🔲 |

---

## 12. Docker & Compose
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 12.1 | Multi-stage Dockerfile (distroless) | `Dockerfile` | DevOps | 🔲 |
| 12.2 | `.dockerignore` | `.dockerignore` | DevOps | 🔲 |
| 12.3 | `docker-compose.yml` (API, Redis, Grafana) | `docker-compose.yml` | DevOps | 🔲 |
| 12.4 | Override file `compose.dev.yml` (volume mounts) | `compose.dev.yml` | DevOps | 🔲 |

---

## 13. scripts/
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 13.1 | `dev-up.sh` – starts KinD + infra/local | `scripts/dev-up.sh` | DevOps | 🔲 |
| 13.2 | `dev-down.sh` – tears down | `scripts/dev-down.sh` | DevOps | 🔲 |
| 13.3 | `generate-ffmpeg-cmd.sh` – helper for CLI | `scripts/ffmpeg.sh` | Backend | 🔲 |
| 13.4 | `load-sample-video.sh` – curl upload demo | `scripts/sample.sh` | QA | 🔲 |

---

## 14. Testing Matrix
| # | Task | Scope | Owner | Status |
|---|------|-------|-------|--------|
| 14.1 | Unit tests (Go) | `go test ./...` | Backend | 🔲 |
| 14.2 | Static analysis (golangci-lint) | CI | Backend | 🔲 |
| 14.3 | Contract tests (OpenAPI spec) | CI | QA | 🔲 |
| 14.4 | Chaos test – kill worker pod | KinD | QA | 🔲 |
| 14.5 | Load test – 100 concurrent jobs | GKE | QA | 🔲 |

---

## 15. Documentation & Ops
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 15.1 | Architecture diagram (C4) | `docs/architecture.png` | Docs | 🔲 |
| 15.2 | API docs (Swagger UI) | `docs/index.html` | Docs | 🔲 |
| 15.3 | Runbook (alerts, SLOs) | `docs/runbook.md` | SRE | 🔲 |
| 15.4 | Cost calculator sheet | `docs/costs.md` | FinOps | 🔲 |

---

## 16. Security & Compliance
| # | Task | File(s) | Owner | Status |
|---|------|---------|-------|--------|
| 16.1 | Container scan (Trivy) in CI | `.github/workflows/security.yml` | Sec | 🔲 |
| 16.2 | SLSA provenance (GitHub OIDC) | `.github/workflows/release.yml` | Sec | 🔲 |
| 16.3 | NetworkPolicies (deny-all, allow-apiserver) | `infra/gcp/network-policies.yaml` | Sec | 🔲 |
| 16.4 | Secret management (GCP Secret Manager) | `infra/gcp/secrets.ts` | Sec | 🔲 |

---
