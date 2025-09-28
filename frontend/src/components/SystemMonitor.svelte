<script>
  let default_monitor_data = {
    cpu_temp: 0,
    cpu_load_percent: 0,
    ram_percent: 0,
  };

  let isLoading = $state(true);
  let monitor_data = $state(default_monitor_data);

  async function getMonitorData() {
    isLoading = true;
    try {
      const request = await fetch("/api/sysinfo");
      const data = await request.json();
      if (request.status != 200) {
        throw Error(data.error);
      }
      monitor_data = data;
    } catch (error) {
      console.log(error);
      monitor_data = default_monitor_data
    } finally {
      isLoading = false;
    }
  }

  getMonitorData();
</script>

<div class="card w-2/4 max-w-[200px]">
  {#if isLoading}
    <div>Loading...</div>
  {:else}
    <div>CPU temp: {monitor_data.cpu_temp}°C</div>
    <div>CPU load: {monitor_data.cpu_load_percent.toFixed(1)}%</div>
    <div>RAM load: {monitor_data.ram_percent.toFixed(1)}%</div>
  {/if}
</div>
