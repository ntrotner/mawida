export interface AppConfig {
  isBackendAware?: boolean,
  user?: boolean,
  navigation?: 'bulky' | 'slim' | 'none'
}

export const AppConfigKey = 'app';

export const defaultAppConfig: { [AppConfigKey]: AppConfig } = {
  [AppConfigKey]: {isBackendAware: true, user: true, navigation: 'slim'}
}
