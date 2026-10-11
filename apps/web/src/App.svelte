<script lang="ts">
  import type { FlowState, NavSection } from './lib/types';
  import { mockProducts, mockVouchers, mockPackages, mockOutlets } from './lib/data/mock';
  import { parseRoute, buildPath } from './lib/router';

  import AuthScreen from './lib/components/auth/AuthScreen.svelte';
  import Sidebar from './lib/components/dashboard/Sidebar.svelte';
  import Topbar from './lib/components/dashboard/Topbar.svelte';
  import OverviewView from './lib/components/dashboard/OverviewView.svelte';
  import ProductsView from './lib/components/dashboard/ProductsView.svelte';
  import PackagesView from './lib/components/dashboard/PackagesView.svelte';
  import VouchersView from './lib/components/dashboard/VouchersView.svelte';
  import OutletsView from './lib/components/dashboard/OutletsView.svelte';

  // Read initial route from URL
  const initial = parseRoute(window.location.pathname);
  let currentStep = $state<FlowState>(initial.step);
  let activeNav = $state<NavSection>(initial.nav);
  let selectedOutlet = $state('all');

  // Push new history state when navigating
  function navigate(step: FlowState, nav: NavSection = 'overview') {
    currentStep = step;
    activeNav = nav;
    const url = buildPath(step, nav);
    if (window.location.pathname !== url) {
      window.history.pushState({ step, nav }, '', url);
    }
  }

  // Handle browser Previous / Next navigation buttons
  function onPopState() {
    const route = parseRoute(window.location.pathname);
    currentStep = route.step;
    activeNav = route.nav;
  }

  const products = $state.raw(mockProducts);
  const vouchers = $state.raw(mockVouchers);
  const packages = $state.raw(mockPackages);
  const outlets = $state.raw(mockOutlets);
</script>

<svelte:window onpopstate={onPopState} />

<div class="app-viewport">
  {#if currentStep === 'LOGIN'}
    <AuthScreen onLoginSuccess={() => navigate('DASHBOARD', 'overview')} />
  {:else if currentStep === 'DASHBOARD'}
    <div class="dash-shell">
      <Sidebar
        {activeNav}
        onSelectNav={(nav) => navigate('DASHBOARD', nav)}
        onLogout={() => navigate('LOGIN')}
      />

      <main class="dash-main">
        <Topbar bind:selectedOutlet />

        <div class="dash-body">
          {#if activeNav === 'overview'}
            <OverviewView
              {products}
              {vouchers}
              onNavigate={(section) => navigate('DASHBOARD', section)}
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
