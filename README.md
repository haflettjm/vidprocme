# vidprocme
Video processing pipeline written in golang

# The Why?

    I started interviewing at a company called sieve data.
    I had worked with Image pipelines before but I wanted to do the whole thing from scratch myself so I had a real idea of how these worked in production.
    I figured maybe it would help me better work with the team and would finally take a slot on the projects for my portfolio that isn't a todo app.

# The What?

So basically what we have here is a video or videos of some length that will be broken out into its component frames; into a small llm (I want to host this entire process locally).

#### I need this project to tackle a few things such as;
  - transcode the video files using goroutines using ffmpeg
  - then use a gpu job scheduler with a queue and priority bins
  - have a retryable task runner
  - then a fleshed out monitoring dashboard
  - I also want to make this deployable with cdktf
  - A Workable CI/CD pipeline.

# The Who?
### github@haflettjm . So like Me

# The Where?
### Probably Hybrid: Local + google cloud

# The When?
### Project Start date: July 8th, 2025

```text
vidcompute/
│
├── cmd/               # main.go, CLI / API entrypoints
├── internal/          # queue, scheduler, transcoder, monitor, config, api
├── web/               # dashboard UI
│
├── infra/             # CDKTF app (TypeScript)
│   ├── gcp/           # Google Cloud stack
│   └── local/         # KinD / k3d stack for dev
│
├── .github/workflows/ # GitHub Actions pipelines
├── scripts/           # ffmpeg helpers, local dev scripts
├── Dockerfile
├── docker-compose.yml # local orchestration (no-GPU sim)
└── README.md

```

# The How?

1. User uploads video via API (front end?)
2. Job enters the appropriate queue
3. gpu scheduler bins select available "gpus"
4. Job runs FFmpeg in worker "pool"
5. On failure should retry with backoff
6. Success/failure tracked
7. Dashboard shows queue status and metrics of everything with alerts on failure.
8. Deploy with docker + prometheus + grafana

# Stretch goals
- Job persistence in redis (required)
- cli submission
- gpu node simulation with multiple workers (cluster sim)


# TODO:
TODO.md
