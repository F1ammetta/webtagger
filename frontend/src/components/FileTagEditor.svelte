<script lang="ts">
  import { onMount } from "svelte";
  import type { CoverUpdate, MusicFile } from "../types";

  export let file: MusicFile;
  export let onClose: () => void;
  export let onSave: (file: MusicFile, cover: CoverUpdate) => void;

  let editedFile: MusicFile = { ...file };
  let isLoading = false;

  function handleSubmit() {
    isLoading = true;

    // Simulate API delay
    onSave(editedFile, {
      update: updateCover,
      bytes: imageBytes,
      mimeType: imageType,
    });
    isLoading = false;
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

  let currentImageSrc: string = `/api/cover/get/${file.uid}?${Date.now()}`;
  let currentImageAlt: string = file.name;
  let imageBytes: Uint8Array = new Uint8Array();
  let updateCover: boolean = false;
  let imagePreviewUrl: string | null = null;
  let imageType: string = "";

  async function handleFileSelect(event: { currentTarget: HTMLInputElement }) {
    const input = event.currentTarget;
    const selectedFile = input.files?.[0];

    if (selectedFile) {
      imageBytes = await selectedFile.bytes();

      // Create a URL for the selected image and update the display immediately
      const objectUrl = URL.createObjectURL(selectedFile);
      currentImageSrc = objectUrl;
      imagePreviewUrl = currentImageSrc;
      currentImageAlt = selectedFile.name;
      imageType = selectedFile.type;
      updateCover = true;
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
  class="fixed inset-0 z-50 flex items-center justify-center bg-black/30 p-4"
  on:click={handleBackdropClick}
  on:keydown={handleKeydown}
  role="link"
  tabindex="0"
>
  <div
    class="w-full max-w-3xl max-h-[90vh] overflow-y-auto rounded-lg bg-gray-800 shadow-xl"
    on:click|stopPropagation
    role="link"
    tabindex="0"
    on:keydown={handleKeydown}
  >
    <div
      class="flex items-center justify-between border-b border-gray-700 px-6 py-4"
    >
      <h2 class="text-xl font-semibold text-white">
        {file.name}
      </h2>
      <button
        class="rounded-full p-2 pb-0.5 text-gray-400 hover:bg-gray-700 hover:text-white"
        on:click={onClose}
      >
        <span class="material-symbols-outlined">close</span>
      </button>
    </div>

    <form on:submit|preventDefault={handleSubmit}>
      <div class="px-6 py-4">
        <div class="flex flex-col items-center pb-8">
          <div
            class="group relative mb-4 flex h-50 w-50 items-center justify-center rounded bg-gray-900 text-cyan-500"
          >
            <img
              src={currentImageSrc}
              class="overflow-hidden rounded hover:fill-gray-50"
              alt={currentImageAlt}
            />
            <div
              class="absolute inset-0 group-hover:bg-black/70 transition-all duration-300 overflow-hidden rounded"
            >
              <span
                class="absolute inset-0 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity duration-300"
                on:click={() => {
                  document.getElementById("fileInput")?.click();
                }}
                role="button"
                tabindex="0"
                on:keypress={() => {}}
              >
                <span
                  class="select-none icon material-symbols-outlined text-white"
                  style="font-size: 3rem;">edit</span
                >
              </span>
            </div>
            <!-- Hidden file input -->
            <input
              id="fileInput"
              type="file"
              accept="image/*"
              style="display: none;"
              on:change={handleFileSelect}
            />
          </div>
          <div class="grid grid-cols-1 gap-6 md:grid-cols-2">
            <div>
              <label for="title" class="block text-sm font-medium text-gray-300"
                >Title</label
              >
              <input
                type="text"
                id="title"
                bind:value={editedFile.metadata.title}
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
                bind:value={editedFile.metadata.artist}
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
                bind:value={editedFile.metadata.album}
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>

            <div>
              <label for="year" class="block text-sm font-medium text-gray-300"
                >Year</label
              >
              <input
                type="number"
                id="year"
                bind:value={editedFile.metadata.year}
                placeholder="YYYY"
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>

            <div>
              <label for="genre" class="block text-sm font-medium text-gray-300"
                >Genre</label
              >

              <input
                type="text"
                id="genre"
                bind:value={editedFile.metadata.genre}
                class="mt-1 block w-full rounded-md border-gray-700 bg-gray-700 px-3 py-2 text-white shadow-sm focus:border-cyan-500 focus:outline-none focus:ring-1 focus:ring-cyan-500"
              />
            </div>
          </div>
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
      </div>
    </form>
  </div>
</div>
