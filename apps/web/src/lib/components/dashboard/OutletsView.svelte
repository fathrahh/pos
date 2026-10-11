<script lang="ts">
  import type { OutletItem } from '../../types';
  import { formatIDR } from '../../data/mock';

  let { outlets }: { outlets: OutletItem[] } = $props();
</script>

<div class="full-content-card">
  <div class="card-header-bar">
    <div>
      <h2 class="card-section-title">Manajemen Cabang (Outlets)</h2>
      <p class="card-section-sub">Daftar lokasi toko fisik yang terhubung dalam tenant bisnis Anda.</p>
    </div>
    <button type="button" class="btn-primary-sm">+ Daftarkan Cabang Baru</button>
  </div>

  <div class="outlets-cards-grid">
    {#each outlets as out (out.id)}
      <div class="outlet-summary-card">
        <h3>{out.name}</h3>
        <p class="outlet-sub-text">{out.address}</p>
        <div class="outlet-card-meta">
          <div>Status: <span class="text-success">Buka ({out.registers} Register Online)</span></div>
          <div>Omzet Hari Ini: <strong>{formatIDR(out.todayRev)}</strong></div>
        </div>
      </div>
    {/each}
  </div>
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

  .outlets-cards-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1rem;
  }

  @media (max-width: 900px) {
    .outlets-cards-grid {
      grid-template-columns: 1fr;
    }
  }

  .outlet-summary-card {
    border: 1px solid var(--gray-2);
    border-radius: var(--radius-md);
    padding: 1.25rem;
    background-color: var(--gray-0);
  }

  .outlet-summary-card h3 {
    margin: 0 0 0.25rem 0;
    font-size: 1rem;
  }

  .outlet-sub-text {
    font-size: 0.75rem;
    color: var(--gray-6);
    margin: 0 0 1rem 0;
  }

  .outlet-card-meta {
    font-size: 0.8125rem;
    display: flex;
    flex-direction: column;
    gap: 0.375rem;
  }

  .text-success {
    color: var(--success);
  }
</style>
