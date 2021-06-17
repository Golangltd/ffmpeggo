package log_ffmpeggo

// 错误日志使用，主要是保证ffmpeg进程正常情况下，无需重启等活动；
// 如果发生故障，可以直接调用重启服务