<script lang="ts">
  import { onMount } from "svelte";
  import { fileStore } from "../stores/files";
  import type { MusicFile } from "../types";
  import { formatFileSize } from "../utils/formatters";

  export let onFileSelect: (file: MusicFile) => void;

  let files: MusicFile[] = [];
  let isLoading = true;
  let searchQuery = "";
  let viewMode: "grid" | "list" = "grid";

  $: filteredFiles = searchQuery
    ? files.filter(
        (file) =>
          file.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          file.metadata.artist
            .toLowerCase()
            .includes(searchQuery.toLowerCase()) ||
          file.metadata.album
            .toLowerCase()
            .includes(searchQuery.toLowerCase()) ||
          file.metadata.title.toLowerCase().includes(searchQuery.toLowerCase()),
      )
    : files;

  const unsubscribe = fileStore.subscribe((state) => {
    files = state.files;
    isLoading = state.loading;
  });

  function getFileIcon(file: MusicFile) {
    const extension = file.name.split(".").pop()?.toLowerCase();

    switch (extension) {
      case "mp3":
        return "music_note";
      case "flac":
        return "high_quality";
      case "wav":
        return "equalizer";
      case "m4a":
      case "aac":
        return "audio_file";
      default:
        return "audio_file";
    }
  }

  onMount(() => {
    return unsubscribe;
  });
</script>

<div class="h-full">
  <!-- Toolbar -->
  <div class="mb-4 flex flex-wrap items-center justify-between gap-4">
    <div class="relative">
      <div
        class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3"
      >
        <span class="material-symbols-outlined text-gray-400">search</span>
      </div>
      <input
        type="text"
        bind:value={searchQuery}
        placeholder="Search files..."
        class="block w-full rounded-lg border border-gray-700 bg-gray-800 p-2.5 pl-10 text-sm text-white placeholder-gray-400 focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
      />
    </div>

    <div class="flex items-center gap-2">
      <button
        class={`rounded-md p-2 ${viewMode === "grid" ? "bg-gray-700 text-cyan-400" : "text-gray-400"}`}
        on:click={() => (viewMode = "grid")}
        title="Grid view"
      >
        <span class="material-symbols-outlined">grid_view</span>
      </button>
      <button
        class={`rounded-md p-2 ${viewMode === "list" ? "bg-gray-700 text-cyan-400" : "text-gray-400"}`}
        on:click={() => (viewMode = "list")}
        title="List view"
      >
        <span class="material-symbols-outlined">view_list</span>
      </button>
    </div>
  </div>

  <!-- Loading state -->
  {#if isLoading}
    <div class="flex h-64 items-center justify-center">
      <div
        class="h-12 w-12 animate-spin rounded-full border-4 border-cyan-500 border-t-transparent"
      ></div>
    </div>
  {:else if filteredFiles.length === 0}
    <div class="flex h-64 flex-col items-center justify-center text-gray-400">
      <span class="material-symbols-outlined mb-4 text-6xl">folder_off</span>
      {searchQuery ? "No files matching your search" : "No music files found"}
    </div>
  {:else}
    <!-- Grid View -->
    {#if viewMode === "grid"}
      <div
        class="grid grid-cols-1 gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5"
      >
        {#each filteredFiles as file (file.uid)}
          <div
            class="group cursor-pointer rounded-lg bg-gray-800 p-4 transition-all hover:bg-gray-700 w-58"
            on:click={() => onFileSelect(file)}
            on:keydown={(e) => e.key === "Enter" && onFileSelect(file)}
            role="button"
            tabindex="0"
          >
            <div
              class="mb-2 flex h-50 w-50 items-center justify-center rounded bg-gray-900 text-cyan-500"
            >
              <img
                src="http://localhost:8080/cover/{file.uid}"
                class="overflow-hidden rounded"
                alt={file.name}
              />
            </div>
            <div
              class="overflow-hidden text-ellipsis whitespace-nowrap text-sm font-medium"
            >
              {file.metadata.title}
            </div>
            <div class="text-xs text-gray-400">
              {file.metadata.artist
                ? `${file.metadata.artist} - ${file.metadata.album || "Unknown Album"}`
                : "Unknown Artist"}
            </div>
            <div class="mt-1 flex justify-between text-xs text-gray-500">
              <span>{formatFileSize(file.size)}</span>
              <!-- <span>{file.format}</span> -->
            </div>
          </div>
        {/each}
      </div>
    {:else}
      <!-- List View -->
      <div class="overflow-hidden rounded-lg border border-gray-700">
        <table class="min-w-full divide-y divide-gray-700">
          <thead class="bg-gray-800">
            <tr>
              <th
                scope="col"
                class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-400"
              >
                Name
              </th>
              <th
                scope="col"
                class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-400"
              >
                Artist
              </th>
              <th
                scope="col"
                class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-400"
              >
                Album
              </th>
              <!-- <th -->
              <!--   scope="col" -->
              <!--   class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-400" -->
              <!-- > -->
              <!--   Format -->
              <!-- </th> -->
              <th
                scope="col"
                class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-400"
              >
                Size
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-700 bg-gray-800">
            {#each filteredFiles as file (file.uid)}
              <tr
                class="cursor-pointer transition-colors hover:bg-gray-700"
                on:click={() => onFileSelect(file)}
                on:keydown={(e) => e.key === "Enter" && onFileSelect(file)}
                tabindex="0"
              >
                <td class="whitespace-nowrap px-4.5 py-2">
                  <div class="flex items-center">
                    <img
                      src="http://localhost:8080/cover/{file.uid}"
                      class="overflow-hidden rounded h-15 mr-3"
                      alt={file.metadata.title}
                    />
                    <span class="font-medium">{file.metadata.title}</span>
                  </div>
                </td>
                <td class="whitespace-nowrap px-6 py-4 text-sm text-gray-300">
                  {file.metadata.artist || "Unknown"}
                </td>
                <td class="whitespace-nowrap px-6 py-4 text-sm text-gray-300">
                  {file.metadata.album || "Unknown"}
                </td>
                <!-- <td class="whitespace-nowrap px-6 py-4 text-sm text-gray-300"> -->
                <!--   {file.format} -->
                <!-- </td> -->
                <td class="whitespace-nowrap px-6 py-4 text-sm text-gray-300">
                  {formatFileSize(file.size)}
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    {/if}
  {/if}
</div>
