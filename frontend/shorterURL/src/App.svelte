<script lang="ts">
  let originalURL = "";
  let shortURL = "";

  function shortenHandler() {
    fetch("http://localhost:8080/shorten", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        url: originalURL
      })
    }).then(res => res.json()).then(data => {
      console.log(data);
      shortURL = "http://localhost:8080/"+data.url;
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
