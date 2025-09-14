<script>
  let props = $props()
  let address = props.address
  let label = props.label

  let isLoading = $state(true)
  let status = $state(false)
  let player_number = $state(0)
  let max_player_number = $state(0)

  function getStatus() {
    isLoading = true
    fetch(`https://api.mcstatus.io/v2/status/java/${address}`)
      .then(res => res.json())
      .then(info => {
        status = info.online
        if (status) {
          player_number = info.players.online
          max_player_number = info.players.max
        }
      })
      .catch(err => console.log(err))
      .finally(() => {isLoading = false})
  }

  getStatus()
</script>

<div>
  <div class="card">
    {#if !isLoading}
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
    {:else}
      <div>Loading...</div>
    {/if}
  </div>
</div>

