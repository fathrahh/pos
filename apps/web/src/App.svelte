<script lang="ts">
  import type { FlowState, NavSection } from './lib/types';
  import { mockProducts, mockVouchers, mockPackages, mockOutlets } from './lib/data/mock';

  import AuthScreen from './lib/components/auth/AuthScreen.svelte';
  import Sidebar from './lib/components/dashboard/Sidebar.svelte';
  import Topbar from './lib/components/dashboard/Topbar.svelte';
  import OverviewView from './lib/components/dashboard/OverviewView.svelte';
  import ProductsView from './lib/components/dashboard/ProductsView.svelte';
  import PackagesView from './lib/components/dashboard/PackagesView.svelte';
  import VouchersView from './lib/components/dashboard/VouchersView.svelte';
  import OutletsView from './lib/components/dashboard/OutletsView.svelte';

  let currentStep = $state<FlowState>('LOGIN');
  let activeNav = $state<NavSection>('overview');
  let selectedOutlet = $state('all');

  const products = $state.raw(mockProducts);
  const vouchers = $state.raw(mockVouchers);
  const packages = $state.raw(mockPackages);
  const outlets = $state.raw(mockOutlets);
</script>

<div class="app-viewport">
  {#if currentStep === 'LOGIN'}
    <AuthScreen onLoginSuccess={() => currentStep = 'DASHBOARD'} />
  {:else if currentStep === 'DASHBOARD'}
    <div class="dash-shell">
      <Sidebar
        {activeNav}
        onSelectNav={(nav) => activeNav = nav}
        onLogout={() => currentStep = 'LOGIN'}
      />

      <main class="dash-main">
        <Topbar bind:selectedOutlet />

        <div class="dash-body">
          {#if activeNav === 'overview'}
            <OverviewView
              {products}
              {vouchers}
              onNavigate={(section) => activeNav = section}
            />
          {:else if activeNav === 'products'}
            <ProductsView {products} />
          {:else if activeNav === 'packages'}
            <PackagesView {packages} />
          {:else if activeNav === 'vouchers'}
            <VouchersView {vouchers} />
          {:else if activeNav === 'outlets'}
            <OutletsView {outlets} />
          {/if}
        </div>
      </main>
    </div>
  {/if}
</div>

<style>
  .app-viewport {
    width: 100%;
    min-height: 100vh;
    display: flex;
    background-color: var(--gray-0);
  }

  .dash-shell {
    display: flex;
    width: 100%;
    min-height: 100vh;
  }

  .dash-main {
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow-x: hidden;
  }

  .dash-body {
    padding: 1.75rem;
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }
</style>
