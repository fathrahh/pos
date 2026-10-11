<script lang="ts">
  import type { VoucherItem } from '../../types';
  import { formatIDR } from '../../data/mock';
  import StatusBadge from './StatusBadge.svelte';

  let { vouchers }: { vouchers: VoucherItem[] } = $props();
</script>

<div class="full-content-card">
  <div class="card-header-bar">
    <div>
      <h2 class="card-section-title">Voucher Diskon & Kupon Promosi</h2>
      <p class="card-section-sub">Kelola kode voucher, batasan minimum belanja, dan kuota pemakaian per cabang.</p>
    </div>
    <button type="button" class="btn-primary-sm">+ Buat Voucher Baru</button>
  </div>

  <table class="data-table">
    <thead>
      <tr>
        <th>Kode Kupon</th>
        <th>Nama Program</th>
        <th>Potongan</th>
        <th class="th-right">Min. Belanja</th>
        <th>Kuota Pemakaian</th>
        <th>Status</th>
        <th class="th-right">Aksi</th>
      </tr>
    </thead>
    <tbody>
      {#each vouchers as v (v.code)}
        <tr>
          <td><code class="voucher-code">{v.code}</code></td>
          <td><strong>{v.name}</strong></td>
          <td><strong>{v.discount}</strong></td>
          <td class="td-right num">{formatIDR(v.minSpend)}</td>
          <td class="num">{v.usage}</td>
          <td>
            <StatusBadge
              status={v.status}
              text={v.status === 'active' ? 'Aktif' : 'Berakhir'}
            />
          </td>
          <td class="td-right">
            <button type="button" class="btn-action-link">Detail</button>
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

  .voucher-code {
    font-family: var(--font-mono);
    font-weight: 700;
    font-size: 0.75rem;
    color: var(--accent);
    background-color: var(--accent-tint);
    padding: 0.15rem 0.45rem;
    border-radius: var(--radius-sm);
    display: inline-block;
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
