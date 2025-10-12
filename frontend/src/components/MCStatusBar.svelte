<script>
  let props = $props();
  let address = props.address;
  let label = props.label;

  let status = $state(false);
  let player_number = $state(0);
  let max_player_number = $state(0);

  async function getStatus() {
    try {
      const response = await fetch(`/api/mcstatus/${address}`);
      if (!response.ok) throw new Error(`unable to fetch ${address} status`);

      const info = await response.json();
      status = info.online;
      if (status) {
        player_number = info.players.online;
        max_player_number = info.players.max;
      }
    } catch (error) {
      console.log(error);
    }
  }

  getStatus();
</script>

<div>
  <div class="card">
    {#await getStatus()}
      <div>Loading...</div>
    {:then}
      <div>
        {label}:
        {#if status}
          <span class="text-green-200">online</span>
        {:else}
          offline
        {/if}
      </div>
      {#if status}
        <div>{player_number}/{max_player_number} players</div>
      {/if}
    {/await}
  </div>
</div>
