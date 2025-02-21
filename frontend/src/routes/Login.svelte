<script lang="ts">
  import { navigate } from "svelte-routing";
  import { authStore } from "../stores/auth";
  import { onMount } from "svelte";

  let username = "";
  let password = "";
  let error = "";
  let isLoading = false;

  async function handleSubmit() {
    error = "";
    isLoading = true;

    const result = await authStore.login(username, password);
    isLoading = false;

    if (result.success) {
      navigate("/");
    } else {
      error = result.message || "Login failed";
    }
  }
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-900 px-4">
  <div class="w-full max-w-md space-y-8 rounded-lg bg-gray-800 p-8 shadow-lg">
    <div>
      <h1 class="text-center text-3xl font-bold text-cyan-500">Yggdrasil</h1>
      <p class="mt-2 text-center text-gray-400">
        Sign in to manage your music collection
      </p>
    </div>

    <form class="mt-8 space-y-6" on:submit|preventDefault={handleSubmit}>
      {#if error}
        <div class="rounded-md bg-red-500 bg-opacity-10 p-3 text-red-400">
          {error}
        </div>
      {/if}

      <div class="space-y-4 rounded-md shadow-sm">
        <div>
          <label for="username" class="sr-only">Username</label>
          <input
            id="username"
            name="username"
            type="text"
            required
            bind:value={username}
            class="relative block w-full appearance-none rounded-md border border-gray-700 bg-gray-700 px-3 py-2 text-gray-200 placeholder-gray-500 focus:border-cyan-500 focus:outline-none focus:ring-cyan-500"
            placeholder="Username"
          />
        </div>
        <div>
          <label for="password" class="sr-only">Password</label>
          <input
            id="password"
            name="password"
            type="password"
            required
            bind:value={password}
            class="relative block w-full appearance-none rounded-md border border-gray-700 bg-gray-700 px-3 py-2 text-gray-200 placeholder-gray-500 focus:border-cyan-500 focus:outline-none focus:ring-cyan-500"
            placeholder="Password"
          />
        </div>
      </div>

      <div>
        <button
          type="submit"
          disabled={isLoading}
          class="group relative flex w-full justify-center rounded-md border border-transparent bg-cyan-600 px-4 py-2 text-sm font-medium text-white hover:bg-cyan-700 focus:outline-none focus:ring-2 focus:ring-cyan-500 focus:ring-offset-2 disabled:opacity-50"
        >
          {#if isLoading}
            <span class="absolute inset-y-0 left-0 flex items-center pl-3">
              <svg
                class="h-5 w-5 animate-spin text-cyan-300"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
              >
                <circle
                  class="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  stroke-width="4"
                ></circle>
                <path
                  class="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                ></path>
              </svg>
            </span>
            Processing...
          {:else}
            Sign in
          {/if}
        </button>
      </div>
    </form>

    <div class="text-center text-sm text-gray-400">
      <p>Demo credentials: admin / password</p>
    </div>
  </div>
</div>
