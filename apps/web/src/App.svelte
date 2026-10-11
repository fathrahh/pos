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

  // Parse initial route including ?redirect=...
  const initial = parseRoute(window.location.pathname, window.location.search);
  let currentStep = $state<FlowState>(initial.step);
  let activeNav = $state<NavSection>(initial.nav);
  let redirectTarget = $state<string | undefined>(initial.redirectUrl);
  let selectedOutlet = $state('all');

  function navigate(step: FlowState, nav: NavSection = 'overview', redirect?: string) {
    currentStep = step;
    activeNav = nav;
    redirectTarget = redirect;

    const url = buildPath(step, nav, redirect);
    const currentFullUrl = window.location.pathname + window.location.search;
    if (currentFullUrl !== url) {
      window.history.pushState({ step, nav, redirect }, '', url);
    }
  }

  function handleLoginSuccess() {
    if (redirectTarget) {
      const target = parseRoute(redirectTarget);
      navigate(target.step, target.nav);
    } else {
      navigate('DASHBOARD', 'overview');
    }
  }

  function handleLogout() {
    // Simpan path aktif saat ini sebagai param ?redirect=...
    const currentDashboardPath = buildPath('DASHBOARD', activeNav);
    navigate('LOGIN', 'overview', currentDashboardPath);
  }

  function onPopState() {
    const route = parseRoute(window.location.pathname, window.location.search);
    currentStep = route.step;
    activeNav = route.nav;
    redirectTarget = route.redirectUrl;
  }

  const products = $state.raw(mockProducts);
  const vouchers = $state.raw(mockVouchers);
  const packages = $state.raw(mockPackages);
  const outlets = $state.raw(mockOutlets);
</script>

<svelte:window onpopstate={onPopState} />

<div class="app-viewport">
  {#if currentStep === 'LOGIN'}
    <AuthScreen onLoginSuccess={handleLoginSuccess} />
  {:else if currentStep === 'DASHBOARD'}
    <div class="dash-shell">
      <Sidebar
        {activeNav}
        onSelectNav={(nav) => navigate('DASHBOARD', nav)}
        onLogout={handleLogout}
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
