<template>
  <div class="dashboard">
    <h1 class="title">Risk Dashboard</h1>
  
    <!-- CARDS -->
     <div class="cards">
      <div class="card">
        <h2>Treat Requests</h2>
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
            <td>82</td>
            <td>true</td>
            <td>flagged</td>
          </tr>

          <tr>
            <td>10.0.0.5</td>
            <td>UK</td>
            <td>35</td>
            <td>false</td>
            <td>clean</td>
          </tr>

          <tr>
            <td>172.16.0.9</td>
            <td>DE</td>
            <td>67</td>
            <td>true</td>
            <td>review</td>
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
  background: #f4f6f8;
  min-height: 100vh;
}

.title  {
  margin-bottom: 20px;
  font-size: 24px;
  font-weight: 600;
  margin-bottom: 24px;
  color: #1f2937;
}

/* CARDS */
.cards {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 16px;
  margin-bottom: 28px;  
}

.card {
  background: white;
  padding: 18px;
  border-radius: 12px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.06);
  transition: transform 0.15s ease;
}

.card:hover {
  transform: translateY(-2px);
  cursor: pointer;
}

.card h2 {
  font-size: 14px;
  margin-bottom: 8px;
  color: #6b7280;
}

.card p {
  font-size: 26px;
  font-weight: 700;
  color: #111827;
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

</style>