<script lang="ts">
  import { onMount, createEventDispatcher } from "svelte";
  import type { MusicFile } from "../types";
  import { formatDuration } from "../utils/formatters";

  export let file: MusicFile;
  export let onClose: () => void;
  export let onSave: (file: MusicFile) => void;

  let editedFile: MusicFile = { ...file };
  let isLoading = false;
  let activeTab: "basic" | "advanced" = "basic";

  const genres = [
    "Rock",
    "Pop",
    "Hip Hop",
    "R&B",
    "Electronic",
    "Jazz",
    "Blues",
    "Classical",
    "Country",
    "Folk",
    "Metal",
    "Punk",
    "Indie",
    "Alternative",
    "Reggae",
    "Soul",
    "Funk",
    "Disco",
    "Ambient",
    "World",
  ];

  function handleSubmit() {
    isLoading = true;

    // Simulate API delay
    setTimeout(() => {
      onSave(editedFile);
      isLoading = false;
    }, 500);
  }

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === "Escape") {
      onClose();
    }
  }

  // Handle clicks outside the modal to close it
  function handleBackdropClick(e: MouseEvent) {
    if (e.target === e.currentTarget) {
      onClose();
    }
  }

  onMount(() => {
    document.addEventListener("keydown", handleKeydown);
    return () => {
      document.removeEventListener("keydown", handleKeydown);
    };
  });
</script>

<div
  class="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-70 p-4"
  on:click={handleBackdropClick}
>
  <div
    class="w-full max-w-3xl max-h-[90vh] overflow-y-auto rounded-lg bg-gray-800 shadow-xl"
    on:click|stopPropagation
  >
    <div
      class="flex items-center justify-between border-b border-gray-700 px-6 py-4"
    >
      <h2 class="text-xl font-semibold text-white">
        Edit Tags: {file.filename}
      </h2>
      <button
        class="rounded-full p-1 text-gray-400 hover:bg-gray-700 hover:text-white"
        on:click={onClose}
      >
        <span class="material-symbols-outlined">close</span>
      </button>
    </div>

    <div class="border-b border-gray-700">
      <nav class="flex px-6">
        <button
          class={`px-4 py-3 text-sm font-medium border-b-2 ${activeTab === "basic" ? "border-cyan-500 text-cyan-500" : "border-transparent text-gray-400 hover:text-gray-300"}`}
          on:click={() => (activeTab = "basic")}
        >
          Basic Info
        </button>
        <button
          class={`px-4 py-3 text-sm font-medium border-b-2 ${activeTab === "advanced" ? "border-cyan-500 text-cyan-500" : "border-transparent text-gray-400 hover:text-gray-300"}`}
          on:click={() => (activeTab = "advanced")}
        >
          Advanced
        </button>
      </nav>
    </div>

    <form on:submit|preventDefault={handleSubmit}>
      <div class="px-6 py-4">
        {#if activeTab === "basic"}
          <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
            <div>
              <label for="title" class="block text-sm font-medium text-gray-300"
                >Title</label
              >
              <input
                type="text"
                id="title"
                bind:value={editedFile.title}
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>

            <div>
              <label
                for="artist"
                class="block text-sm font-medium text-gray-300">Artist</label
              >
              <input
                type="text"
                id="artist"
                bind:value={editedFile.artist}
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>

            <div>
              <label for="album" class="block text-sm font-medium text-gray-300"
                >Album</label
              >
              <input
                type="text"
                id="album"
                bind:value={editedFile.album}
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>

            <div>
              <label for="year" class="block text-sm font-medium text-gray-300"
                >Year</label
              >
              <input
                type="text"
                id="year"
                bind:value={editedFile.year}
                placeholder="YYYY"
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>

            <div>
              <label
                for="trackNumber"
                class="block text-sm font-medium text-gray-300"
                >Track Number</label
              >
              <input
                type="text"
                id="trackNumber"
                bind:value={editedFile.trackNumber}
                placeholder="1"
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>

            <div>
              <label for="genre" class="block text-sm font-medium text-gray-300"
                >Genre</label
              >
              <select
                id="genre"
                bind:value={editedFile.genre}
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              >
                <option value="">Select genre</option>
                {#each genres as genre}
                  <option value={genre}>{genre}</option>
                {/each}
              </select>
            </div>
          </div>
        {:else}
          <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
            <div>
              <label
                for="format"
                class="block text-sm font-medium text-gray-300">Format</label
              >
              <input
                type="text"
                id="format"
                value={editedFile.format}
                disabled
                class="mt-1 block w-full cursor-not-allowed rounded-md border-gray-700 bg-gray-800 px-3 py-2 text-gray-400 shadow-sm"
              />
            </div>

            <div>
              <label for="size" class="block text-sm font-medium text-gray-300"
                >File Size</label
              >
              <input
                type="text"
                id="size"
                value={formatFileSize(editedFile.size)}
                disabled
                class="mt-1 block w-full cursor-not-allowed rounded-md border-gray-700 bg-gray-800 px-3 py-2 text-gray-400 shadow-sm"
              />
            </div>

            <div>
              <label
                for="bitrate"
                class="block text-sm font-medium text-gray-300"
                >Bitrate (kbps)</label
              >
              <input
                type="number"
                id="bitrate"
                bind:value={editedFile.bitrate}
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>

            <div>
              <label
                for="sampleRate"
                class="block text-sm font-medium text-gray-300"
                >Sample Rate (Hz)</label
              >
              <input
                type="number"
                id="sampleRate"
                bind:value={editedFile.sampleRate}
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>

            <div>
              <label
                for="duration"
                class="block text-sm font-medium text-gray-300">Duration</label
              >
              <input
                type="text"
                id="duration"
                value={formatDuration(editedFile.duration)}
                disabled
                class="mt-1 block w-full cursor-not-allowed rounded-md border-gray-700 bg-gray-800 px-3 py-2 text-gray-400 shadow-sm"
              />
            </div>

            <div>
              <label for="path" class="block text-sm font-medium text-gray-300"
                >File Path</label
              >
              <input
                type="text"
                id="path"
                value={editedFile.path + editedFile.filename}
                disabled
                class="mt-1 block w-full cursor-not-allowed rounded-md border-gray-700 bg-gray-800 px-3 py-2 text-gray-400 shadow-sm"
              />
            </div>
          </div>
        {/if}
      </div>

      <div
        class="flex justify-end space-x-3 border-t border-gray-700 bg-gray-800 px-6 py-4"
      >
        <button
          type="button"
          class="rounded-md border border-gray-600 bg-transparent px-4 py-2 text-sm font-medium text-gray-300 hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-cyan-500"
          on:click={onClose}
          disabled={isLoading}
        >
          Cancel
        </button>
        <button
          type="submit"
          class="rounded-md bg-cyan-600 px-4 py-2 text-sm font-medium text-white hover:bg-cyan-700 focus:outline-none focus:ring-2 focus:ring-cyan-500 disabled:opacity-50"
          disabled={isLoading}
        >
          {#if isLoading}
            <span class="flex items-center">
              <svg class="mr-2 h-4 w-4 animate-spin" viewBox="0 0 24 24">
                <circle
                  class="opacity-25"
                  cx="12"
                  cy="12"
                  r="10"
                  stroke="currentColor"
                  stroke-width="4"
                  fill="none"
                />
                <path
                  class="opacity-75"
                  fill="currentColor"
                  d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                />
              </svg>
              Saving...
            </span>
          {:else}
            Save Changes
          {/if}
        </button>
      </div>
    </form>
  </div>
</div>
