package config

type Options struct {
    MaxConcurrentPerRouter int      // Max concurrent requests per worker
    QueueSize              int      // Channel buffer size per worker
    EnableLogging          bool     // Whether to enable verbose logging
    GracefulTimeoutSeconds int      // Request processing timeout
    MaxWorkers             int      // Maximum number of workers to spawn
    WorkerGroupSize        int      // Number of workers per group
    DBInitFunc             func(groupID string) any // DB init per group
}