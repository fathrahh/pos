export type FlowState = 'LOGIN' | 'DASHBOARD';

export type NavSection = 'overview' | 'products' | 'packages' | 'vouchers' | 'outlets';

export interface ProductItem {
  id: string;
  name: string;
  category: string;
  cogs: number;
  price: number;
  margin: string;
  stock: number;
  status: 'active' | 'low' | 'out';
}

export interface VoucherItem {
  code: string;
  name: string;
  discount: string;
  minSpend: number;
  usage: string;
  status: 'active' | 'expired';
}

export interface PackageItem {
  id: string;
  name: string;
  items: string;
  cogs: number;
  bundlePrice: number;
  normalPrice: number;
}

export interface OutletItem {
  id: string;
  name: string;
  address: string;
  registers: number;
  todayRev: number;
}
