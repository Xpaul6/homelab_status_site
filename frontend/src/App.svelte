<script>
  import MCStatusBar from "./components/MCStatusBar.svelte";
  import SystemMonitor from "./components/SystemMonitor.svelte";

  const mcservers = [
    { address: "carapuziki.ddns.net:25000", label: "CARAPUZIKI" },
    { address: "carapuziki.ddns.net:25565", label: "YapSMP" },
  ];

  let server_status = $state("Loading...");

  async function getStatus() {
    try {
      const response = await fetch("/api/server-status");
      if (!response.ok) throw new Error("server is offline");

      const data = await response.text();
      server_status = data;
    } catch (error) {
      console.log(error);
      server_status = "Server is down :(";
    }
  }

  getStatus();
</script>

<svelte:head>
  <title>who the hell is giganig...</title>
  <link rel="icon" type="image" href="/gnome.png" />
</svelte:head>

<main>
  <div class="flex flex-col justify-center items-center">
    <h1 class="mt-20">{server_status}</h1>

    <h2 class="m-3 mt-20">System monitor:</h2>
    <SystemMonitor />

    <h2 class="m-3">Services status:</h2>
    <div class="flex flex-row flex-wrap w-3/4 max-w-[700px] justify-center">
      {#each mcservers as server}
        <MCStatusBar address={server.address} label={server.label} />
      {/each}
    </div>
  </div>
</main>
