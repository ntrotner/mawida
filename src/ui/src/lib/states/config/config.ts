import { Observable, distinctUntilChanged, map } from "rxjs";
import { DefaultState } from "../common/state";
import type { Config, Configuration } from "./model";
import { defaultLoggerConfig } from "./collection/logger";
import { defaultAppConfig } from "../status";

/**
 * ConfigState class is used to manage the state of the configs.
 */
export class ConfigState extends DefaultState<Configuration> {

  constructor() {
    super();
    this.setState({loaded: false, configMap: {
      ...defaultAppConfig,
      ...defaultLoggerConfig
    }})
  }

  public isLoaded(): Observable<boolean> {
    return this.observable().pipe(
      map(configs => configs.loaded),
    ) as Observable<boolean>
  }

  /**
   * Updates configuration and indicates that the configs were loaded
   * 
   * @param configMap 
   */
  public setConfigMap(configMap: Record<string, Config>): void {
    const previousState = this.getSyncState();
    this.setState({ ...(previousState || {}), loaded: true, configMap: {...(previousState?.configMap || {}), ...configMap} });
  }

  /**
   * Creates an observable for the requested config
   * 
   * @param configName 
   * @returns undefined value if configs are not ready
   */
  public getConfig<T>(configName: string): Observable<T | undefined> {
    return this.observable().pipe(
      map(configs => configs.configMap[configName]),
      distinctUntilChanged((a, b) => JSON.stringify(a) === JSON.stringify(b))
    ) as Observable<T | undefined>
  }
}
export const configState = new ConfigState();

export const CONFIG_STATE = 'configState';