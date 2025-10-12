<script>
  let default_monitor_data = {
    cpu_temp: 0,
    cpu_load_percent: 0,
    ram_percent: 0,
  };

  let monitor_data = $state(default_monitor_data);

  async function getMonitorData() {
    try {
      const response = await fetch("/api/sysinfo");
      if (!response.ok) throw new Error("unable to fetch system data");

      const data = await response.json();
      monitor_data = data;
    } catch (error) {
      console.log(error);
    }
  }

  getMonitorData();
  setInterval(() => getMonitorData(), 5000);
</script>

<div class="card w-2/4 max-w-[200px]">
  {#await getMonitorData()}
    <div>----------</div>
    <div>Loading...</div>
    <div>----------</div>
  {:then}
    <div>CPU temp: {monitor_data.cpu_temp}°C</div>
    <div>CPU load: {monitor_data.cpu_load_percent.toFixed(1)}%</div>
    <div>RAM load: {monitor_data.ram_percent.toFixed(1)}%</div>
  {/await}
</div>
