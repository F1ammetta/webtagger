<script>
  import { navigate } from "svelte-routing";
</script>

<div
  class="flex h-screen flex-col items-center justify-center bg-gray-900 px-4 text-center"
>
  <h1 class="mb-4 text-9xl font-bold text-cyan-500">404</h1>
  <h2 class="mb-8 text-2xl font-medium text-white">Page Not Found</h2>
  <p class="mb-8 max-w-md text-gray-400">
    The page you are looking for might have been removed, had its name changed,
    or is temporarily unavailable.
  </p>
  <button
    on:click={() => navigate("/")}
    class="rounded-md bg-cyan-600 px-5 py-3 font-medium text-white hover:bg-cyan-700 focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:ring-offset-2 focus:ring-offset-gray-900"
  >
    Back to Home
  </button>
</div>
