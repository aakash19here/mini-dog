### Designing the bidirectional streaming RPC

The interesting part of bidi streaming isn't the syntax — it's deciding *why* the server needs to talk back while the client is still sending.

## start from the use case, not the call type

ask: "what would a real datadog agent need a long-lived two-way conversation for?"

in log pipelines, the classic answer: **stream logs up while getting per-message acknowledgements (or control signals) back**.

- **client → server stream:** log entries (can reuse the existing entry message)
- **server → client stream:** something the client actually *acts on*. options, in order of how instructive they are to build:

1. **per-log acks** — server replies with the `id` of each log it persisted plus a status.
   - teaches the key bidi insight: the two streams are **independent** — gRPC does NOT enforce a 1:1 request/response pairing. you correlate them yourself via the `id` field.
2. **batched acks** — server acks every N logs or every T seconds.
   - more realistic (datadog-style), forces you to handle "streams move at different rates".
3. **control messages** — server pushes back "slow down" (backpressure), "change log level filter", "stop sending DEBUG".
   - the most datadog-like: agents get remote config over the same channel.

## naming

follow the verb-based pattern already in the service:

| rpc | meaning |
|---|---|
| `SendLog` | one |
| `SubmitLogs` | many, then a summary |
| `StreamLogs` | ongoing conversation (bidi) |

other candidates:
- `TailLogs` — only if the *server* streams logs out (a viewer, like `kubectl logs -f`). that's a **server-streaming** RPC — a different, also worthwhile addition.
- `LogSession` / `Collect` — if going the "long-lived agent session with acks + control" route.

## message shapes to consider

- request stream can reuse the existing entry message.
- naming smell to notice: `LogEntryRequest` is doing double duty across three RPCs now. many people would name the data itself `LogEntry` and keep `Request`/`Response` wrappers per-RPC. not urgent.
- for the response stream, one message type can carry different meanings:

```proto
message StreamResponse {
  oneof payload {
    LogAck ack = 1;
    ControlSignal control = 2;
  }
}
```

very idiomatic way to multiplex "acks" and "commands" over one server stream.

- give the ack a way to point back at what it acknowledges: log `id` (per-log) or `last_id` / count (batched). without that field the client can't do anything useful with the ack.

## questions to answer before writing the proto

1. does the server respond once per log, per batch, or only on events? → decides the response message shape.
2. what does the client *do* with a failed ack — retry, drop, buffer? → decides what status info the ack needs.
3. who closes the stream first, and what happens to in-flight logs? → bidi's trickiest part in the Go implementation: the send and recv loops usually live in **separate goroutines**.

## suggested first pass

```proto
rpc StreamLogs(stream LogEntryRequest) returns (stream LogAck);
```

per-log acks keyed by `id` — the smallest design that still forces learning the correlation + concurrent-loops parts. then evolve the ack into a `oneof` with control signals once that works.
