# Z-stream backport (RHAI-50090)

Pinned copy of `github.com/llm-d/llm-d-inference-payload-processor` at
`a8bbe6a6a75a` (`v0.1.0-rc.4.0.20260721200012-a8bbe6a6a75a`) plus the
equivalent of opendatahub-io/ai-gateway-payload-processing#494:

- default gRPC `MaxRecvMsgSize` raised to 32 MiB
- optional `GRPC_MAX_RECV_MSG_SIZE` env override

On `main`/`rhoai` the runner is in-tree (`cmd/runner.go`), so #494 landed
there. On `rhoai-3.5` the runner is still this external module; a literal
cherry-pick of #494 does not apply. This in-tree replace carries the same
behavior for Konflux builds without depending on a personal fork at build time.
