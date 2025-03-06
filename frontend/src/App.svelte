<script lang="ts">
  import { onMount } from "svelte";
  import { Router, Route } from "svelte-routing";
  import { authStore } from "./stores/auth";
  import Login from "./routes/Login.svelte";
  import Dashboard from "./routes/Dashboard.svelte";
  import NotFound from "./routes/NotFound.svelte";
  // import ProtectedRoute from "./components/ProtectedRoute.svelte";

  export let url = "";
  let isLoading = true;

  onMount(async () => {
    // Check if user is already logged in (token in localStorage)
    await authStore.checkAuth();
    isLoading = false;
  });
</script>

{#if isLoading}
  <div class="flex h-screen w-screen items-center justify-center bg-gray-900">
    <div
      class="h-12 w-12 animate-spin rounded-full border-4 border-cyan-500 border-t-transparent"
    ></div>
  </div>
{:else}
  <Router {url}>
    <div class="min-h-screen bg-gray-900 text-gray-100">
      <Route path="/login" component={Login} />
      <Route path="/" component={Dashboard} />
      <Route path="*" component={NotFound} />
    </div>
  </Router>
{/if}
