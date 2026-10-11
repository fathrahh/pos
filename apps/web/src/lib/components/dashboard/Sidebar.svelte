<script lang="ts">
  import type { NavSection } from '../../types';

  let {
    activeNav,
    onSelectNav,
    onLogout
  }: {
    activeNav: NavSection;
    onSelectNav: (nav: NavSection) => void;
    onLogout: () => void;
  } = $props();
</script>

<aside class="dash-sidebar">
  <div class="sidebar-brand">
    <div class="brand-logo-sq"></div>
    <div class="brand-text-col">
      <span class="brand-title">Kopi Nusantara</span>
      <span class="brand-sub">Tenant Owner</span>
    </div>
  </div>

  <nav class="sidebar-nav">
    <button
      type="button"
      class="nav-link"
      class:active={activeNav === 'overview'}
      onclick={() => onSelectNav('overview')}
    >
      <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <rect x="3" y="3" width="7" height="7"/>
        <rect x="14" y="3" width="7" height="7"/>
        <rect x="14" y="14" width="7" height="7"/>
        <rect x="3" y="14" width="7" height="7"/>
      </svg>
      <span>Overview & Finansial</span>
    </button>

    <button
      type="button"
      class="nav-link"
      class:active={activeNav === 'products'}
      onclick={() => onSelectNav('products')}
    >
      <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-2z"/>
        <line x1="3" y1="6" x2="21" y2="6"/>
        <path d="M16 10a4 4 0 0 1-8 0"/>
      </svg>
      <span>Produk & HPP</span>
    </button>

    <button
      type="button"
      class="nav-link"
      class:active={activeNav === 'packages'}
      onclick={() => onSelectNav('packages')}
    >
      <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <polygon points="12 2 2 7 12 12 22 7 12 2"/>
        <polyline points="2 17 12 22 22 17"/>
        <polyline points="2 12 12 17 22 12"/>
      </svg>
      <span>Paket & Bundle</span>
    </button>

    <button
      type="button"
      class="nav-link"
      class:active={activeNav === 'vouchers'}
      onclick={() => onSelectNav('vouchers')}
    >
      <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <rect x="2" y="5" width="20" height="14" rx="2"/>
        <line x1="2" y1="10" x2="22" y2="10"/>
      </svg>
      <span>Voucher Diskon</span>
    </button>

    <button
      type="button"
      class="nav-link"
      class:active={activeNav === 'outlets'}
      onclick={() => onSelectNav('outlets')}
    >
      <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
        <polyline points="9 22 9 12 15 12 15 22"/>
      </svg>
      <span>Manajemen Cabang</span>
    </button>
  </nav>

  <div class="sidebar-footer">
    <button type="button" class="btn-logout" onclick={onLogout}>
      <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
        <polyline points="16 17 21 12 16 7"/>
        <line x1="21" y1="12" x2="9" y2="12"/>
      </svg>
      <span>Keluar Akun</span>
    </button>
  </div>
</aside>

<style>
  .dash-sidebar {
    width: 240px;
    background-color: var(--gray-1);
    border-right: 1px solid var(--gray-2);
    display: flex;
    flex-direction: column;
    padding: 1.5rem 1rem;
    flex-shrink: 0;
  }

  .sidebar-brand {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0 0.5rem 1.5rem 0.5rem;
    border-bottom: 1px solid var(--gray-2);
  }

  .brand-logo-sq {
    width: 28px;
    height: 28px;
    border-radius: var(--radius-sm);
    background-color: var(--accent);
  }

  .brand-text-col {
    display: flex;
    flex-direction: column;
  }

  .brand-title {
    font-size: 0.9375rem;
    font-weight: 700;
    color: var(--gray-9);
    letter-spacing: -0.015em;
  }

  .brand-sub {
    font-size: 0.6875rem;
    color: var(--gray-6);
    font-family: var(--font-mono);
  }

  .sidebar-nav {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
    padding: 1.25rem 0;
    flex: 1;
  }

  .nav-link {
    display: flex;
    align-items: center;
    gap: 0.75rem;
    padding: 0.625rem 0.75rem;
    background: none;
    border: none;
    border-radius: var(--radius-sm);
    font-size: 0.8125rem;
    font-weight: 500;
    color: var(--gray-7);
    cursor: pointer;
    text-align: left;
    transition: all var(--motion-fast);
  }

  .nav-link:hover {
    background-color: var(--gray-2);
    color: var(--gray-9);
  }

  .nav-link.active {
    background-color: var(--accent-tint-strong);
    color: var(--accent);
    font-weight: 600;
  }

  .nav-icon {
    width: 17px;
    height: 17px;
  }

  .sidebar-footer {
    padding-top: 1rem;
    border-top: 1px solid var(--gray-2);
  }

  .btn-logout {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    width: 100%;
    padding: 0.5rem 0.75rem;
    background: none;
    border: none;
    color: var(--gray-6);
    font-size: 0.75rem;
    font-weight: 500;
    cursor: pointer;
    border-radius: var(--radius-sm);
  }

  .btn-logout:hover {
    background-color: var(--gray-2);
    color: var(--danger);
  }
</style>
