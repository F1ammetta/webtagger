<script lang="ts">
  import { onMount } from "svelte";
  import { navigate } from "svelte-routing";
  import { authStore } from "../stores/auth";
  import Dashboard from "../routes/Dashboard.svelte";

  let isAuthenticated: boolean;

  const unsubscribe = authStore.subscribe((state) => {
    isAuthenticated = state.isAuthenticated;
  });

  onMount(() => {
    if (!isAuthenticated) {
      navigate("/login", { replace: true });
      window.location.reload();
    }
    return unsubscribe;
  });
</script>

{#if isAuthenticated}
  <Dashboard />
{/if}
