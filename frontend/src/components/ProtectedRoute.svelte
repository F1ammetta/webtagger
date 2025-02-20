<script lang="ts">
  import { onMount } from "svelte";
  import { navigate } from "svelte-routing";
  import { authStore } from "../stores/auth";

  let isAuthenticated: boolean;

  const unsubscribe = authStore.subscribe((state) => {
    isAuthenticated = state.isAuthenticated;
  });

  onMount(() => {
    if (!isAuthenticated) {
      navigate("/login", { replace: true });
    }
    return unsubscribe;
  });
</script>

{#if isAuthenticated}
  <slot></slot>
{/if}
