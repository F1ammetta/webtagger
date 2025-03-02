<script lang="ts">
  import { onMount } from "svelte";
  import Navbar from "../components/Navbar.svelte";
  import FileExplorer from "../components/FileExplorer.svelte";
  import FileTagEditor from "../components/FileTagEditor.svelte";
  import { fileStore } from "../stores/files";
  import type { MusicFile, CoverUpdate } from "../types";

  let selectedFile: MusicFile | null = null;
  let isTagEditorOpen = false;

  let fileExplorer: FileExplorer;
  let idx: number;

  function handleFileSelect(file: MusicFile, index: number) {
    selectedFile = file;
    isTagEditorOpen = true;
    idx = index;
  }

  function closeTagEditor() {
    isTagEditorOpen = false;
    selectedFile = null;
    fileExplorer.reloadImage(idx);
  }

  onMount(() => {
    // Load files on component mount
    fileStore.loadFiles();
  });
</script>

<div class="flex h-screen flex-col">
  <Navbar />

  <main class="flex-1 overflow-y-auto">
    <div class="container mx-auto max-w-7xl px-4 py-6 pb-40">
      <FileExplorer bind:this={fileExplorer} onFileSelect={handleFileSelect} />
    </div>
  </main>

  {#if isTagEditorOpen && selectedFile}
    <FileTagEditor
      file={selectedFile}
      onClose={closeTagEditor}
      onSave={async (updatedFile, cover) => {
        await fileStore.updateFile(updatedFile, cover);
        closeTagEditor();
      }}
    />
  {/if}
</div>
