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