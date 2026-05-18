<script setup>
  import { ref, onMounted } from 'vue';
  
  const stats = ref({})
  const risks = ref([])

  async function fetchRiskData() {
    const response = await fetch('http://localhost:8080/risk')
    const data = await response.json();

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
</script>

<template>

  <header class="topbar">
    <div class="brand">RiskPulse</div>
    <div class="status">
      <span class="dot"></span>
        API Connected
    </div>
  </header>

  <div class="dashboard">
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
        <div class="demo-bar high"></div>
        <div class="demo-bar medium"></div>
        <div class="demo-bar low"></div>
      </div>
      <div class="chart-box">
        <h3>VPN Activity</h3>
        <div class="circle-mock"></div>
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
</template>

<style>
.topbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 32px;
  background: white;
  border-bottom: 1px solid #eee;
  margin-bottom: 20px;
}

.brand {
  font-weight: 700;
  font-size: 18px;
}

.dot {
  width: 8px;
  height: 8px;
  background: #22c55e;
  border-radius: 50%;
  display: inline-block;
  margin-right: 8px;
}

.dashboard {
  padding: 32px;
  font-family: Arial, sans-serif;
  background: linear-gradient(to bottom, #f5f7fb, #ffffff);
  min-height: 100vh;
}

.title  {
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 24px;
  color: #1f2937;
}

/* CARDS */
.cards {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 16px;
  margin-bottom: 28px;  
}

.card {
  background: white;
  padding: 18px;
  border-radius: 12px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.06);
  transition: all 0.2s ease;
}

.card:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 18px rgba(0, 0, 0, 0.08);
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
  margin-bottom: 24px;
}

.chart-box {
  background: white;
  padding: 16px;
  border-radius: 12px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.06);
}

.demo-bar {
  height: 10px;
  margin: 8px 0;
  border-radius: 6px;
}

.demo-bar.high { background: #b42318; width: 80%; }
.demo-bar.medium { background: #d87b34; width: 55%; }
.demo-bar.low { background: #1f7a3f; width: 30%; }

.circle-mock {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: #f1f5f9;
  margin-top: 20px;;
}

/* TABLE */
.table-wrapper {
  background: white;
  padding: 16px;
  border-radius: 12px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.06);
}

table {
  width: 100%;
  border-collapse: collapse;
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