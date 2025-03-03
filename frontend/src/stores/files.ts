import { writable } from "svelte/store";
import type { MusicFile, CoverUpdate, CoverReq, FileUpload } from "../types";

interface FileState {
  files: MusicFile[];
  loading: boolean;
  error: string | null;
  layout: "grid" | "list";
}
function uint8ArrayToBase64(array: Uint8Array): string {
  // For large arrays, we need to chunk to avoid "Maximum call stack size exceeded"
  const chunk = 8192;
  let binary = '';

  for (let i = 0; i < array.length; i += chunk) {
    const slice = array.subarray(i, Math.min(i + chunk, array.length));
    binary += String.fromCharCode.apply(null, Array.from(slice));
  }

  return btoa(binary);
}

async function uploadFile(file: File, endpoint: string): Promise<Response> {
  const formData = new FormData();
  formData.append('data', file);

  const response = await fetch(endpoint, {
    method: 'POST',
    body: formData,
  });

  if (!response.ok) {
    throw new Error(`Failed to upload ${file.name}: ${response.statusText}`);
  }

  return response;
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

        const res = await fetch("/api/songs");

        files = await res.json();

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
    upload: async (files: FileList) => {
      const uploadPromises = Array.from(files).map(file => uploadFile(file, '/api/upload'));

      let newFiles: MusicFile[] = [];

      try {
        const responses = await Promise.all(uploadPromises);
        console.log('All files uploaded successfully:', responses);
        for (var res of responses) {
          let file: MusicFile = await res.json();
          newFiles.push(file);
        }
      } catch (error) {
        console.error('Error uploading files:', error);
      }


      update((state) => {
        // In production, this would be a real API call
        // For now, update the local state

        return { ...state, loading: false, files: [...state.files, ...newFiles] };
      });
    },
    updateFile: async (updatedFile: MusicFile, cover: CoverUpdate) => {
      var updatedFiles: MusicFile[] = []

      const res = await fetch(`/api/edit/${updatedFile.uid}`, {
        method: "POST",
        body: JSON.stringify(updatedFile),
        headers: {
          "Content-type": "application/json; charset=UTF-8"
        },
      });

      console.log(res);

      if (res.ok) {
        console.log("sucess");
      } else {
        updatedFile.uid = "";
      }


      update((state) => {
        // In production, this would be a real API call
        // For now, update the local state
        updatedFiles = state.files.map((file) =>
          file.uid == updatedFile.uid ? updatedFile : file
        );

        return { ...state, files: updatedFiles };
      });


      let data: CoverReq = {
        data: uint8ArrayToBase64(cover.bytes),
        type: cover.mimeType,
      };

      if (cover.update) {
        const res = await fetch(`/api/cover/set/${updatedFile.uid}`, {
          method: "POST",
          body: JSON.stringify(data),
          headers: {
            "Content-type": "application/json; charset=UTF-8"
          }
        });

        if (res.ok) {
          // TODO: Update File img in the UI
        }

      }

      // Simulate API call to update the file
      // In production: await fetch(`/file/${updatedFile.id}`, { method: 'POST', body: JSON.stringify(updatedFile) });
      console.log("File updated:", updatedFile);
    },
  };
}

export const fileStore = createFileStore();
