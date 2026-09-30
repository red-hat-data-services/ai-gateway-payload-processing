# Provider Guides

Configuration guides for each supported external model provider.

| Provider | Type | Translation | Auth | Guide |
|----------|------|------------|------|-------|
| [OpenAI](openai.md) | `openai` | Pass-through | `Authorization: Bearer` | [openai.md](openai.md) |
| [Anthropic](anthropic.md) | `anthropic` | OpenAI ↔ Messages API | `x-api-key` | [anthropic.md](anthropic.md) |
| [AWS Bedrock](bedrock-openai.md) | `bedrock-openai` | Pass-through (Mantle) | `Authorization: Bearer` | [bedrock-openai.md](bedrock-openai.md) |
| [Azure OpenAI](azure-openai.md) | `azure-openai` | Path rewrite + field stripping | `api-key` | [azure-openai.md](azure-openai.md) |
| [Vertex AI](vertex-openai.md) | `vertex-openai` | Path rewrite + field stripping | `Authorization: Bearer` (OAuth) | [vertex-openai.md](vertex-openai.md) |

## Credential secrets

Every provider's API-key Secret **must** carry the label:

```yaml
metadata:
  labels:
    inference.llm-d.ai/ipp-managed: "true"
```

The apikey-injection plugin's reconciler only watches Secrets with this label
(`pkg/plugins/apikey-injection/reconciler.go`). A Secret without it is invisible
to the credential store even if the `ExternalProvider`'s `secretRef` points at it,
and every request to that provider fails with:

```
HTTP 500 — inference error: Internal - authType 'apikey' credentials not found
```

To fix an existing unlabeled Secret:

```bash
kubectl label secret <name> -n <namespace> inference.llm-d.ai/ipp-managed=true
```

## Transport annotations

`ExternalProvider` supports optional annotations to control upstream transport:

| Annotation | Default | Description |
|------------|---------|-------------|
| `inference.opendatahub.io/port` | `443` | Backend port (1-65535) |
| `inference.opendatahub.io/tls` | `"true"` | TLS origination: `"true"` or `"false"` |

Notes:
- `inference.opendatahub.io/tls` must be exactly `"true"` or `"false"` (lowercase string)
- Disabling TLS (`"false"`) is not allowed with `apikey` or `sigv4` auth
- Legacy annotations `maas.opendatahub.io/port` and `maas.opendatahub.io/tls` are migrated to the new keys
