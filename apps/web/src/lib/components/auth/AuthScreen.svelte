<script lang="ts">
  let { onLoginSuccess }: { onLoginSuccess: () => void } = $props();

  let email = $state('owner@nusantara-fnb.id');
  let password = $state('••••••••••••');
  let showPassword = $state(false);
  let isSubmitting = $state(false);

  function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    isSubmitting = true;
    setTimeout(() => {
      isSubmitting = false;
      onLoginSuccess();
    }, 350);
  }
</script>

<div class="auth-layout">
  <!-- Left Editorial Cover Panel -->
  <aside class="auth-panel-cover">
    <div class="cover-image-container">
      <img
        src="https://images.unsplash.com/photo-1554118811-1e0d58224f24?auto=format&fit=crop&w=1200&q=85"
        alt="Warm architectural cafe counter"
        class="cover-image"
      />
      <div class="cover-overlay"></div>
    </div>

    <div class="cover-topbar">
      <div class="brand-badge-light">
        <div class="brand-logo-placeholder"></div>
        <span class="brand-name-light">Logo Brand</span>
      </div>
    </div>

    <div class="cover-content">
      <h2 class="cover-title">Kelola seluruh cabang dalam satu genggaman.</h2>
      <p class="cover-subtitle">Pusat kendali katalog menu, HPP produk, voucher promo, dan performa omzet semua outlet.</p>
      
      <div class="cover-carousel-dots">
        <span class="dot active"></span>
        <span class="dot"></span>
        <span class="dot"></span>
      </div>
    </div>

    <footer class="cover-footer">
      <span>v2.4 Backoffice Suite</span>
      <span>•</span>
      <span>Owner & Branch Manager Portal</span>
    </footer>
  </aside>

  <!-- Right Form Container -->
  <main class="auth-form-wrap">
    <div class="auth-form-card">
      <div class="form-title-group">
        <h1 class="form-title">Masuk ke Dasbor</h1>
        <p class="form-sub">Akses manajemen terpusat untuk pemilik bisnis & store manager</p>
      </div>

      <!-- SSO Google Button -->
      <div class="sso-group">
        <button type="button" class="sso-btn" onclick={onLoginSuccess}>
          <svg class="sso-icon" viewBox="0 0 24 24">
            <path fill="#EA4335" d="M12 5c1.6 0 3 .6 4.1 1.6l3.1-3.1C17.3 1.8 14.8 1 12 1 7.5 1 3.7 3.6 1.9 7.3l3.7 2.9C6.5 7.4 9 5 12 5z"/>
            <path fill="#4285F4" d="M23.5 12.3c0-.8-.1-1.6-.2-2.3H12v4.6h6.5c-.3 1.5-1.1 2.8-2.4 3.7l3.7 2.9c2.2-2 3.7-5 3.7-8.9z"/>
            <path fill="#FBBC05" d="M5.6 14.8c-.2-.7-.4-1.5-.4-2.3 0-.8.2-1.6.4-2.3L1.9 7.3C.7 9.7 0 12 0 14.5s.7 4.8 1.9 7.2l3.7-2.9z"/>
            <path fill="#34A853" d="M12 23.5c3.2 0 6-1.1 8-3l-3.7-2.9c-1.1.7-2.5 1.2-4.3 1.2-3 0-5.5-2-6.4-4.8L1.9 17c1.8 3.7 5.6 6.5 10.1 6.5z"/>
          </svg>
          <span>Lanjut dengan Google</span>
        </button>
      </div>

      <div class="auth-divider">
        <span class="divider-text">atau dengan email</span>
      </div>

      <form onsubmit={handleSubmit} class="main-form">
        <div class="field-item">
          <label for="email" class="field-label">Email akun bisnis</label>
          <input
            id="email"
            type="email"
            bind:value={email}
            required
            autocomplete="username"
            class="field-input"
            placeholder="owner@brand.com"
          />
        </div>

        <div class="field-item">
          <div class="field-header-row">
            <label for="password" class="field-label">Kata sandi</label>
            <a href="#forgot" class="field-link">Lupa kata sandi?</a>
          </div>
          <div class="input-with-action">
            <input
              id="password"
              type={showPassword ? 'text' : 'password'}
              bind:value={password}
              required
              autocomplete="current-password"
              class="field-input"
            />
            <button
              type="button"
              class="btn-reveal"
              onclick={() => showPassword = !showPassword}
              aria-label={showPassword ? 'Sembunyikan password' : 'Lihat password'}
            >
              {showPassword ? 'Hide' : 'Show'}
            </button>
          </div>
        </div>

        <button type="submit" class="btn-primary" disabled={isSubmitting}>
          {#if isSubmitting}
            Memverifikasi...
          {:else}
            Masuk ke Dasbor
          {/if}
        </button>

        <div class="form-bottom-link">
          <span>Belum memiliki tenant bisnis?</span>
          <a href="#register" class="field-link">Daftarkan brand</a>
        </div>
      </form>
    </div>
  </main>
</div>

<style>
  .auth-layout {
    display: grid;
    grid-template-columns: minmax(480px, 42vw) 1fr;
    width: 100%;
    min-height: 100vh;
  }

  @media (max-width: 900px) {
    .auth-layout {
      grid-template-columns: 1fr;
    }
    .auth-panel-cover {
      display: none;
    }
  }

  .auth-panel-cover {
    position: relative;
    padding: 3.5rem 3rem;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    overflow: hidden;
  }

  .cover-image-container {
    position: absolute;
    inset: 0;
    z-index: 0;
  }

  .cover-image {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .cover-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(
      180deg,
      rgba(15, 23, 42, 0.35) 0%,
      rgba(15, 23, 42, 0.55) 45%,
      rgba(15, 23, 42, 0.90) 100%
    );
  }

  .cover-topbar {
    position: relative;
    z-index: 1;
  }

  .brand-badge-light {
    display: flex;
    align-items: center;
    gap: 0.75rem;
  }

  .brand-logo-placeholder {
    width: 28px;
    height: 28px;
    border-radius: var(--radius-sm);
    background: rgba(255, 255, 255, 0.2);
    border: 1.5px dashed rgba(255, 255, 255, 0.6);
    backdrop-filter: blur(4px);
  }

  .brand-name-light {
    font-size: 1.0625rem;
    font-weight: 600;
    letter-spacing: -0.015em;
    color: #ffffff;
  }

  .cover-content {
    position: relative;
    z-index: 1;
    margin-bottom: 2rem;
  }

  .cover-title {
    font-size: 1.75rem;
    font-weight: 700;
    letter-spacing: -0.025em;
    color: #ffffff;
    line-height: 1.3;
    margin: 0 0 0.75rem 0;
    max-width: 380px;
  }

  .cover-subtitle {
    font-size: 0.875rem;
    color: rgba(255, 255, 255, 0.8);
    line-height: 1.5;
    margin: 0 0 1.75rem 0;
    max-width: 360px;
  }

  .cover-carousel-dots {
    display: flex;
    gap: 0.5rem;
  }

  .cover-carousel-dots .dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: rgba(255, 255, 255, 0.3);
  }

  .cover-carousel-dots .dot.active {
    width: 20px;
    border-radius: 3px;
    background: #ffffff;
  }

  .cover-footer {
    position: relative;
    z-index: 1;
    display: flex;
    align-items: center;
    gap: 0.5rem;
    font-size: 0.75rem;
    color: rgba(255, 255, 255, 0.6);
  }

  .auth-form-wrap {
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2.5rem 1.5rem;
  }

  .auth-form-card {
    width: 100%;
    max-width: 380px;
  }

  .form-title-group {
    margin-bottom: 1.75rem;
  }

  .form-title {
    font-size: 1.625rem;
    font-weight: 700;
    letter-spacing: -0.025em;
    color: var(--gray-9);
    margin: 0 0 0.375rem 0;
  }

  .form-sub {
    font-size: 0.875rem;
    color: var(--gray-6);
    margin: 0;
    line-height: 1.4;
  }

  .sso-group {
    margin-bottom: 1.25rem;
  }

  .sso-btn {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.625rem;
    padding: 0.6875rem 1rem;
    background-color: #ffffff;
    border: 1px solid var(--gray-3);
    border-radius: var(--radius-sm);
    color: var(--gray-9);
    font-size: 0.875rem;
    font-weight: 500;
    cursor: pointer;
    box-shadow: 0 1px 2px oklch(0% 0 0 / 0.04);
    transition: background-color var(--motion-fast), border-color var(--motion-fast);
  }

  .sso-btn:hover {
    background-color: var(--gray-1);
    border-color: var(--gray-6);
  }

  .sso-icon {
    width: 18px;
    height: 18px;
  }

  .auth-divider {
    display: flex;
    align-items: center;
    margin: 1.25rem 0 1.5rem 0;
    position: relative;
  }

  .auth-divider::before {
    content: '';
    position: absolute;
    left: 0;
    right: 0;
    height: 1px;
    background-color: var(--gray-2);
  }

  .divider-text {
    position: relative;
    background-color: var(--gray-0);
    padding: 0 0.75rem;
    margin: 0 auto;
    font-size: 0.75rem;
    color: var(--gray-6);
  }

  .main-form {
    display: flex;
    flex-direction: column;
    gap: 1.125rem;
  }

  .field-item {
    display: flex;
    flex-direction: column;
    gap: 0.375rem;
  }

  .field-header-row {
    display: flex;
    justify-content: space-between;
    align-items: baseline;
  }

  .field-label {
    font-size: 0.8125rem;
    font-weight: 500;
    color: var(--gray-7);
  }

  .field-link {
    font-size: 0.8125rem;
    color: var(--accent);
    text-decoration: none;
    font-weight: 500;
  }

  .field-link:hover {
    text-decoration: underline;
  }

  .field-input {
    width: 100%;
    padding: 0.625rem 0.75rem;
    background-color: #ffffff;
    border: 1px solid var(--gray-3);
    border-radius: var(--radius-sm);
    font-size: 0.875rem;
    color: var(--gray-9);
    outline: none;
    transition: border-color var(--motion-fast), box-shadow var(--motion-fast);
  }

  .field-input:focus {
    border-color: var(--accent);
    box-shadow: 0 0 0 3px var(--accent-tint);
  }

  .input-with-action {
    position: relative;
    display: flex;
    align-items: center;
  }

  .btn-reveal {
    position: absolute;
    right: 0.625rem;
    background: none;
    border: none;
    font-size: 0.75rem;
    font-weight: 600;
    color: var(--gray-6);
    cursor: pointer;
    padding: 0.25rem 0.5rem;
  }

  .btn-primary {
    width: 100%;
    padding: 0.6875rem 1rem;
    background-color: var(--accent);
    color: var(--on-accent);
    border: none;
    border-radius: var(--radius-sm);
    font-size: 0.875rem;
    font-weight: 500;
    cursor: pointer;
    box-shadow: 0 1px 2px oklch(0% 0 0 / 0.1);
    transition: background-color var(--motion-fast);
  }

  .btn-primary:hover:not(:disabled) {
    background-color: var(--accent-hover);
  }

  .btn-primary:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }

  .form-bottom-link {
    margin-top: 0.75rem;
    text-align: center;
    font-size: 0.8125rem;
    color: var(--gray-6);
    display: flex;
    gap: 0.375rem;
    justify-content: center;
  }
</style>
