import { filter, switchMap, take } from "rxjs";
import { configState } from "../../states/config";
import { type LoggerConfig, LoggerConfigKey } from "../../states/config/collection/logger";
import { MessageType, type Log } from "./model";

class Logger {
  enabled: boolean = false;
  config: LoggerConfig = {};
  constructor() {
    configState.isLoaded().pipe(
      filter(isLoaded => isLoaded),
      take(1),
      switchMap(() => configState.getConfig<LoggerConfig>(LoggerConfigKey))
    ).subscribe(config => {
      if (config?.url) {
        this.enabled = true;
        this.config = config;
      }
    })
  }

  private logFactory(type: MessageType, message: string): Log {
    return {
      message,
      type,
      timestamp: (new Date()).getTime()
    }
  }

  private async sendLog(messageType: Log): Promise<boolean> {
    if (this.enabled && this.config.url) {
      try {
        const response = await fetch(this.config.url, { method: 'POST', body: JSON.stringify(messageType) });
        return response.status === 200;
      } catch {
        return false;
      }
    }
    return false;
  }

  public debug(msg: string): Promise<boolean> {
    const newLog = this.logFactory(MessageType.debug, msg);
    return this.sendLog(newLog);
  }

  public info(msg: string): Promise<boolean> {
    const newLog = this.logFactory(MessageType.info, msg);
    return this.sendLog(newLog);
  }

  public error(msg: string): Promise<boolean> {
    const newLog = this.logFactory(MessageType.error, msg);
    return this.sendLog(newLog);
  }

}

export const logger = new Logger();