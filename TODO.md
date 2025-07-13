# vidprocme – master todo (clickable in obsidian)
---

## high-level data flow

```text
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│  simple Frontend   │───▶│  http api   │───▶│  scheduler  │───▶│   worker    │
│             │    │  /submit    │    │  + queue    │    │  (gke pod)  │
└─────────────┘    └─────────────┘    └─────────────┘    └─────┬───────┘
                                                               │
                                                               │ (1) fetch
                                                               ▼
┌─────────────┐    ┌─────────────┐    ┌─────────────┐    ┌─────────────┐
│   gcs in    │◀───┤   ffmpeg    │◀───┤  llm batch  │◀───┤ frame pngs  │
│  original   │    │  splitter   │    │  (gpt-4o)   │    │  (tmpfs)    │
└─────────────┘    └─────┬───────┘    └─────┬───────┘    └─────────────┘
                         │                  │
                         ▼                  ▼
                ┌─────────────┐    ┌─────────────┐
                │  gcs out    │    │  postgres   │
                │ encoded mp4 │    │ frame_ctx   │
                └─────────────┘    └─────────────┘
```

---

## 0. meta & setup
- [x] `mkdir -p vidcompute && cd vidcompute`
- [x] `git init`
- [x] add `.gitignore` (go, node, terraform, os)
- [x] commit initial scaffold
- [x] push to github (private repo)
- [x] enable github actions
- [x] install prerequisites (docker, kubectl, terraform, go 1.22+, node 20+)

---

## 1. root directory skeleton
- [x] `mkdir -p cmd/ internal/ web/ infra/ scripts/ .github/workflows`
- [x] `touch readme.md license makefile dockerfile docker-compose.yml`

---

## 2. core entry points (`cmd/`)
- [x] `cmd/server/main.go` – bootstrap http server, graceful shutdown
- [ ] wire config → logger → router → queue → scheduler
- [ ] ~~`cmd/vidctl/main.go` – cobra root~~ ➜ **stretch goal / post-MVP CLI tool**
- [ ] ~~`cmd/vidctl/submit.go` – post job json~~
- [ ] ~~`cmd/vidctl/status.go` – get job status~~
- [ ] ~~`cmd/vidctl/watch.go` – sse logs~~
- [ ] ~~embed version via `ldflags`~~

---

## 3. configuration (`internal/config/`)
- [x] `config.go` – struct with env tags
- [ ] load via viper (.yaml + env override)
- [ ] validation (required, regex, numeric ranges)
- [ ] unit tests with table-driven cases

---

## 4. http api (`internal/api/`)
- [ ] router (`gin`) + middleware (request-id, cors, recovery, metrics)
- [ ] `post /api/v1/jobs` – accept jobspec json OR multipart file upload
- [ ] `get /api/v1/jobs/{id}` – job status & progress
- [ ] `get /api/v1/jobs/{id}/logs` – sse log tail
- [ ] `get /api/v1/jobs/{id}/frames` – paginated frame context
- [ ] `get /api/v1/search?q=...` – full-text on frame json
- [ ] serve static `/dashboard` (from `web/`)
- [ ] openapi spec + swagger ui

---

## 5. queue (`internal/queue/`)
- [ ] interface `queue` with enqueue, dequeue, ack, nack
- [ ] in-memory impl (slice + mutex) for local dev
- [ ] redis impl (`go-redis/v9`) with ttl, retry, dead-letter
- [ ] exponential back-off policy config
- [ ] unit tests with `miniredis`

---

## 6. scheduler (`internal/scheduler/`)
- [ ] gpu node discovery interface
- [ ] fifo + priority queue (high/normal/low)
- [ ] bin-packing (first-fit decreasing)
- [ ] metrics: `scheduler_queue_depth`, `scheduler_jobs_scheduled_total`
- [ ] integration test on kind

---

