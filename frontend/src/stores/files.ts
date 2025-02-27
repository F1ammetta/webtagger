import { writable } from "svelte/store";
import type { MusicFile } from "../types";

interface FileState {
  files: MusicFile[];
  loading: boolean;
  error: string | null;
  layout: "grid" | "list";
}

function createFileStore() {
  const initialState: FileState = {
    files: [],
    loading: false,
    error: null,
    layout: "grid",
  };

  const { subscribe, set, update } = writable<FileState>(initialState);

  return {
    subscribe,
    loadFiles: async () => {
      update((state) => ({ ...state, loading: true, error: null }));

      try {
        // In production, this would be a real API call
        // Simulating API delay
        await new Promise((resolve) => setTimeout(resolve, 800));

        // Mock data response
        let files: MusicFile[] = [];

        const res = await fetch("http://localhost:8080/songs");

        files = await res.json();

        console.log(files);



        update((state) => ({ ...state, files: files, loading: false }));
      } catch (error) {
        console.error("Failed to load files:", error);
        update((state) => ({
          ...state,
          loading: false,
          error: "Failed to load music files. Please try again.",
        }));
      }
    },
    updateFile: async (updatedFile: MusicFile) => {
      update((state) => {
        // In production, this would be a real API call
        // For now, update the local state
        const updatedFiles = state.files.map((file) =>
          file.uid === updatedFile.uid ? updatedFile : file
        );

        return { ...state, files: updatedFiles };
      });

      // Simulate API call to update the file
      // In production: await fetch(`/file/${updatedFile.id}`, { method: 'POST', body: JSON.stringify(updatedFile) });
      console.log("File updated:", updatedFile);
    },
  };
}

export const fileStore = createFileStore();
