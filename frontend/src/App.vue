<script setup>
  import { ref, onMounted, computed } from 'vue';
  
  const stats = ref({})
  const risks = ref([])
  const lastUpdated = ref('')

  async function fetchRiskData() {
    const response = await fetch('https://riskpulse-production-efb9.up.railway.app/risk')
    const data = await response.json();
    lastUpdated.value = new Date().toLocaleTimeString()
    stats.value = data.stats
    risks.value = data.risks
  }

  onMounted(() => {
    fetchRiskData()
  })

  function getRiskClass(risk) {
    if (risk >= 80) return 'high';
    if (risk >= 50) return 'medium';
    return 'low';
  }

  const riskStats = computed(() => {
    const total = risks.value.length;
    const low = risks.value.filter(r => r.risk < 50).length;
    const medium = risks.value.filter(r => r.risk >= 50 && r.risk < 80).length;
    const high = risks.value.filter(r => r.risk >= 80).length; 

    return { low, medium, high, total }
  })

  const vpnStats = computed(() => {
    const total = risks.value.length;
    const vpn = risks.value.filter(r => r.vpn).length;
    const noVpn = total - vpn;

    return { vpn, noVpn, total }
  })

  const navItems = [
    'Dashboard',
    'Team',
    'Analytics',
    'Completed',
    'Settings',
  ]

</script>

<template>

  <div class="app-shell">
    <aside class="sidebar">
      <div class="sidebar-brand">
        <img src="/logo.png" alt="RiskPulse Logo" class="logo">
      </div>
      <nav class="sidebar-nav" aria-label="Main navigation">
        <a
          v-for="item in navItems"
          :key="item"
          href="#"
          class="nav-item"
          @click.prevent
        >{{ item }}</a>
      </nav>
    </aside>

    <div class="dashboard">
    <div class="dashboard-meta">
      <time class="last-updated">
        Last Updated: {{ lastUpdated }}
      </time>
      <div class="api-status" aria-label="API connection status">
        <span class="dot" aria-hidden="true"></span>
        API Connected
      </div>
    </div>
    <h1 class="title">Risk Dashboard</h1>
  
    <!-- CARDS -->
     <div class="cards">
      <div class="card">
        <h2>Threat Requests</h2>
        <p>{{ stats.threat_requests }}</p>
      </div>

      <div class="card">
        <h2>VPN Detections</h2>
        <p>{{ stats.vpn_detections }}</p>
      </div>

      <div class="card">
        <h2>Fraud Attempts</h2>
        <p>{{ stats.fraud_attempts }}</p>
      </div>

      <div class="card">
        <h2>Risk Alerts</h2>
        <p>{{ stats.risk_alerts }}</p>
      </div>
     </div>

    <!-- CHART ROW-->
    <div class="chart-row">
      <div class="chart-box">
        <h3>Risk Distribution</h3>
        <div class="bar-label">High Risk</div>
        <div class="demo-bar high"
          :style="{ width: (riskStats.high / riskStats.total * 100) + '%' }">
        </div>
        <div class="bar-label">Medium Risk</div>
        <div class="demo-bar medium"
          :style="{ width: (riskStats.medium / riskStats.total * 100) + '%' }">
        </div>
        <div class="bar-label">Low Risk</div>
        <div class="demo-bar low"
          :style="{ width: (riskStats.low / riskStats.total * 100) + '%' }">
        </div>
        <div class="risk-labels">
          <span>High: {{ riskStats.high }}</span>
          <span>Medium: {{ riskStats.medium }}</span>
          <span>Low: {{ riskStats.low }}</span>
        </div>
      </div>
      <div class="chart-box">
        <h3>VPN Activity</h3>
        <div class="vpn-bar">
          <div class="vpn-yes" :style="{ width: vpnStats.vpn / vpnStats.total * 100 + '%' }"></div>
          <div class="vpn-no" :style="{ width: vpnStats.noVpn / vpnStats.total * 100 + '%' }"></div>
        </div>

        <div class="vpn-labels">
          <span>VPN: {{ vpnStats.vpn }}</span>
          <span>No VPN: {{ vpnStats.noVpn }}</span>
        </div>
      </div>
    </div>
    <!-- TABLE -->
    <div class="table-wrapper">
      <table>
        <thead>
          <tr>
            <th>IP</th>
            <th>Country</th>
            <th>Risk</th>
            <th>VPN</th>
            <th>Status</th>
          </tr>
        </thead>

        <tbody>
          <tr v-for="item in risks" :key="item.ip">
            <td>{{ item.ip }}</td>
            <td>{{ item.country }}</td>
            <td><span :class="['risk', getRiskClass(item.risk)]">{{ item.risk }}</span></td>
            <td>
              <span :class="['vpn', item.vpn ? 'yes-vpn' : 'no-vpn']">
                {{ item.vpn ? 'YES' : 'NO' }}
              </span>
            </td>
            <td><span :class="['status', item.status]">{{ item.status }}</span></td>
          </tr>
        </tbody>
      </table>
    </div>
    </div>
  </div>
