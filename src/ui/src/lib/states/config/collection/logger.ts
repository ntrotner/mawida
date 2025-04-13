export interface LoggerConfig {
  url?: string
}

export const LoggerConfigKey = 'logger';

export const defaultLoggerConfig: { [LoggerConfigKey]: LoggerConfig } = {
  [LoggerConfigKey]: {}
}
