<script lang="ts">
  let {
    selectedOutlet = $bindable()
  }: {
    selectedOutlet: string;
  } = $props();
</script>

<header class="dash-topbar">
  <div class="topbar-left">
    <div class="outlet-dropdown-wrap">
      <label for="outlet-select" class="dropdown-label">Cabang Aktif:</label>
      <select id="outlet-select" bind:value={selectedOutlet} class="outlet-select">
        <option value="all">Semua Cabang (Konsolidasi)</option>
        <option value="kemang">Cabang Kemang, Jaksel</option>
        <option value="senopati">Cabang Senopati, Jaksel</option>
        <option value="bandung">Cabang Riau, Bandung</option>
      </select>
    </div>
  </div>

  <div class="topbar-right">
    <span class="user-chip">
      <span class="user-avatar-dot"></span>
      <span>Budi Santoso (Owner)</span>
    </span>
  </div>
</header>

<style>
  .dash-topbar {
    height: 56px;
    background-color: #ffffff;
    border-bottom: 1px solid var(--gray-2);
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 1.75rem;
  }

  .outlet-dropdown-wrap {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }

  .dropdown-label {
    font-size: 0.8125rem;
    color: var(--gray-6);
    font-weight: 500;
  }

  .outlet-select {
    padding: 0.375rem 0.625rem;
    border: 1px solid var(--gray-2);
    border-radius: var(--radius-sm);
    background-color: var(--gray-0);
    color: var(--gray-9);
    font-size: 0.8125rem;
    font-weight: 600;
    cursor: pointer;
  }

  .user-chip {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-size: 0.8125rem;
    font-weight: 500;
    color: var(--gray-7);
  }

  .user-avatar-dot {
    width: 8px;
    height: 8px;
    border-radius: 50%;
    background-color: var(--success);
  }
</style>
