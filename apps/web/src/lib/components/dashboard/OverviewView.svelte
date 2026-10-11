<script lang="ts">
  import type { ProductItem, VoucherItem } from '../../types';
  import { formatIDR } from '../../data/mock';
  import StatusBadge from './StatusBadge.svelte';

  let {
    products,
    vouchers,
    onNavigate
  }: {
    products: ProductItem[];
    vouchers: VoucherItem[];
    onNavigate: (section: 'products' | 'vouchers') => void;
  } = $props();
</script>

<div class="overview-container">
  <!-- KPI Metrics Grid -->
  <div class="kpi-grid">
    <div class="metric-card primary">
      <div class="metric-head">
        <span class="metric-label">Total Omzet Hari Ini</span>
        <span class="tag-live">LIVE</span>
      </div>
      <div class="metric-big-num">Rp 18.420.000</div>
      <div class="metric-footnote">+14.2% dari target harian konsolidasi</div>
    </div>

    <div class="metric-card">
      <div class="metric-head">
        <span class="metric-label">Gross Profit (Margin Kotor)</span>
      </div>
      <div class="metric-num">Rp 11.604.600</div>
      <div class="metric-footnote">Rata-rata margin kotor 63.0%</div>
    </div>

    <div class="metric-card">
      <div class="metric-head">
        <span class="metric-label">Total Transaksi</span>
      </div>
      <div class="metric-num">524 <small>struk</small></div>
      <div class="metric-footnote">Rata-rata basket size Rp 35.150</div>
    </div>

    <div class="metric-card">
      <div class="metric-head">
        <span class="metric-label">Alert Inventori</span>
      </div>
      <div class="metric-num text-warning">3 <small>item menipis</small></div>
      <div class="metric-footnote">Butuh restock bahan baku</div>
    </div>
  </div>

  <!-- Content Split: Product Master Preview + Mini Voucher Widget -->
  <div class="dashboard-split">
    <!-- Left Table -->
    <div class="table-container-card">
      <div class="card-header-bar">
        <div>
          <h3 class="card-section-title">Performa Menu Utama</h3>
          <p class="card-section-sub">Daftar harga jual & kalkulasi snapshot HPP (COGS)</p>
        </div>
        <button type="button" class="btn-outline-sm" onclick={() => onNavigate('products')}>
          Lihat Semua Katalog
        </button>
      </div>

      <table class="data-table">
        <thead>
          <tr>
            <th>Nama Menu</th>
            <th>Kategori</th>
            <th class="th-right">HPP (COGS)</th>
            <th class="th-right">Harga Jual</th>
            <th class="th-right">Margin</th>
            <th>Stok</th>
          </tr>
        </thead>
        <tbody>
          {#each products as p (p.id)}
            <tr>
              <td><strong>{p.name}</strong></td>
              <td><span class="badge-category">{p.category}</span></td>
              <td class="td-right num">{formatIDR(p.cogs)}</td>
              <td class="td-right num"><strong>{formatIDR(p.price)}</strong></td>
              <td class="td-right num text-accent">{p.margin}</td>
              <td>
                <StatusBadge
                  status={p.status}
                  text={p.status === 'out' ? 'Habis' : p.status === 'low' ? `Sisa ${p.stock}` : `${p.stock} porsi`}
                />
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>

    <!-- Right Vouchers -->
    <div class="side-card">
      <div class="card-header-bar">
        <div>
          <h3 class="card-section-title">Voucher Aktif</h3>
          <p class="card-section-sub">Promo diskon berjalan</p>
        </div>
        <button type="button" class="btn-outline-sm" onclick={() => onNavigate('vouchers')}>
          + Buat
        </button>
      </div>

      <div class="voucher-list-vertical">
        {#each vouchers.slice(0, 2) as v (v.code)}
          <div class="voucher-item-box">
            <div class="voucher-code-badge">{v.code}</div>
            <div class="voucher-detail">
              <div class="v-name">{v.name}</div>
              <div class="v-meta">Min. belanja {formatIDR(v.minSpend)} • Terpakai: {v.usage}</div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </div>
</div>

<style>
  .overview-container {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  .kpi-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 1rem;
  }

  @media (max-width: 1200px) {
    .kpi-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  .metric-card {
    background-color: #ffffff;
    border: 1px solid var(--gray-2);
    border-radius: var(--radius-md);
    padding: 1.25rem;
    box-shadow: var(--shadow-card);
  }

  .metric-card.primary {
    background-color: var(--accent-tint);
    border-color: var(--accent);
  }

  .metric-head {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 0.5rem;
  }

  .metric-label {
    font-size: 0.75rem;
    color: var(--gray-6);
    font-weight: 600;
  }

  .tag-live {
    font-size: 0.625rem;
    font-weight: 700;
    color: var(--accent);
    background: white;
    padding: 0.125rem 0.375rem;
    border-radius: var(--radius-sm);
  }

  .metric-big-num {
    font-size: 1.625rem;
    font-weight: 700;
    font-family: var(--font-mono);
    color: var(--accent);
    letter-spacing: -0.03em;
    margin-bottom: 0.25rem;
  }

  .metric-num {
    font-size: 1.375rem;
    font-weight: 700;
    font-family: var(--font-mono);
    color: var(--gray-9);
    letter-spacing: -0.02em;
    margin-bottom: 0.25rem;
  }

  .metric-footnote {
    font-size: 0.75rem;
    color: var(--gray-6);
  }

  .dashboard-split {
    display: grid;
    grid-template-columns: 1fr 340px;
    gap: 1.25rem;
  }

  @media (max-width: 1100px) {
    .dashboard-split {
      grid-template-columns: 1fr;
    }
  }

  .table-container-card, .side-card {
    background-color: #ffffff;
    border: 1px solid var(--gray-2);
    border-radius: var(--radius-md);
    box-shadow: var(--shadow-card);
    padding: 1.25rem;
  }

  .card-header-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1.25rem;
  }

  .card-section-title {
    margin: 0 0 0.25rem 0;
    font-size: 1.0625rem;
    font-weight: 600;
    letter-spacing: -0.015em;
  }

  .card-section-sub {
    margin: 0;
    font-size: 0.75rem;
    color: var(--gray-6);
  }

  .data-table {
    width: 100%;
    border-collapse: collapse;
    font-size: 0.8125rem;
  }

  .data-table th {
    text-align: left;
    padding: 0.625rem 0.75rem;
    font-weight: 600;
    color: var(--gray-6);
    border-bottom: 1px solid var(--gray-2);
    font-size: 0.75rem;
  }

  .data-table td {
    padding: 0.75rem;
    border-bottom: 1px solid var(--gray-1);
    color: var(--gray-9);
  }

  .data-table tr:hover td {
    background-color: var(--gray-0);
  }

  .th-right, .td-right {
    text-align: right;
  }

  .num {
    font-family: var(--font-mono);
    font-variant-numeric: tabular-nums;
  }

  .text-accent {
    color: var(--accent);
    font-weight: 600;
  }

  .text-warning {
    color: var(--warning);
  }

  .badge-category {
    background-color: var(--gray-1);
    color: var(--gray-7);
    padding: 0.2rem 0.5rem;
    border-radius: var(--radius-sm);
    font-size: 0.6875rem;
    font-weight: 500;
  }

  .btn-outline-sm {
    background: none;
    border: 1px solid var(--gray-3);
    padding: 0.375rem 0.75rem;
    border-radius: var(--radius-sm);
    font-size: 0.75rem;
    font-weight: 500;
    color: var(--gray-7);
    cursor: pointer;
  }

  .btn-outline-sm:hover {
    background-color: var(--gray-1);
  }

  .voucher-list-vertical {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  .voucher-item-box {
    border: 1px solid var(--gray-2);
    border-radius: var(--radius-sm);
    padding: 0.75rem;
    display: flex;
    flex-direction: column;
    gap: 0.375rem;
  }

  .voucher-code-badge {
    font-family: var(--font-mono);
    font-weight: 700;
    font-size: 0.75rem;
    color: var(--accent);
    background-color: var(--accent-tint);
    padding: 0.15rem 0.45rem;
    border-radius: var(--radius-sm);
    display: inline-block;
    width: fit-content;
  }

  .v-name {
    font-size: 0.8125rem;
    font-weight: 600;
    color: var(--gray-9);
  }

  .v-meta {
    font-size: 0.6875rem;
    color: var(--gray-6);
  }
</style>