</template>

<style>
/* FONTS */
* {
  box-sizing: border-box;
}

body {
  font-family: "Inter", sans-serif;
  background: #f6f8fc;
  margin: 0;
}

h1, h2, h3 {
  font-weight: 600;
  letter-spacing: -0.02em;
}

p, td, span {
  font-weight: 400;
}
/* APP SHELL & SIDEBAR */
.app-shell {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  flex-shrink: 0;
  width: 220px;
  min-height: 100vh;
  padding: 20px 16px;
  font-family: "Inter", sans-serif;
  background: linear-gradient(to bottom right, #ffffff 65%, #efeff3);
  border-right: 1px solid #e7ebf2;
}

.sidebar-brand {
  margin-bottom: 24px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e7ebf2;
}

.sidebar .logo {
  width: 100%;
  max-width: 180px;
  display: block;
}

.sidebar-nav {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.nav-item {
  display: block;
  padding: 10px 14px;
  border-radius: 8px;
  font-family: "Inter", sans-serif;
  font-size: 14px;
  font-weight: 500;
  color: #374151;
  text-decoration: none;
  transition: all 0.2s ease-in-out;
}

.nav-item:hover {
  background: rgba(17, 24, 39, 0.06);
  color: #111827;
  transform: translateX(3px);
}

@media (max-width: 768px) {
  .app-shell {
    flex-direction: column;
    min-height: 100vh;
  }

  .sidebar {
    width: 100%;
    min-height: auto;
    border-right: none;
    border-bottom: 1px solid #e7ebf2;
    padding: 14px 20px;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }

  .sidebar-brand {
    margin-bottom: 0;
    padding-bottom: 0;
    border-bottom: none;
  }

  .sidebar .logo {
    max-width: 140px;
  }

  .sidebar-nav {
    flex-direction: row;
    flex-wrap: wrap;
    width: 100%;
    gap: 6px;
  }

  .nav-item {
    padding: 8px 12px;
  }

  .nav-item:hover {
    transform: translateY(-2px);
  }
}

.dashboard-meta {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 16px;
  margin-bottom: 12px;
}

.last-updated {
  font-size: 12px;
  color: #ffffff;
}

.api-status {
  display: flex;
  align-items: center;
  font-size: 13px;
  font-weight: 500;
  color: #ffffff;
}

.dot {
  display: inline-block;
  width: 8px;
  height: 8px;
  background: #22c55e;
  border-radius: 50%;
  margin-right: 8px;
}

/* DASHBOARD */
.dashboard {
  flex: 1;
  min-width: 0;
  padding: 25px 32px;
  background: linear-gradient(to right, #1d1d1d 10%, #131313 50%, #1d1d1d 90%);
  min-height: 100vh;
}

@keyframes fade {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.title,
.card,
.chart-box,
.table-wrapper {
  animation: fade 0.55s ease-out both;
}

.title {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 8px;
  color: #ffffff;
  animation-delay: 0.05s;
}

.cards .card:nth-child(1) { animation-delay: 0.12s; }
.cards .card:nth-child(2) { animation-delay: 0.2s; }
.cards .card:nth-child(3) { animation-delay: 0.28s; }
.cards .card:nth-child(4) { animation-delay: 0.36s; }

.chart-row .chart-box:nth-child(1) { animation-delay: 0.44s; }
.chart-row .chart-box:nth-child(2) { animation-delay: 0.52s; }

.table-wrapper {
  animation-delay: 0.6s;
}

@media (prefers-reduced-motion: reduce) {
  .title,
  .card,
  .chart-box,
  .table-wrapper {
    animation: none;
    opacity: 1;
    transform: none;
  }
}

/* CARDS */
.cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 16px;
  margin-bottom: 18px; 
}

@media (max-width: 768px) {
  .cards {
    grid-template-columns: 1fr;
  }
}

.card {
  background: linear-gradient(to bottom right, #ffffff 65%, #efeff3);
  padding: 18px;
  border-radius: 12px;
  border: 1px solid #e7ebf2;
  transition: all 0.2s ease-in-out;
  backdrop-filter: blur(8px);
}

.card:hover {
  transform: translateY(-3px);
  cursor: pointer;
}

.card h2 {
  font-size: 14px;
  margin-bottom: 8px;
  color: #6b7280;
  letter-spacing: 0.3px;
  text-transform: uppercase;
}

.card p {
  font-size: 28px;
  font-weight: 800;
  color: #111827;
  margin-top: 6px;
}

/* CHART ROW */
.chart-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
  margin-bottom: 18px;
}

@media (max-width: 768px) {
  .chart-row {
    grid-template-columns: 1fr;
  }
}

.chart-box {
  background: white;
  border: 1px solid #eef0f4;
  padding: 16px;
  border-radius: 12px;
  background: linear-gradient(to bottom right, #ffffff 65%, #efeff3);

}

.bar-label {
  font-size: 12px;
}

.demo-bar {
  height: 10px;
  margin: 8px 0;
  border-radius: 6px;
  transition: width 0.6s ease;
}

.demo-bar.high { background: #b42318; width: 80%; }
.demo-bar.medium { background: #d87b34; width: 55%; }
.demo-bar.low { background: #1f7a3f; width: 30%; }

.risk-labels {
  display: flex;
  justify-content: start;
  font-size: 12px;
  font-weight: 500;
  margin-top: 8px;
  gap: 10px;
  color: black;
}

.circle-mock {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: #f1f5f9;
  margin-top: 20px;;
}

.vpn-bar {
  display: flex;
  height: 14px;
  border-radius: 8px;
  overflow: hidden;
  background: #e5e7eb;
  margin-top: 12px;
}

.vpn-yes {
  background: #b42318;
}

.vpn-no {
  background: #1f7a3f;
}

.vpn-labels {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  font-weight: 500;
  margin-top: 8px;
  color: black;
}

/* TABLE */
.table-wrapper {
  background: white;
  border: 1px solid #eef0f4;
  padding: 16px;
  border-radius: 12px;
  max-height: 600px;
  overflow-y: auto;
  background: linear-gradient(to bottom right, #ffffff 65%, #efeff3);

}

table {
  width: 100%;
  border-collapse: collapse;
  min-width: 500px;
}

th {
  text-align: left;
  font-size: 12px;
  color: #6b7280;
  padding-bottom: 10px;
}


td {
  padding: 12px 0;
  border-top: 1px solid #f1f1f1;
  font-size: 14px;
  color: #111827;
}

td:nth-child(3) {
  font-variant-numeric: tabular-nums;
}

tbody tr:hover {
  background: #f9fafb;
  cursor: pointer;
}

/* STATUS */
.status {
  padding: 4px 10px;
  border-radius: 25px;
  font-size: 12px;
  font-weight: 600;
  display: inline-block;
}

.clean {
  background: #e7f8ee;
  color: #1f7a3f;
}

.flagged {
  background: #fde8e8;
  color: #b42318;
}

.review {
  background: #fff7e6;
  color: #b45309;
}

/* VPN */
.vpn {
  font-weight: 600;
}

.yes-vpn {
  color: #b42318;
}

.no-vpn {
  color: #1f7a3f;
}

/* RISK */
.risk {
  font-weight: 700;
}

.low {
  color: #1f7a3f;
}

.medium {
  color: #d87b34;
}

.high {
  color: #b42318;
}

</style>