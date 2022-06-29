# Elastic Agent Shipper Client API

## Publication acknowledgements

Some inputs need a way to persist their position in a data stream, such as a log file or a Kafka topic. To guarantee at-least-once delivery of these events, it's not enough to just send a publish request to the shipper, since the shipper may shut down before successfully publishing all its queued events, and in this case the input's persisted position should not include events that were dropped during shutdown.

To support this, the shipper process maintains an internal ordering of queued events by an ascending ID. With each API call, the shipper reports the position of the highest sequential event ID that has been "persisted" -- either written to the configured output, or written to disk when the disk queue is in use (meaning that even if it is not published to an output during this run, it has been saved and will be published the next time the shipper starts). Once an event is reported as persisted, an input may safely update its internal position to reflect that those events have been processed.

Because the IDs assigned to events are specific to the shipper process, they are not preserved between restarts. To help inputs recognize and invalidate old IDs when the shipper restarts, the shipper process is assigned a UUID on startup, which is reported in API responses. While the shipper doesn't restart under normal operation, this gives inputs a way to provide additional robustness when the system is recovering from an error.

### Technical considerations

- Lack of metadata
- One event ID per publish request
- Keying the event ordering by the shipper process UUID
- 