<script lang="ts">
  // Flow states
  type FlowState = 'LOGIN' | 'DASHBOARD';
  let currentStep = $state<FlowState>('LOGIN');

  // Login form states
  let email = $state('owner@nusantara-fnb.id');
  let password = $state('••••••••••••');
  let showPassword = $state(false);
  let isSubmitting = $state(false);

  // Dashboard nav & filter
  type NavSection = 'overview' | 'products' | 'pricing' | 'vouchers' | 'packages' | 'outlets';
  let activeNav = $state<NavSection>('overview');
  let selectedOutlet = $state('all');

  // Raw mock state for performance (avoid deep proxy on read-only arrays)
  const products = $state.raw([
    { id: 'prd-01', name: 'Signature Aren Latte', category: 'Kopi Susu', cogs: 8500, price: 24000, margin: '64.5%', stock: 142, status: 'active' },
    { id: 'prd-02', name: 'Cold Brew Oat Milk', category: 'Kopi Hitam', cogs: 11000, price: 32000, margin: '65.6%', stock: 84, status: 'active' },
    { id: 'prd-03', name: 'Butter Croissant Artisanal', category: 'Pastry', cogs: 9500, price: 22000, margin: '56.8%', stock: 19, status: 'low' },
    { id: 'prd-04', name: 'Matcha Kyoto Fusion', category: 'Non-Coffee', cogs: 12000, price: 28000, margin: '57.1%', stock: 65, status: 'active' },
    { id: 'prd-05', name: 'Pain au Chocolat', category: 'Pastry', cogs: 11500, price: 26000, margin: '55.7%', stock: 0, status: 'out' }
  ]);

  const vouchers = $state.raw([
    { code: 'SENINSEGAR', name: 'Diskon 20% Awal Pekan', discount: '20%', minSpend: 50000, usage: '184 / 300', status: 'active' },
    { code: 'BARISTAFAV', name: 'Potongan Flat Rp 10.000', discount: 'Rp 10.000', minSpend: 60000, usage: '92 / 150', status: 'active' },
    { code: 'PAYDAYFEST', name: 'Cashback Gajian Member', discount: '25%', minSpend: 100000, usage: 'Expired', status: 'expired' }
  ]);

  const packages = $state.raw([
    { id: 'pkg-01', name: 'Paket Sarapan Hemat', items: 'Aren Latte + Croissant', cogs: 18000, bundlePrice: 38000, normalPrice: 46000 },
    { id: 'pkg-02', name: 'Work From Cafe Duo', items: '2x Cold Brew + 2x Brownies', cogs: 32000, bundlePrice: 75000, normalPrice: 92000 }
  ]);

  const outlets = $state.raw([
    { id: 'out-1', name: 'Cabang Kemang, Jakarta Selatan', address: 'Jl. Kemang Raya No. 14', registers: 2, todayRev: 7850000 },
    { id: 'out-2', name: 'Cabang Senopati, Jakarta Selatan', address: 'Jl. Senopati No. 88', registers: 3, todayRev: 8120000 },
    { id: 'out-3', name: 'Cabang Riau, Bandung', address: 'Jl. L.L.R.E. Martadinata No. 42', registers: 1, todayRev: 2450000 }
  ]);

  function handleLogin(e: SubmitEvent) {
    e.preventDefault();
    isSubmitting = true;
    setTimeout(() => {
      isSubmitting = false;
      currentStep = 'DASHBOARD';
    }, 350);
  }

  function formatIDR(num: number): string {
    return 'Rp ' + num.toLocaleString('id-ID');
  }
</script>

