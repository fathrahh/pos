<script lang="ts">
  let {
    status,
    text
  }: {
    status: 'active' | 'low' | 'out' | 'expired';
    text: string;
  } = $props();
</script>

{#if status === 'out' || status === 'expired'}
  <span class="badge-status out">{text}</span>
{:else if status === 'low'}
  <span class="badge-status low">{text}</span>
{:else}
  <span class="badge-status active">{text}</span>
{/if}

<style>
  .badge-status {
    display: inline-block;
    padding: 0.2rem 0.5rem;
    border-radius: var(--radius-sm);
    font-size: 0.6875rem;
    font-weight: 600;
  }

  .badge-status.active {
    background-color: oklch(52% 0.14 155 / 0.12);
    color: var(--success);
  }

  .badge-status.low {
    background-color: oklch(62% 0.14 75 / 0.15);
    color: var(--warning);
  }

  .badge-status.out {
    background-color: oklch(52% 0.18 25 / 0.12);
    color: var(--danger);
  }
</style>
