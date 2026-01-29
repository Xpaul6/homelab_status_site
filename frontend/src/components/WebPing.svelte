<script>
  let { address, label } = $props();

  let status = $state(false);

  async function getSiteStatus() {
    try {
      const response = await fetch(`/api/websitestatus?address=${address}`);
      if (!response.ok) throw new Error(`unable to fetch ${address} status`);
      status = true;
    } catch (error) {
      console.log(error);
    }
  }

  getSiteStatus();
</script>

<div>
  <div class="card">
    {#await getSiteStatus()}
      <div>Loading...</div> 
    {:then} 
      {label}:
      {#if status}
        <span class="text-green-200">online</span>
      {:else}
        <span class="text-red-200">offline</span>
      {/if}
    {/await} 
  </div>
</div>