{#snippet statusBadge(status: string, text: string)}
  {#if status === 'out' || status === 'expired'}
    <span class="badge-status out">{text}</span>
  {:else if status === 'low'}
    <span class="badge-status low">{text}</span>
  {:else}
    <span class="badge-status active">{text}</span>
  {/if}
{/snippet}

<div class="app-viewport">
  {#if currentStep === 'LOGIN'}
    <!-- Split Panel Auth View (Owner & Branch Manager Portal) -->
    <div class="auth-layout">
      <!-- Left: High-Fashion / Editorial Photo Cover Panel -->
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

      <!-- Right: Pure Modern Form -->
      <main class="auth-form-wrap">
        <div class="auth-form-card">
          <div class="form-title-group">
            <h1 class="form-title">Masuk ke Dasbor</h1>
            <p class="form-sub">Akses manajemen terpusat untuk pemilik bisnis & store manager</p>
          </div>

          <!-- OAuth Provider Row (Equal weight, neutral outline) -->
          <div class="sso-group">
            <button type="button" class="sso-btn" onclick={() => currentStep = 'DASHBOARD'}>
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

          <form onsubmit={handleLogin} class="main-form">
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

  {:else if currentStep === 'DASHBOARD'}
    <!-- Backoffice Management Dashboard (Operator / Owner Scope) -->
    <div class="dash-shell">
      <!-- Left Sidebar: Graphite Tinted Surface -->
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
            onclick={() => activeNav = 'overview'}
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
            onclick={() => activeNav = 'products'}
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
            onclick={() => activeNav = 'packages'}
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
            onclick={() => activeNav = 'vouchers'}
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
            onclick={() => activeNav = 'outlets'}
          >
            <svg class="nav-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>
              <polyline points="9 22 9 12 15 12 15 22"/>
            </svg>
            <span>Manajemen Cabang</span>
          </button>
        </nav>

        <div class="sidebar-footer">
          <button type="button" class="btn-logout" onclick={() => currentStep = 'LOGIN'}>
            <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
              <polyline points="16 17 21 12 16 7"/>
              <line x1="21" y1="12" x2="9" y2="12"/>
            </svg>
            <span>Keluar Akun</span>
          </button>
        </div>
      </aside>

      <!-- Main Workspace -->
      <main class="dash-main">
        <!-- Topbar: Global Branch Switcher + Actions -->
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

        <!-- Dynamic Content Body -->
        <div class="dash-body">
          {#if activeNav === 'overview'}
            <!-- Overview & KPI Metrics Tier -->
            <div class="kpi-grid">
              <div class="metric-card primary">
                <div class="metric-head">
                  <span class="metric-label">Total Omzet Hari Ini</span>
                  <span class="tag-live">LIVE</span>
                </div>
                <div class="metric-big-num">Rp 18.420.000</div>
                <div class="metric-footnote">+14.2% dari target harian konsolidasi</div>
              </div>

              <div class="metric-card">
                <div class="metric-head">
                  <span class="metric-label">Gross Profit (Margin Kotor)</span>
                </div>
                <div class="metric-num">Rp 11.604.600</div>
                <div class="metric-footnote">Rata-rata margin kotor 63.0%</div>
              </div>

              <div class="metric-card">
                <div class="metric-head">
                  <span class="metric-label">Total Transaksi</span>
                </div>
                <div class="metric-num">524 <small>struk</small></div>
                <div class="metric-footnote">Rata-rata basket size Rp 35.150</div>
              </div>

              <div class="metric-card">
                <div class="metric-head">
                  <span class="metric-label">Alert Inventori</span>
                </div>
                <div class="metric-num text-warning">3 <small>item menipis</small></div>
                <div class="metric-footnote">Butuh restock bahan baku</div>
              </div>
            </div>

            <!-- Content Split: Recent Sales Table + Fast Promo Activity -->
            <div class="dashboard-split">
              <!-- Left: Product Master Overview Table -->
              <div class="table-container-card">
                <div class="card-header-bar">
                  <div>
                    <h3 class="card-section-title">Performa Menu Utama</h3>
                    <p class="card-section-sub">Daftar harga jual & kalkulasi snapshot HPP (COGS)</p>
                  </div>
                  <button type="button" class="btn-outline-sm" onclick={() => activeNav = 'products'}>
                    Lihat Semua Katalog
                  </button>
                </div>

                <table class="data-table">
                  <thead>
                    <tr>
                      <th>Nama Menu</th>
                      <th>Kategori</th>
                      <th class="th-right">HPP (COGS)</th>
                      <th class="th-right">Harga Jual</th>
                      <th class="th-right">Margin</th>
                      <th>Stok</th>
                    </tr>
                  </thead>
                  <tbody>
                    {#each products as p (p.id)}
                      <tr>
                        <td><strong>{p.name}</strong></td>
                        <td><span class="badge-category">{p.category}</span></td>
                        <td class="td-right num">{formatIDR(p.cogs)}</td>
                        <td class="td-right num"><strong>{formatIDR(p.price)}</strong></td>
                        <td class="td-right num text-accent">{p.margin}</td>
                        <td>
                          {@render statusBadge(p.status, p.status === 'out' ? 'Habis' : p.status === 'low' ? `Sisa ${p.stock}` : `${p.stock} porsi`)}
                        </td>
                      </tr>
                    {/each}
                  </tbody>
                </table>
              </div>

              <!-- Right: Active Vouchers Mini View -->
              <div class="side-card">
                <div class="card-header-bar">
                  <div>
                    <h3 class="card-section-title">Voucher Aktif</h3>
                    <p class="card-section-sub">Promo diskon berjalan</p>
                  </div>
                  <button type="button" class="btn-outline-sm" onclick={() => activeNav = 'vouchers'}>
                    + Buat
                  </button>
                </div>

                <div class="voucher-list-vertical">
                  {#each vouchers.slice(0, 2) as v (v.code)}
                    <div class="voucher-item-box">
                      <div class="voucher-code-badge">{v.code}</div>
                      <div class="voucher-detail">
                        <div class="v-name">{v.name}</div>
                        <div class="v-meta">Min. belanja {formatIDR(v.minSpend)} • Terpakai: {v.usage}</div>
                      </div>
                    </div>
                  {/each}
                </div>
              </div>
            </div>

          {:else if activeNav === 'products'}
            <!-- Products & COGS / HPP Management View -->
            <div class="full-content-card">
              <div class="card-header-bar">
                <div>
                  <h2 class="card-section-title">Katalog Produk & Snapshot HPP</h2>
                  <p class="card-section-sub">Atur harga retail, kontrol HPP per cabang, dan pantau margin laba per item.</p>
                </div>
                <div class="btn-group">
                  <button type="button" class="btn-primary-sm">+ Tambah Produk Baru</button>
                </div>
              </div>

              <table class="data-table">
                <thead>
                  <tr>
                    <th>Menu</th>
                    <th>Kategori</th>
                    <th class="th-right">HPP Bahan Baku</th>
                    <th class="th-right">Harga Jual Retail</th>
                    <th class="th-right">Margin Laba</th>
                    <th>Status Stok</th>
                    <th class="th-right">Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  {#each products as p (p.id)}
                    <tr>
                      <td><strong>{p.name}</strong></td>
                      <td><span class="badge-category">{p.category}</span></td>
                      <td class="td-right num">{formatIDR(p.cogs)}</td>
                      <td class="td-right num"><strong>{formatIDR(p.price)}</strong></td>
                      <td class="td-right num text-accent">{p.margin}</td>
                      <td>
                        {@render statusBadge(p.status, p.status === 'out' ? 'Habis' : p.status === 'low' ? `Sisa ${p.stock}` : `${p.stock} porsi`)}
                      </td>
                      <td class="td-right">
                        <button type="button" class="btn-action-link">Ubah Harga</button>
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            </div>

          {:else if activeNav === 'packages'}
            <!-- Bundles & Combos View -->
            <div class="full-content-card">
              <div class="card-header-bar">
                <div>
                  <h2 class="card-section-title">Paket Menu & Combo Bundle</h2>
                  <p class="card-section-sub">Gabungkan beberapa item menjadi paket promo dengan harga spesial.</p>
                </div>
                <button type="button" class="btn-primary-sm">+ Buat Paket Baru</button>
              </div>

              <table class="data-table">
                <thead>
                  <tr>
                    <th>Nama Paket</th>
                    <th>Isi Paket (Menu)</th>
                    <th class="th-right">Total HPP</th>
                    <th class="th-right">Harga Normal</th>
                    <th class="th-right">Harga Paket</th>
                    <th class="th-right">Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  {#each packages as pkg (pkg.id)}
                    <tr>
                      <td><strong>{pkg.name}</strong></td>
                      <td><span class="badge-category">{pkg.items}</span></td>
                      <td class="td-right num">{formatIDR(pkg.cogs)}</td>
                      <td class="td-right num strike">{formatIDR(pkg.normalPrice)}</td>
                      <td class="td-right num text-accent"><strong>{formatIDR(pkg.bundlePrice)}</strong></td>
                      <td class="td-right">
                        <button type="button" class="btn-action-link">Edit Paket</button>
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            </div>

          {:else if activeNav === 'vouchers'}
            <!-- Vouchers & Promo View -->
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
                        {@render statusBadge(v.status, v.status === 'active' ? 'Aktif' : 'Berakhir')}
                      </td>
                      <td class="td-right">
                        <button type="button" class="btn-action-link">Detail</button>
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            </div>

          {:else if activeNav === 'outlets'}
            <!-- Outlets List View -->
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

  /* ================= Auth Layout (Split Panel) ================= */
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

  /* Left Side: Editorial Image Panel */
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
    transition: all 200ms ease;
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

  /* Right Side: Form Container */
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

  /* ================= Backoffice Dashboard Layout ================= */
  .dash-shell {
    display: flex;
    width: 100%;
    min-height: 100vh;
  }

  /* Sidebar: 240px Graphite Tinted Surface */
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

  /* Main Workspace Area */
  .dash-main {
    flex: 1;
    display: flex;
    flex-direction: column;
    overflow-x: hidden;
  }

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

  .dash-body {
    padding: 1.75rem;
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
  }

  /* KPI Cards Grid */
  .kpi-grid {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 1rem;
  }

  @media (max-width: 1200px) {
    .kpi-grid {
      grid-template-columns: repeat(2, 1fr);
    }
  }

  .metric-card {
    background-color: #ffffff;
    border: 1px solid var(--gray-2);
    border-radius: var(--radius-md);
    padding: 1.25rem;
    box-shadow: var(--shadow-card);
  }

  .metric-card.primary {
    background-color: var(--accent-tint);
    border-color: var(--accent);
  }

  .metric-head {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 0.5rem;
  }

  .metric-label {
    font-size: 0.75rem;
    color: var(--gray-6);
    font-weight: 600;
  }

  .tag-live {
    font-size: 0.625rem;
    font-weight: 700;
    color: var(--accent);
    background: white;
    padding: 0.125rem 0.375rem;
    border-radius: var(--radius-sm);
  }

  .metric-big-num {
    font-size: 1.625rem;
    font-weight: 700;
    font-family: var(--font-mono);
    color: var(--accent);
    letter-spacing: -0.03em;
    margin-bottom: 0.25rem;
  }

  .metric-num {
    font-size: 1.375rem;
    font-weight: 700;
    font-family: var(--font-mono);
    color: var(--gray-9);
    letter-spacing: -0.02em;
    margin-bottom: 0.25rem;
  }

  .metric-footnote {
    font-size: 0.75rem;
    color: var(--gray-6);
  }

  /* Split Layout */
  .dashboard-split {
    display: grid;
    grid-template-columns: 1fr 340px;
    gap: 1.25rem;
  }

  @media (max-width: 1100px) {
    .dashboard-split {
      grid-template-columns: 1fr;
    }
  }

  .table-container-card, .side-card, .full-content-card {
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

  /* Data Table Styling */
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

  .strike {
    text-decoration: line-through;
    color: var(--gray-6);
  }

  .text-accent {
    color: var(--accent);
    font-weight: 600;
  }

  .text-warning {
    color: var(--warning);
  }

  .text-success {
    color: var(--success);
  }

  .badge-category {
    background-color: var(--gray-1);
    color: var(--gray-7);
    padding: 0.2rem 0.5rem;
    border-radius: var(--radius-sm);
    font-size: 0.6875rem;
    font-weight: 500;
  }

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

  .btn-outline-sm {
    background: none;
    border: 1px solid var(--gray-3);
    padding: 0.375rem 0.75rem;
    border-radius: var(--radius-sm);
    font-size: 0.75rem;
    font-weight: 500;
    color: var(--gray-7);
    cursor: pointer;
  }

  .btn-outline-sm:hover {
    background-color: var(--gray-1);
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

  .voucher-list-vertical {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  .voucher-item-box {
    border: 1px solid var(--gray-2);
    border-radius: var(--radius-sm);
    padding: 0.75rem;
    display: flex;
    flex-direction: column;
    gap: 0.375rem;
  }

  .voucher-code-badge, .voucher-code {
    font-family: var(--font-mono);
    font-weight: 700;
    font-size: 0.75rem;
    color: var(--accent);
    background-color: var(--accent-tint);
    padding: 0.15rem 0.45rem;
    border-radius: var(--radius-sm);
    display: inline-block;
    width: fit-content;
  }

  .v-name {
    font-size: 0.8125rem;
    font-weight: 600;
    color: var(--gray-9);
  }

  .v-meta {
    font-size: 0.6875rem;
    color: var(--gray-6);
  }

  /* Outlets Grid */
  .outlets-cards-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 1rem;
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
</style>
