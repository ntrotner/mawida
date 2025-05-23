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
  ADMIN_CUSTOMER_DETAIL = 'admin/customers/{customerId}',
  CUSTOMER = 'customer',
  CUSTOMER_PICKUP = 'customer/pickup?rentContractId={rentContractId}',
  CUSTOMER_RETURN = 'customer/return?rentContractId={rentContractId}',
  CUSTOMER_PROFILE = 'customer/dashboard',
  SHOP = 'shop',
  SHOP_SEARCH = 'shop/search',
  SHOP_PRODUCT = 'shop/product?productId={productId}',
  SHOP_CHECKOUT = 'shop/checkout?productId={productId}&locationId={locationId}',
  SHOP_LOCATION = 'shop/location?locationId={locationId}',
  SETTINGS = 'settings',
  PRIVACY = 'privacy',
  TERMS = 'terms',
  COOKIES = 'cookies',
}