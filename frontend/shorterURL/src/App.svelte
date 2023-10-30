<script lang="ts">
  let originalURL = "";
  let shortURL = "";

  function shortenHandler() {
    fetch("https://railway.app/project/f22076ea-5cc1-4c6e-97fd-d8b6c3ce55d9/", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        url: originalURL
      })
    }).then(res => res.json()).then(data => {
      console.log(data);
      shortURL = "https://railway.app/project/f22076ea-5cc1-4c6e-97fd-d8b6c3ce55d9/"+data.url;
    });
  }
</script>

<main>
  <label for="url">URL</label>
  <input id="url" type="text" placeholder="https://example.com" bind:value={originalURL}/>
  <button on:click={
    shortenHandler
  }>Shorten</button>

  {#if shortURL !== ""}
    <p>
      {shortURL}
      <button on:click={
        () => {
          window.location.href = shortURL;
        }
      }>Go to this link</button>
      <button on:click={
        () => {
          navigator.clipboard.writeText(shortURL);
        }
      }>Copy</button>
    </p>
    
  {/if}
</main>

<style>
</style>
