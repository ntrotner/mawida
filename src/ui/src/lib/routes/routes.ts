export enum ROUTES {
  HOME = '',
  LOGIN = 'login',
  LOGOUT = 'logout',
  PROFILE = 'profile',
  ADMIN = 'admin',
  ADMIN_LOCATIONS = 'admin/locations',
  ADMIN_LOCATIONS_CREATE = 'admin/locations/create',
  ADMIN_PRODUCT_CREATE = 'admin/locations/{locationId}/products/create',
  ADMIN_PRODUCTS = 'admin/locations/{locationId}/products',
  ADMIN_CUSTOMERS = 'admin/customers',
  ADMIN_STATISTICS = 'admin/statistics',
}