<script lang="ts">
  let healthStatus = $state<string>('checking...');
  let tenantId = $state<string>('tenant-demo-123');

  async function checkBackend() {
    try {
      const res = await fetch('/api/v1/');
      if (res.ok) {
        const json = await res.json();
        healthStatus = json.data?.message || 'online';
      } else {
        healthStatus = 'error: ' + res.statusText;
      }
    } catch (e: any) {
      healthStatus = 'offline / failed to connect';
    }
  }

  $effect(() => {
    checkBackend();
  });
</script>

<main style="padding: 2rem; max-width: 800px; margin: 0 auto;">
  <header style="border-bottom: 1px solid #e5e7eb; padding-bottom: 1rem; margin-bottom: 2rem;">
    <h1 style="margin: 0; font-size: 1.875rem; color: #1f2937;">POS SaaS Dashboard</h1>
    <p style="margin: 0.5rem 0 0 0; color: #6b7280;">Svelte 5 SPA + Golang Clean Architecture</p>
  </header>

  <div style="background: white; border: 1px solid #e5e7eb; border-radius: 8px; padding: 1.5rem; box-shadow: 0 1px 3px rgba(0,0,0,0.1);">
    <h2 style="margin-top: 0; font-size: 1.25rem;">Backend Connection Status</h2>
    <p>API Status: <strong style="color: {healthStatus === 'online' || healthStatus.includes('API') ? '#059669' : '#dc2626'}">{healthStatus}</strong></p>
    <p>Active Tenant: <code>{tenantId}</code></p>
    <button
      style="padding: 0.5rem 1rem; background: #2563eb; color: white; border: none; border-radius: 6px; cursor: pointer;"
      onclick={checkBackend}
    >
      Re-check API
    </button>
  </div>
</main>
