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
      <h1 class="text-center text-3xl font-bold text-cyan-500">
        Yggdrasil
        <!--<svg
        viewBox="0 0 64 64"
        xmlns="http://www.w3.org/2000/svg"
        xmlns:xlink="http://www.w3.org/1999/xlink"
        aria-hidden="true"
        role="img"
        class="iconify iconify--emojione h-60 fill-orange-400"
        preserveAspectRatio="xMidYMid meet"
        style="--darkreader-inline-fill: var(--darkreader-background-000000, #000000);"
        data-darkreader-inline-fill=""
        ><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g
          id="SVGRepo_tracerCarrier"
          stroke-linecap="round"
          stroke-linejoin="round"
        ></g><g id="SVGRepo_iconCarrier">
          <g
            style="--darkreader-inline-fill: var(--darkreader-background-5d6d74, #4f5559);"
            data-darkreader-inline-fill=""
          >
            <path
              class="fill-cyan-700"
              d="M38.5 42.6c-4.6-1.5-3.3-2.6 1.1-4.3c2.1-.8.8-2.5.9-3.8c0-.6 2.5-.5 2.3-2.9c-.1-1.7-3.8-4.1-4.8-5.1c-.6-.6 1.2-2.2-.1-3.6c-1.7-1.9-2-5.3-3-7.2c0 0 .8-1.2.2-1.9C29.9 8 10.6 8.6 5.6 17.2c-5.5 9.7-5.6 23 5.9 30.2c5.1 3.2-1.4 14.6-1.4 14.6h20.3c0-1.9-2.3-8.9 1.7-8.6c3.4.3 7.7.1 7.3-3.8c-.1-1.2-.2-2.2.6-3.2c.8-.9 2-2.7-1.5-3.8"
            >
            </path> <path d="M43.1 40.8L62 43.3v-5z"> </path>
            <path d="M58.5 57.1l2-4.3l-17.4-9.4z"> </path>
            <path d="M60.5 28.8l-2-4.3l-15.4 13.6z"> </path>
          </g>
        </g></svg
      >-->
      </h1>
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
