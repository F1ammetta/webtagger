<script lang="ts">
  import { onMount } from "svelte";
  import Navbar from "../components/Navbar.svelte";
  import FileExplorer from "../components/FileExplorer.svelte";
  import FileTagEditor from "../components/FileTagEditor.svelte";
  import { fileStore } from "../stores/files";
  import type { MusicFile } from "../types";

  let selectedFile: MusicFile | null = null;
  let isTagEditorOpen = false;

  function handleFileSelect(file: MusicFile) {
    selectedFile = file;
    isTagEditorOpen = true;
  }

  function closeTagEditor() {
    isTagEditorOpen = false;
    selectedFile = null;
  }

  onMount(() => {
    // Load files on component mount
    fileStore.loadFiles();
  });
</script>

<div class="flex h-screen flex-col">
  <Navbar />

  <main class="flex-1 overflow-hidden">
    <div class="container mx-auto h-full max-w-7xl px-4 py-6">
      <FileExplorer onFileSelect={handleFileSelect} />
    </div>
  </main>

  {#if isTagEditorOpen && selectedFile}
    <FileTagEditor
      file={selectedFile}
      onClose={closeTagEditor}
      onSave={(updatedFile) => {
        fileStore.updateFile(updatedFile);
        closeTagEditor();
      }}
    />
  {/if}
</div>
