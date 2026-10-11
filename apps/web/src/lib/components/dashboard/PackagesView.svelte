<script lang="ts">
  import type { PackageItem } from '../../types';
  import { formatIDR } from '../../data/mock';

  let { packages }: { packages: PackageItem[] } = $props();
</script>

<div class="full-content-card">
  <div class="card-header-bar">
    <div>
      <h2 class="card-section-title">Paket Menu & Combo Bundle</h2>
      <p class="card-section-sub">Gabungkan beberapa item menjadi paket promo dengan harga spesial.</p>
    </div>
    <button type="button" class="btn-primary-sm">+ Buat Paket Baru</button>
  </div>

  <table class="data-table">
    <thead>
      <tr>
        <th>Nama Paket</th>
        <th>Isi Paket (Menu)</th>
        <th class="th-right">Total HPP</th>
        <th class="th-right">Harga Normal</th>
        <th class="th-right">Harga Paket</th>
        <th class="th-right">Aksi</th>
      </tr>
    </thead>
    <tbody>
      {#each packages as pkg (pkg.id)}
        <tr>
          <td><strong>{pkg.name}</strong></td>
          <td><span class="badge-category">{pkg.items}</span></td>
          <td class="td-right num">{formatIDR(pkg.cogs)}</td>
          <td class="td-right num strike">{formatIDR(pkg.normalPrice)}</td>
          <td class="td-right num text-accent"><strong>{formatIDR(pkg.bundlePrice)}</strong></td>
          <td class="td-right">
            <button type="button" class="btn-action-link">Edit Paket</button>
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

  .strike {
    text-decoration: line-through;
    color: var(--gray-6);
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
