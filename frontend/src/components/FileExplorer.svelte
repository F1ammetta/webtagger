<script lang="ts">
  import { onMount } from "svelte";
  import { fileStore } from "../stores/files";
  import type { MusicFile } from "../types";
  import { formatFileSize } from "../utils/formatters";
  import LazyImage from "./LazyImage.svelte";
  import Fuse from "fuse.js";

  export let onFileSelect: (file: MusicFile, index: number) => void;

  let files: MusicFile[] = [];
  let isLoading = true;
  let searchQuery = "";
  let viewMode: "grid" | "list";

  let fuse: Fuse<MusicFile>;
  $: {
    fuse = new Fuse(files, {
      keys: ["name", "metadata.artist", "metadata.album", "metadata.title"],
      includeScore: true,
      threshold: 0.3,
    });
  }

  $: filteredFiles = searchQuery
    ? fuse.search(searchQuery).map((result) => result.item) // Extract matched items
    : files;

  const options = ["Title", "Artist", "Album", "Size"];

  var sort = "Title";
  var inverted = false;

  let imageRefs: LazyImage[] = [];

  export function reloadImage(index: number) {
    if (imageRefs[index]) {
      imageRefs[index].reloadImage();
    }
  }

  const unsubscribe = fileStore.subscribe((state) => {
    files = state.files.sort((a, b) =>
      a.metadata.title.localeCompare(b.metadata.title),
    );
    isLoading = state.loading;
    viewMode = state.layout;
  });

  function handleSort(option: string) {
    filteredFiles = filteredFiles.sort((a, b) => {
      var val = 0;
      switch (option) {
        case options[0]:
          val = a.metadata.title.localeCompare(b.metadata.title);
          break;
        case options[1]:
          val = a.metadata.artist.localeCompare(b.metadata.artist);
          break;
        case options[2]:
          val = a.metadata.album.localeCompare(b.metadata.album);
          break;
        case options[3]:
          val = a.size - b.size;
          break;
      }
      val = sort == option && !inverted ? -val : val;
      return val;
    });
    sort = option;
    inverted = !inverted;
  }

  var showSortDropdown = false;

  onMount(() => {
    return unsubscribe;
  });
</script>

<div class="h-full">
  <!-- Toolbar -->
  <div class="mb-4 flex flex-wrap items-center justify-between gap-4">
    <div class="flex items-center gap-4">
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

      <div class="relative">
        <button
          class="rounded-md p-2 pb-0.5 hover:bg-gray-700 hover:text-cyan-400 text-gray-400"
          on:click={() => (showSortDropdown = !showSortDropdown)}
          title="Sort"
        >
          <span class="material-symbols-outlined">sort</span>
        </button>

        <!-- Dropdown menu -->
        {#if showSortDropdown}
          <div
            class="absolute z-10 right-0 mt-2 w-48 rounded-lg border border-gray-700 bg-gray-800 shadow-lg"
          >
            <div class="py-1">
              {#each options as option}
                <div
                  class="cursor-pointer px-4 py-2 text-sm text-gray-400 hover:bg-gray-700"
                  on:click={() => handleSort(option)}
                  on:keypress={() => {}}
                  role="button"
                  tabindex="0"
                >
                  Sort by {option}
                </div>
              {/each}
            </div>
          </div>
        {/if}
      </div>
    </div>

    <div class="flex items-center gap-2">
      <button
        class={`rounded-md p-2 pb-0.5 ${viewMode === "grid" ? "bg-gray-700  text-cyan-400" : "text-gray-400"}`}
        on:click={() => fileStore.changeLayout("grid")}
        title="Grid view"
      >
        <span class="material-symbols-outlined">grid_view</span>
      </button>
      <button
        class={`rounded-md p-2 pb-0.5 ${viewMode === "list" ? "bg-gray-700 text-cyan-400" : "text-gray-400"}`}
        on:click={() => fileStore.changeLayout("list")}
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
        {#each filteredFiles as file, index}
          <div
            class="group cursor-pointer rounded-lg bg-gray-800 p-4 transition-all hover:bg-gray-700 w-58"
            on:click={() => onFileSelect(file, index)}
            on:keydown={(e) => e.key === "Enter" && onFileSelect(file, index)}
            role="button"
            tabindex="0"
          >
            <div
              class="mb-2 flex h-50 w-50 items-center justify-center rounded bg-gray-900 text-cyan-500"
            >
              <LazyImage
                bind:this={imageRefs[index]}
                src={`/api/cover/get/${file.uid}`}
                alt={file.metadata.title}
                className="overflow-hidden select-none rounded"
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
              <th
                scope="col"
                class="px-6 py-3 text-left text-xs font-medium uppercase tracking-wider text-gray-400"
              >
                Size
              </th>
            </tr>
          </thead>
          <tbody class="divide-y divide-gray-700 bg-gray-800">
            {#each filteredFiles as file, index}
              <tr
                class="cursor-pointer h-25 transition-colors hover:bg-gray-700"
                on:click={() => onFileSelect(file, index)}
                on:keydown={(e) =>
                  e.key === "Enter" && onFileSelect(file, index)}
                tabindex="0"
              >
                <td class="whitespace-nowarp px-4.5 py-2">
                  <div class="flex items-center">
                    <LazyImage
                      bind:this={imageRefs[index]}
                      src={`/api/cover/get/${file.uid}`}
                      alt={file.metadata.title}
                      className="overflow-hidden w-15 select-none rounded mr-4 mt-2 mb-2"
                    />
                    <span class="font-medium w-64 truncate">
                      {file.metadata.title}</span
                    >
                  </div>
                </td>
                <td class="w-64 px-6 py-4 text-sm text-gray-300">
                  {file.metadata.artist || "Unknown"}
                </td>
                <td class="w-64 px-6 py-4 text-sm text-gray-300">
                  {file.metadata.album || "Unknown"}
                </td>
                <!-- <td class="whitespace-nowarp px-6 py-4 text-sm text-gray-300"> -->
                <!--   {file.format} -->
                <!-- </td> -->
                <td class="px-6 py-4 text-sm text-gray-300">
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
