import i18n from 'sveltekit-i18n';

export const languages = [
  {key: 'common.en', locale: 'en'},
  {key: 'common.de', locale: 'de'}
];

const config = ({
  loaders: [
    {
      locale: 'en',
      key: 'admin',
      loader: async () => (
        await import('./languages/admin/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'admin',
      loader: async () => (
        await import('./languages/admin/de.json')
      ).default,
    },
    {
      locale: 'en',
      key: 'common',
      loader: async () => (
        await import('./languages/common/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'common',
      loader: async () => (
        await import('./languages/common/de.json')
      ).default,
    },
    {
      locale: 'en',
      key: 'home',
      loader: async () => (
        await import('./languages/home/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'home',
      loader: async () => (
        await import('./languages/home/de.json')
      ).default,
    },
    {
      locale: 'en',
      key: 'login',
      loader: async () => (
        await import('./languages/login/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'login',
      loader: async () => (
        await import('./languages/login/de.json')
      ).default,
    },
    {
      locale: 'en',
      key: 'logout',
      loader: async () => (
        await import('./languages/logout/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'logout',
      loader: async () => (
        await import('./languages/logout/de.json')
      ).default,
    },
    {
      locale: 'en',
      key: 'profile',
      loader: async () => (
        await import('./languages/profile/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'profile',
      loader: async () => (
        await import('./languages/profile/de.json')
      ).default,
    },
    {
      locale: 'en',
      key: 'product',
      loader: async () => (
        await import('./languages/product/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'product',
      loader: async () => (
        await import('./languages/product/de.json')
      ).default,
    },
    {
      locale: 'en',
      key: 'shop',
      loader: async () => (
        await import('./languages/shop/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'shop',
      loader: async () => (
        await import('./languages/shop/de.json')
      ).default,
    },
    {
      locale: 'en',
      key: 'location',
      loader: async () => (
        await import('./languages/location/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'location',
      loader: async () => (
        await import('./languages/location/de.json')
      ).default,
    },
    {
      locale: 'en',
      key: 'customer',
      loader: async () => (
        await import('./languages/customer/en.json')
      ).default,
    },
    {
      locale: 'de',
      key: 'customer',
      loader: async () => (
        await import('./languages/customer/de.json')
      ).default,
    }
  ],
});
export const { t, loading, locales, locale, translations, loadTranslations, addTranslations, setLocale, setRoute } = new i18n(config);