## 7. transcoder (`internal/transcoder/`)
- [ ] `jobspec` struct (input, outputs, codec, bitrate, enable_llm bool)
- [ ] `worker.run()` – download → ffmpeg → upload
- [ ] ffmpeg progress parser (`frame=`, `fps=`, `time=`)
- [ ] side-car json manifest for multi-bitrate hls
- [ ] pre-signed url generation

---

## 8. llm frame-context (new)
- [ ] `internal/llm/extractor.go` – ffmpeg scene-detect or fixed fps
- [ ] `internal/llm/client.go` – interface + openai impl
- [ ] rate-limiting & retry logic
- [ ] `internal/llm/store.go` – batch insert to postgres
- [ ] migration `migrations/0001_create_frame_context.sql`
- [ ] add secrets for llm key (gcp secret manager)
- [ ] update infra node pool cpu/ram for llm

---

## 9. monitor (`internal/monitor/`)
- [ ] prometheus registry singleton
- [ ] counters: `jobs_total`, `jobs_failed_total`
- [ ] histograms: `job_duration_seconds`, `llm_inference_seconds`
- [ ] structured logs (zap)
- [ ] grafana dashboard json (infra auto-import)

---

## 10. web dashboard (`web/`)
- [ ] `dashboard.html` – bulma css skeleton
- [ ] `js/jobs.js` – auto-refresh table
- [ ] `js/logs.js` – sse modal for live logs
- [ ] `js/frames.js` – timeline scrubber with frame-by-frame context
- [ ] dark-mode toggle
- [ ] bundle with esbuild

---

## 11. infrastructure (`infra/`)
- [ ] `cdktf.json` – typescript cdktf init
- [ ] `infra/gcp/main.ts` – gke autopilot + cloudsql + gcs + secret manager
- [ ] `infra/local/main.ts` – kind + redis + prometheus + grafana
- [ ] shared constructs (`vidcomputeservice`, `vidcomputequeue`)
- [ ] `make infra/local/up` one-command dev stack

---

## 12. docker & compose
- [ ] multi-stage dockerfile (distroless)
- [ ] `.dockerignore`
- [ ] `docker-compose.yml` (local dev)
- [ ] `compose.dev.yml` override (volume mounts, hot-reload)

---

## 13. ci/cd (`.github/workflows/`)
- [ ] `build.yml` – docker build & push
- [ ] `test.yml` – `go test ./...`
- [ ] `infra-plan.yml` – terraform plan on pr
- [ ] `infra-apply.yml` – manual apply
- [ ] `e2e-test.yml` – kind + submit job + assert completion
- [ ] ~~`release.yml` – goreleaser for `vidctl` binaries~~  ➜ **stretch goal / post-MVP CLI tool**

---

## 14. scripts (`scripts/`)
- [ ] `dev-up.sh` – start kind + infra/local
- [ ] `dev-down.sh` – tear down
- [ ] `load-sample-video.sh` – upload test clip
- [ ] `generate-ffmpeg-cmd.sh` – helper script

---

## 15. security & compliance
- [ ] container scan (trivy) in ci
- [ ] slsa provenance via github oidc
- [ ] networkpolicies (deny-all, allow-apiserver)
- [ ] workload identity for gke → gcs/cloudsql

---

## 16. docs & runbooks
- [ ] `docs/architecture.png` – c4 diagram
- [ ] `docs/openapi.yaml` – openapi 3.1
- [ ] `docs/runbook.md` – on-call playbooks
- [ ] `docs/costs.md` – gcp cost calculator
- [ ] `docs/build-log.md` – obsidian daily log

---

## 17. final validation
- [ ] local e2e: upload 5 s clip → encoded + 150 frame rows
- [ ] cloud e2e: same on gke autopilot
- [ ] chaos: kill worker, assert retry & dlq
- [ ] load: 100 concurrent jobs under 5 min

---

---

## Stretch Goals / Post-MVP
- [ ] `vidctl` CLI tool (cobra, submit/status/watch commands)
- [ ] release binaries via goreleaser

# author notes


this is probably good enough for what i want to do for now
