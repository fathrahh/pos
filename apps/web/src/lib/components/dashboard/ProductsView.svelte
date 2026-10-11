<script lang="ts">
  import type { ProductItem } from '../../types';
  import { formatIDR } from '../../data/mock';
  import StatusBadge from './StatusBadge.svelte';

  let { products }: { products: ProductItem[] } = $props();
</script>

<div class="full-content-card">
  <div class="card-header-bar">
    <div>
      <h2 class="card-section-title">Katalog Produk & Snapshot HPP</h2>
      <p class="card-section-sub">Atur harga retail, kontrol HPP per cabang, dan pantau margin laba per item.</p>
    </div>
    <div class="btn-group">
      <button type="button" class="btn-primary-sm">+ Tambah Produk Baru</button>
    </div>
  </div>

  <table class="data-table">
    <thead>
      <tr>
        <th>Menu</th>
        <th>Kategori</th>
        <th class="th-right">HPP Bahan Baku</th>
        <th class="th-right">Harga Jual Retail</th>
        <th class="th-right">Margin Laba</th>
        <th>Status Stok</th>
        <th class="th-right">Aksi</th>
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
          <td class="td-right">
            <button type="button" class="btn-action-link">Ubah Harga</button>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>

<style>
  .full-content-card {
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

  .badge-category {
    background-color: var(--gray-1);
    color: var(--gray-7);
    padding: 0.2rem 0.5rem;
    border-radius: var(--radius-sm);
    font-size: 0.6875rem;
    font-weight: 500;
  }

  .btn-primary-sm {
    background-color: var(--accent);
    color: var(--on-accent);
    border: none;
    padding: 0.45rem 0.875rem;
    border-radius: var(--radius-sm);
    font-size: 0.75rem;
    font-weight: 600;
    cursor: pointer;
  }

  .btn-action-link {
    background: none;
    border: none;
    color: var(--accent);
    font-weight: 600;
    cursor: pointer;
    font-size: 0.75rem;
  }
</style>
