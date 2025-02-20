<script lang="ts">
  import { authStore } from "../stores/auth";
  import { navigate } from "svelte-routing";

  let user: any;

  const unsubscribe = authStore.subscribe((state) => {
    user = state.user;
  });

  function handleLogout() {
    authStore.logout();
    navigate("/login");
  }
</script>

<nav class="bg-gray-800 shadow-md">
  <div class="container mx-auto px-4">
    <div class="flex h-16 items-center justify-between">
      <div class="flex items-center">
        <span class="text-xl font-bold text-cyan-500">Music Tag Editor</span>
      </div>

      <div class="ml-4 flex items-center md:ml-6">
        <div class="relative ml-3">
          <div class="flex items-center">
            <span class="mr-4 text-sm font-medium text-gray-300">
              {user?.username || "User"}
            </span>

            <button
              type="button"
              on:click={handleLogout}
              class="rounded-md bg-gray-700 px-3 py-2 text-sm font-medium text-gray-200 hover:bg-gray-600 hover:text-white"
            >
              Sign out
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</nav>
