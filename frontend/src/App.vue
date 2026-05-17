<template>
  <div class="dashboard">
    <h1 class="title">Risk Dashboard</h1>
  
    <!-- CARDS -->
     <div class="cards">
      <div class="card">
        <h2>Threat Requests</h2>
        <p>128</p>
      </div>

      <div class="card">
        <h2>VPN Detections</h2>
        <p>42</p>
      </div>

      <div class="card">
        <h2>Fraud Attempts</h2>
        <p>19</p>
      </div>

      <div class="card">
        <h2>Risk Alerts</h2>
        <p>7</p>
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
          <tr>
            <td>192.168.1.0</td>
            <td>US</td>
            <td><span class="risk high">82</span></td>
            <td>true</td>
            <td><span class="status flagged">flagged</span></td>
          </tr>

          <tr>
            <td>10.0.0.5</td>
            <td>UK</td>
            <td><span class="risk low">35</span></td>
            <td>false</td>
            <td><span class="status clean">clean</span></td>
          </tr>

          <tr>
            <td>172.16.0.9</td>
            <td>DE</td>
            <td><span class="risk medium">67</span></td>
            <td>true</td>
            <td><span class="status review">review</span></td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style>
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