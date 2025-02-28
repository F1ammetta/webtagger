import { writable } from "svelte/store";
import type { MusicFile, CoverUpdate } from "../types";

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
    changeLayout: async (layout: "grid" | "list") => {
      localStorage.setItem("layout", layout);
      update((state) => ({ ...state, layout: layout }));
    },
    loadFiles: async () => {
      update((state) => ({ ...state, loading: true, error: null }));

      try {
        // In production, this would be a real API call
        // Simulating API delay
        function getLayout(): "grid" | "list" {
          const val = localStorage.getItem("layout");
          if (val === "grid" || val === "list") {
            return val;
          }
          return "grid";
        }

        const layout = getLayout();

        update((state) => ({ ...state, layout: layout }));

        // Mock data response
        let files: MusicFile[] = [];

        const res = await fetch("http://localhost:8080/songs");

        files = await res.json();

        console.log(files);



        update((state) => ({ ...state, files: files, loading: false, }));
      } catch (error) {
        console.error("Failed to load files:", error);
        update((state) => ({
          ...state,
          loading: false,
          error: "Failed to load music files. Please try again.",
        }));
      }
    },
    updateFile: async (updatedFile: MusicFile, cover: CoverUpdate) => {
      update((state) => {
        // In production, this would be a real API call
        // For now, update the local state
        //
        var updatedFiles: MusicFile[] = []

        fetch(`http://localhost:8080/edit/${updatedFile.uid}`, {
          method: "POST",
          body: JSON.stringify(updatedFile),
          headers: {
            "Content-type": "application/json; charset=UTF-8"
          },
        }).then((res) => {
          if (res.ok) {
            updatedFiles = state.files.map((file) =>
              file.uid === updatedFile.uid ? updatedFile : file
            );
          }
        })

        return { ...state, files: updatedFiles };
      });

      // Simulate API call to update the file
      // In production: await fetch(`/file/${updatedFile.id}`, { method: 'POST', body: JSON.stringify(updatedFile) });
      console.log("File updated:", updatedFile);
    },
  };
}

export const fileStore = createFileStore();
