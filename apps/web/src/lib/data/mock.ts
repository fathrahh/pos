import type { ProductItem, VoucherItem, PackageItem, OutletItem } from '../types';

export const mockProducts: ProductItem[] = [
  { id: 'prd-01', name: 'Signature Aren Latte', category: 'Kopi Susu', cogs: 8500, price: 24000, margin: '64.5%', stock: 142, status: 'active' },
  { id: 'prd-02', name: 'Cold Brew Oat Milk', category: 'Kopi Hitam', cogs: 11000, price: 32000, margin: '65.6%', stock: 84, status: 'active' },
  { id: 'prd-03', name: 'Butter Croissant Artisanal', category: 'Pastry', cogs: 9500, price: 22000, margin: '56.8%', stock: 19, status: 'low' },
  { id: 'prd-04', name: 'Matcha Kyoto Fusion', category: 'Non-Coffee', cogs: 12000, price: 28000, margin: '57.1%', stock: 65, status: 'active' },
  { id: 'prd-05', name: 'Pain au Chocolat', category: 'Pastry', cogs: 11500, price: 26000, margin: '55.7%', stock: 0, status: 'out' }
];

export const mockVouchers: VoucherItem[] = [
  { code: 'SENINSEGAR', name: 'Diskon 20% Awal Pekan', discount: '20%', minSpend: 50000, usage: '184 / 300', status: 'active' },
  { code: 'BARISTAFAV', name: 'Potongan Flat Rp 10.000', discount: 'Rp 10.000', minSpend: 60000, usage: '92 / 150', status: 'active' },
  { code: 'PAYDAYFEST', name: 'Cashback Gajian Member', discount: '25%', minSpend: 100000, usage: 'Expired', status: 'expired' }
];

export const mockPackages: PackageItem[] = [
  { id: 'pkg-01', name: 'Paket Sarapan Hemat', items: 'Aren Latte + Croissant', cogs: 18000, bundlePrice: 38000, normalPrice: 46000 },
  { id: 'pkg-02', name: 'Work From Cafe Duo', items: '2x Cold Brew + 2x Brownies', cogs: 32000, bundlePrice: 75000, normalPrice: 92000 }
];

export const mockOutlets: OutletItem[] = [
  { id: 'out-1', name: 'Cabang Kemang, Jakarta Selatan', address: 'Jl. Kemang Raya No. 14', registers: 2, todayRev: 7850000 },
  { id: 'out-2', name: 'Cabang Senopati, Jakarta Selatan', address: 'Jl. Senopati No. 88', registers: 3, todayRev: 8120000 },
  { id: 'out-3', name: 'Cabang Riau, Bandung', address: 'Jl. L.L.R.E. Martadinata No. 42', registers: 1, todayRev: 2450000 }
];

export function formatIDR(num: number): string {
  return 'Rp ' + num.toLocaleString('id-ID');
}
